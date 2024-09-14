package handler

import (
	"net/http"

	"github.com/VincentDevi/JobsTrack/internal/view/page"
)

type Messages struct{}

func NewMessages() *Messages {
	return &Messages{}
}

func (o *Messages) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	page.OrganisationsPage().Render(r.Context(), w)
}
