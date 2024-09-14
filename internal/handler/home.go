package handler

import (
	"net/http"

	"github.com/VincentDevi/JobsTrack/internal/view"
	"github.com/VincentDevi/JobsTrack/internal/view/layout"
)

type Home struct {
	title string
}

func NewHomeHandler() *Home {
	return &Home{
		title: "Job Tracks",
	}
}

func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	home := NewHomeHandler()
	c := view.Hello()
	p := view.SideClose()
	layout.PageLayoutWithTable(c, p, home.title).Render(r.Context(), w)

}

func (h *Home) Open(w http.ResponseWriter, r *http.Request) {
	view.SideOpen().Render(r.Context(), w)
}

func (h *Home) Close(w http.ResponseWriter, r *http.Request) {
	view.SideClose().Render(r.Context(), w)
}
