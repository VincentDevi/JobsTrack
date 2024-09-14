package handler

import (
	"net/http"

	"github.com/VincentDevi/JobsTrack/internal/view/page"
)

type Organisations struct{}

func NewOrganisations() *Organisations {
	return &Organisations{}
}

func (o *Organisations) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	page.OrganisationsPage().Render(r.Context(), w)
}
