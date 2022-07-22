package routes

import (
	"net/http"
	"runtime"

	"github.com/dimfeld/httptreemux"
	"github.com/hawken-im/supergroup.mixin.one/config"
	"github.com/hawken-im/supergroup.mixin.one/views"
)

func RegisterRoutes(router *httptreemux.TreeMux) {
	registerHanders(router)

	router.GET("/", root)
	router.GET("/_hc", healthCheck)
	registerUsers(router)
	registerPackets(router)
	registerMesseages(router)
	registerProperties(router)
	registerBroadcasters(router)
}

func root(w http.ResponseWriter, r *http.Request, params map[string]string) {
	views.RenderDataResponse(w, r, map[string]string{
		"build":      config.BuildVersion + "-" + runtime.Version(),
		"developers": "https://developers.mixin.one",
	})
}

func healthCheck(w http.ResponseWriter, r *http.Request, params map[string]string) {
	views.RenderBlankResponse(w, r)
}
