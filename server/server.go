package server

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	assetsRoot      = "/__gmdd__/"
	shutdownTimeout = 5 * time.Second
)

// Start gmdd server
func Start(addr string, port int, root string) {
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt)

	contentHandler := newContentHandler(root)
	assetsHandler := newAssetsHandler()

	mux := http.NewServeMux()
	mux.Handle("/", contentHandler)
	mux.Handle("/favicon.ico", withPrefix(assetsRoot, assetsHandler))
	mux.Handle(assetsRoot, assetsHandler)

	listenAddr := fmt.Sprintf("%s:%d", addr, port)
	srv := &http.Server{Addr: listenAddr, Handler: mux}

	go func() {
		log.WithFields(log.Fields{
			"address": listenAddr,
			"root":    root,
		}).Info("Server started.")

		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.WithFields(log.Fields{
				"error": err,
			}).Fatal("Error listening.")
		}
	}()

	<-shutdown
	log.Info("Shutting down...")

	deadline := time.Now().Add(shutdownTimeout)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Error shutting down.")
	}
}

func withPrefix(prefix string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r2 := new(http.Request)
		*r2 = *r
		r2.URL = new(url.URL)
		*r2.URL = *r.URL
		r2.URL.Path = path.Join(prefix, r.URL.Path)
		h.ServeHTTP(w, r2)
	})
}
