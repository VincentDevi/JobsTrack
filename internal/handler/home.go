package handler

import (
	"net/http"

	"github.com/VincentDevi/JobsTrack/internal/view/page"
	"github.com/VincentDevi/JobsTrack/internal/view/ui"
)

type Home struct{}

func NewHomeHandler() *Home {
	return &Home{}
}

func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	page.HomePage().Render(r.Context(), w)

}

func (h *Home) Open(w http.ResponseWriter, r *http.Request) {
	ui.SideOpen().Render(r.Context(), w)
}

func (h *Home) Close(w http.ResponseWriter, r *http.Request) {
	ui.SideClose().Render(r.Context(), w)
}
