// +build ignore

package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/benbjohnson/ego"
)

func main() {
	err := os.MkdirAll("templates", 0755)
	if err != nil {
		log.Panic(err)
	}

	prefix := "generate/templates"

	filepath.Walk(prefix, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, ".go") {
			return nil
		}

		log.Printf("Processing template: %s ...", path)

		tpl, err := ego.ParseFile(path)
		if err != nil {
			log.Panic("Error parsing template.", err)
		}

		tplPath := strings.TrimPrefix(path, prefix)
		tplPath = strings.Join([]string{tplPath, ".go"}, "")
		tplPath = filepath.Join("templates", tplPath)

		f, err := os.Create(tplPath)
		if err != nil {
			log.Panic("Error generating template code.", err)
		}
		defer f.Close()

		_, err = tpl.WriteTo(f)
		if err != nil {
			log.Panic("Error generating template code.", err)
		}

		return nil
	})
}
