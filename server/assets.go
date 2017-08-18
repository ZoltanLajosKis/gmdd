package server

import (
	"net/http"

	"github.com/ZoltanLajosKis/gmdd/assets"
)

type assetsHandler struct {
}

func newAssetsHandler() http.Handler {
	return http.FileServer(assets.Assets)
}
