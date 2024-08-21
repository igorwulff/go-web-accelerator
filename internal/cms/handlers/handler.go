package cms

import (
	"net/http"
)

type Handler struct{}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	/*u := models.Cms{}

	components.Show(u).Render(r.Context(), w)*/
}
