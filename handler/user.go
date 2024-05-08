package handler

import (
	"net/http"

	"youwe.com/go-web-accelerator/model"
	"youwe.com/go-web-accelerator/view/user"
)

type UserHandler struct{}

func (h *UserHandler) HandleUserShow(w http.ResponseWriter, r *http.Request) {
	u := model.User{
		Email: r.Context().Value("email").(string),
	}
	user.Show(u).Render(r.Context(), w)
}
