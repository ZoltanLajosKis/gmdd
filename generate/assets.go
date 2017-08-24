// +build ignore

package main

import (
	"log"

	as "github.com/ZoltanLajosKis/go-assets"
)

var (
	sources = []*as.Source{
		{"favicon.ico",
			"generate/assets/favicon.ico", nil, nil},
		{"__gmdd__/gmdd.css",
			"generate/assets/gmdd.css", nil, nil},
		{"__gmdd__/gmdd-md.js",
			"generate/assets/gmdd-md.js", nil, nil},
		{"__gmdd__/gmdd-dir.js",
			"generate/assets/gmdd-dir.js", nil, nil},
		{"__gmdd__/github-markdown.min.css",
			"https://cdnjs.cloudflare.com/ajax/libs/github-markdown-css/2.8.0/github-markdown.min.css",
			&as.Checksum{as.MD5, "0d424ff347a923913a99682bffda185b"}, nil},
		{"__gmdd__/katex",
			"https://github.com/Khan/KaTeX/releases/download/v0.8.1/katex.zip",
			&as.Checksum{as.MD5, "90766d8ff1ac3cefef02c461b5e071d2"},
			&as.Archive{as.Zip, as.ReMap(
				"(katex/(contrib.*|fonts.*|images.*|katex\\.min\\.(css|js)))",
				"__gmdd__/${1}")}},
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
		{"__gmdd__/moment.min.js",
			"https://cdnjs.cloudflare.com/ajax/libs/moment.js/2.18.1/moment.min.js",
			&as.Checksum{as.MD5, "aeb7908241d9f6d5a45e504cc4f2ec15"}, nil},
		{"__gmdd__/filesize.min.js",
			"https://cdnjs.cloudflare.com/ajax/libs/filesize/3.5.10/filesize.min.js",
			&as.Checksum{as.MD5, "c38d33dff49cb465987577c14c0c9200"}, nil},
		{"__gmdd__/tablesort",
			"https://github.com/tristen/tablesort/archive/v5.0.1.zip",
			&as.Checksum{as.MD5, "33d8f31e23de760cd20e65d7ad709777"},
			&as.Archive{as.Zip, as.ReMap(
				"tablesort-5.0.1/dist/(?:(tablesort.min.js)|sorts/(tablesort.number.min.js))",
				"__gmdd__/${1}${2}")}},
	}
)

func main() {
	err := as.Compile(sources, "assets/assets.go", "assets", "Assets", nil)
	if err != nil {
		log.Panic(err)
	}
}
