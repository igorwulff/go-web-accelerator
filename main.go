package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"youwe.com/go-web-accelerator/handlers"
)

func main() {
	userHandler := handlers.UserHandler{}

	mux := http.NewServeMux()

	mux.Handle(
		"GET /user",
		withCors(
			withUser(
				http.HandlerFunc(
					userHandler.HandleUserShow,
				),
			),
		),
	)

	log.Fatal(http.ListenAndServe(":3000", mux))
}

func withCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

// var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var store = sessions.NewCookieStore([]byte("TEST"))

func withUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")

		if session.Values["email"] == nil {
			session.Values["email"] = "igorwulff@gmail.coasdsadm"

			// Save it before we write to the response/return from the handler.
			err := session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		//https://mkfeuhrer.medium.com/sessions-using-golang-and-redis-2b8fa91b573b

		// Add docker with Redis???

		ctx := context.WithValue(r.Context(), "email", session.Values["email"])
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
