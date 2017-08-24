package server

import (
	"fmt"
	"hash/crc32"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"

	log "github.com/sirupsen/logrus"
)

type contentHandler struct {
	root      string
	dirServer *directoryServer
	mdServer  *markdownServer
}

func newContentHandler(root string) *contentHandler {
	return &contentHandler{
		root:      root,
		dirServer: newDirectoryServer(root),
		mdServer:  newMarkdownServer(root),
	}
}

func (h *contentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.WithFields(log.Fields{
				"error": r,
			}).Fatal(string(debug.Stack()))

			respondStatus(w, http.StatusInternalServerError)
			return
		}
	}()

	l := log.WithFields(log.Fields{
		"handler":    "content",
		"req.remote": r.RemoteAddr,
		"req.method": r.Method,
		"req.path":   r.URL.Path,
	})

	if r.Method != http.MethodGet {
		l.Warn("Method not allowed.")
		respondStatus(w, http.StatusMethodNotAllowed)
		return
	}

	// ServeMux already cleaned the URL path
	path := filepath.Join(h.root, r.URL.Path)

	fi, err := os.Stat(path)
	if err != nil {
		l.Warn("File not found.")
		respondStatus(w, http.StatusNotFound)
		return
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		h.serveDir(w, r, path, l)
	case mode.IsRegular():
		h.serveFile(w, r, path, l)
	default:
		l.WithFields(log.Fields{
			"mode": mode,
		}).Warn("Unexpected file mode.")

		respondStatus(w, http.StatusInternalServerError)
		return
	}
}

func (h *contentHandler) serveDir(w http.ResponseWriter, r *http.Request, path string, l *log.Entry) {
	h.dirServer.serve(w, r, path, l)
}

func (h *contentHandler) serveFile(w http.ResponseWriter, r *http.Request, path string, l *log.Entry) {
	_, raw := r.URL.Query()["raw"]
	if raw {
		l.Info("Serving raw file.")
		http.ServeFile(w, r, path)
		return
	}

	switch filepath.Ext(path) {
	case ".md", ".mdwn", ".mkd", ".mkdn", ".markdown":
		h.mdServer.serve(w, r, path, l)
	default:
		l.Info("Serving file.")
		http.ServeFile(w, r, path)
	}
}

func etag(data []byte) string {
	crc := crc32.ChecksumIEEE(data)
	return fmt.Sprintf("%08X", crc)
}

func respondStatus(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(status)
	w.Write([]byte(http.StatusText(status)))
}
