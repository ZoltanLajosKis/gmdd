package main

//go:generate go run generate/assets.go
//go:generate go run generate/templates.go

import (
	"os"

	srv "github.com/ZoltanLajosKis/gmdd/server"
	kpn "github.com/alecthomas/kingpin"
	log "github.com/sirupsen/logrus"
)

var (
	version  string
	revision string

	root = kpn.Arg("rootDir", "Root directory to serve.").String()
	addr = kpn.Flag("addr", "Address to bind to.").Short('a').Default("").
		OverrideDefaultFromEnvar("GMDD_Addr").String()
	port = kpn.Flag("port", "Listen port.").Short('p').Default("8000").
		OverrideDefaultFromEnvar("GMDD_PORT").Int()
	logLevel = kpn.Flag("loglevel", "Log level").Short('l').Default("warn").
			OverrideDefaultFromEnvar("GMDD_LOGLEVEL").String()
)

func main() {
	kpn.Parse()

	lvl, err := log.ParseLevel(*logLevel)
	if err != nil {
		lvl = log.WarnLevel
		log.WithFields(log.Fields{
			"error": err,
		}).Warn("Invalid log level; using warn.")
	}
	log.SetLevel(lvl)

	if *root == "" {
		*root, err = os.Getwd()
		if err != nil {
			log.Panic(err)
		}
	}

	log.WithFields(log.Fields{
		"version":  version,
		"revision": revision,
	}).Debug("gmdd")

	srv.Start(*addr, *port, *root)
}
