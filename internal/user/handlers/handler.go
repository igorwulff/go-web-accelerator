package user

import (
	"net/http"

	components "youwe.com/go-web-accelerator/internal/user/components"
	models "youwe.com/go-web-accelerator/internal/user/models"
)

type Handler struct{}

func (h *Handler) HandleUserShow(w http.ResponseWriter, r *http.Request) {
	u := models.User{
		Email: r.Context().Value("email").(string),
	}
	components.Show(u).Render(r.Context(), w)
}
