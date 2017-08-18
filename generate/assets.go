// +build ignore

package main

import (
	"log"
	"os"
	"regexp"
	"strings"

	as "github.com/ZoltanLajosKis/go-assets"
)

var (
	sources = []*as.Source{
		{"favicon.ico",
			"generate/assets/favicon.ico", nil, nil},
		{"__gmdd__/gmdd.css",
			"generate/assets/gmdd.css", nil, nil},
		{"__gmdd__/gmdd.js",
			"generate/assets/gmdd.js", nil, nil},
		{"__gmdd__/github-markdown.css",
			"https://raw.githubusercontent.com/sindresorhus/github-markdown-css/v2.8.0/github-markdown.css",
			&as.Checksum{as.MD5, "8f563f252b6ce60044212a5f6465494d"}, nil},
		{"__gmdd__/katex",
			"https://github.com/Khan/KaTeX/releases/download/v0.8.1/katex.zip",
			&as.Checksum{as.MD5, "90766d8ff1ac3cefef02c461b5e071d2"},
			&as.Archive{as.Zip, func(path string) string {
				re := regexp.MustCompile("katex/(contrib.*|fonts.*|images.*|katex\\.min\\.(css|js))")
				if re.MatchString(path) {
					return strings.Join([]string{"__gmdd__/", path}, "")
				}
				return ""
			}}},
		{"__gmdd__/mermaid.min.css",
			"https://unpkg.com/mermaid@7.0.4/dist/mermaid.min.css",
			&as.Checksum{as.MD5, "d8ba2f2dc1bda6ab3bca4a9ba21dbd88"}, nil},
		{"__gmdd__/mermaid.min.js",
			"https://unpkg.com/mermaid@7.0.4/dist/mermaid.min.js",
			&as.Checksum{as.MD5, "0518927a4ae49af017d50ebdebbc48b0"}, nil},
		{"__gmdd__/highlightjs.min.css",
			"http://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/styles/default.min.css",
			&as.Checksum{as.MD5, "5133d11fbaf87d3978cf403eba33c764"}, nil},
		{"__gmdd__/highlightjs.min.js",
			"http://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/highlight.min.js",
			&as.Checksum{as.MD5, "87cfd4f9aaf9cbe85f70454128541748"}, nil},
	}
)

func main() {
	err := os.MkdirAll("assets", 0755)
	if err != nil {
		log.Panic(err)
	}
	err = as.Compile(sources, "assets/assets.go", "assets", "Assets", nil)
	if err != nil {
		log.Panic(err)
	}
}
