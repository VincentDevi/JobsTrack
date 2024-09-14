package handler

import (
	"net/http"
)

type SidePanel struct {
	display bool
}

func New(side *SidePanel) *SidePanel {
	return &SidePanel{display: false}
}

func (panel *SidePanel) OpenPanel(w http.ResponseWriter, r *http.Request) {}
