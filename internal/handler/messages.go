package handler

import (
	"net/http"

	"github.com/VincentDevi/JobsTrack/internal/model"
	"github.com/VincentDevi/JobsTrack/internal/view/page"
)

type Messages struct {
	messages []model.MessagesPageElement
}

func NewMessages() *Messages {
	return &Messages{}
}

func (o *Messages) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	messages := model.DummiesMessagesPageElement()
	page.MessagesPage(messages).Render(r.Context(), w)
}
