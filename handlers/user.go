package handlers

import (
	"net/http"

	"youwe.com/go-web-accelerator/components"
	"youwe.com/go-web-accelerator/models"
)

type UserHandler struct{}

func (h *UserHandler) HandleUserShow(w http.ResponseWriter, r *http.Request) {
	u := models.User{
		Email: r.Context().Value("email").(string),
	}
	components.Show(u).Render(r.Context(), w)
}
