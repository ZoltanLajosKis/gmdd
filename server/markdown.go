package server

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	tpl "github.com/ZoltanLajosKis/gmdd/templates"

	blf "github.com/russross/blackfriday"
	log "github.com/sirupsen/logrus"
)

type markdownServer struct {
	root string
}

type renderer struct {
	title        string
	htmlRenderer *blf.HTMLRenderer
}

var (
	mathInlineDelim = []byte("$")
	mathFenceDelim  = []byte("$$")
	mathFenceInfo   = []byte("math")
)

const (
	mdExts = blf.NoExtensions | blf.NoIntraEmphasis | blf.Tables | blf.FencedCode |
		blf.Autolink | blf.Strikethrough | blf.SpaceHeadings | blf.Footnotes |
		blf.NoEmptyLineBeforeBlock | blf.HeadingIDs | blf.Titleblock |
		blf.AutoHeadingIDs | blf.BackslashLineBreak | blf.DefinitionLists
	htmlFlags = blf.HTMLFlagsNone | blf.CompletePage
)

func newMarkdownServer(root string) *markdownServer {
	return &markdownServer{
		root: root,
	}
}

func (s *markdownServer) serve(w http.ResponseWriter, r *http.Request, path string, l *log.Entry) {
	l.Info("Serving markdown file.")

	data, err := ioutil.ReadFile(path)
	if err != nil {
		l.WithFields(log.Fields{
			"error": err,
		}).Warn("Could not read file.")

		respondStatus(w, http.StatusInternalServerError)
		return
	}

	etag := etag(data)

	if r.Header.Get("If-None-Match") == etag {
		l.Debug("File not modified.")
		w.Header().Set("ETag", etag)
		w.WriteHeader(http.StatusNotModified)
		return
	}

	title := strings.TrimPrefix(path, s.root)
	renderer := &renderer{
		title,
		blf.NewHTMLRenderer(blf.HTMLRendererParameters{
			HeadingIDPrefix: "heading-",
			Flags:           htmlFlags,
		}),
	}

	html := blf.Run(data, blf.WithExtensions(mdExts), blf.WithRenderer(renderer))

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("ETag", etag)
	w.WriteHeader(http.StatusOK)
	w.Write(html)
}

func (r *renderer) RenderHeader(w io.Writer, node *blf.Node) {
	tpl.MarkdownHeader(w, r.title)
}

func (r *renderer) RenderNode(w io.Writer, node *blf.Node, entering bool) blf.WalkStatus {
	switch node.Type {
	case blf.Code:
		if len(node.Literal) > 2 && bytes.HasPrefix(node.Literal, mathInlineDelim) &&
			bytes.HasSuffix(node.Literal, mathInlineDelim) {
			tpl.MarkdownInlineMath(w, node.Literal)
			return blf.GoToNext
		}
	case blf.CodeBlock:
		if bytes.Equal(node.Info, mathFenceInfo) {
			if !bytes.HasPrefix(node.Literal, mathFenceDelim) {
				node.Literal = bytes.Join([][]byte{mathFenceDelim, []byte(" "), node.Literal}, nil)
			}
			if !bytes.HasSuffix(node.Literal, mathFenceDelim) {
				node.Literal = bytes.Join([][]byte{node.Literal, []byte(" "), mathFenceDelim}, nil)
			}
		}
	}

	return r.htmlRenderer.RenderNode(w, node, entering)
}

func (r *renderer) RenderFooter(w io.Writer, node *blf.Node) {
	tpl.MarkdownFooter(w)
}
