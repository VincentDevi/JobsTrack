package handler

import (
	"net/http"

	"github.com/VincentDevi/JobsTrack/internal/view/page"
)

type Offers struct{}

func NewOffers() *Offers {
	return &Offers{}
}

func (o *Offers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	page.OrganisationsPage().Render(r.Context(), w)
}
