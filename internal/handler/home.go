package handler

import (
	"net/http"

	"github.com/VincentDevi/JobsTrack/internal/view"
)

type Home struct{}

func NewHomeHandler() *Home {
	return &Home{}
}

func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := view.Hello()
	view.Layout(c, "job tracks").Render(r.Context(), w)

}
