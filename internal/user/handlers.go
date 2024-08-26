package user

import (
	"net/http"

	layout "youwe.com/go-web-accelerator/internal/shared"
	components "youwe.com/go-web-accelerator/internal/user/components"
	models "youwe.com/go-web-accelerator/internal/user/models"
)

func Handlers(prefix string, mux *http.ServeMux) {
	mux.HandleFunc("GET "+prefix+"/", getUser)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	meta := layout.MetaProps{
		Title: "Users",
	}

	menu := []*layout.NavProps{
		{
			Url:   "/",
			Label: "Home",
		},
		{
			Url:   "/user/",
			Label: "User",
		},
		{
			Url:   "/cms/",
			Label: "CMS",
		},
	}

	email := r.Context().Value("email")
	u := models.User{}

	if email != nil {
		u.Email = email.(string)
	}

	components.Show(meta, menu, u).Render(r.Context(), w)
}
