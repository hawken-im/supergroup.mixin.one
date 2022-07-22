package routes

import (
	"encoding/json"
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/hawken-im/supergroup.mixin.one/middlewares"
	"github.com/hawken-im/supergroup.mixin.one/models"
	"github.com/hawken-im/supergroup.mixin.one/session"
	"github.com/hawken-im/supergroup.mixin.one/views"
)

type propertyImpl struct{}

func registerProperties(router *httptreemux.TreeMux) {
	impl := propertyImpl{}

	router.POST("/properties", impl.create)
}

func (impl *propertyImpl) create(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	var body struct {
		Value bool `json:"value"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		views.RenderErrorResponse(w, r, session.BadRequestError(r.Context()))
		return
	}
	if middlewares.CurrentUser(r).GetRole() != "admin" {
		views.RenderErrorResponse(w, r, session.ForbiddenError(r.Context()))
		return
	}
	_, err := models.CreateProperty(r.Context(), models.ProhibitedMessage, body.Value)
	if err != nil {
		views.RenderErrorResponse(w, r, err)
	} else {
		views.RenderBlankResponse(w, r)
	}
}
