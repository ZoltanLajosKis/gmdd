package server

import (
	"net/http"
	"os"
	"sort"
	"strings"

	tpl "github.com/ZoltanLajosKis/gmdd/templates"
	log "github.com/sirupsen/logrus"
)

type directoryServer struct {
	root string
}

func newDirectoryServer(root string) *directoryServer {
	return &directoryServer{
		root: root,
	}
}

func (s *directoryServer) serve(w http.ResponseWriter, r *http.Request, dirPath string, l *log.Entry) {
	l.Info("Serving directory.")

	f, err := os.Open(dirPath)
	if err != nil {
		l.WithFields(log.Fields{
			"error": err,
		}).Warn("Could not open directory.")

		respondStatus(w, http.StatusInternalServerError)
		return
	}

	list, err := f.Readdir(-1)
	if err != nil {
		l.WithFields(log.Fields{
			"error": err,
		}).Warn("Could not read directory.")

		respondStatus(w, http.StatusInternalServerError)
		return
	}

	var size int64

	for _, info := range list {
		if !info.IsDir() {
			size += info.Size()
		}
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].IsDir() && !list[j].IsDir() {
			return true
		}
		if !list[i].IsDir() && list[j].IsDir() {
			return false
		}
		return list[i].Name() < list[j].Name()
	})

	crumbs := strings.Split(r.URL.Path, "/")

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	tpl.Directory(w, crumbs, list, size)
}
