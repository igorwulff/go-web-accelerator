package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	cms "youwe.com/go-web-accelerator/internal/cms/handlers"
	user "youwe.com/go-web-accelerator/internal/user/handlers"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func main() {
	user := user.Handler{}
	cms := cms.Handler{}

	mux := http.NewServeMux()

	middlewareAuth := withCors(
		withUser(
			http.HandlerFunc(user.HandleUserShow),
		),
	)

	mux.Handle("GET /", http.HandlerFunc(cms.Handle))
	mux.Handle("GET /pages/", http.HandlerFunc(cms.Handle))

	mux.Handle("GET /user", middlewareAuth)

	log.Fatal(http.ListenAndServe(":3000", mux))
}

func withCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

// https://www.alexedwards.net/blog/working-with-cookies-in-go
// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/06.2.html
func withUser(next http.Handler) http.Handler {
	// https://www.alexedwards.net/blog/working-with-cookies-in-go
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("session-id")
		if cookie == nil {
			expire := time.Now().AddDate(0, 0, 1)
			uuid, _ := uuid.NewV7()
			sid := uuid.String()

			//client.Set(context.Background(), sid, "test", expire.Sub(time.Now()))

			client.Expire(context.Background(), sid, expire.Sub(time.Now()))
			client.HSet(
				context.Background(),
				sid,
				"email",
				"igorwulff@gmail.com",
			).Err()

			newCookie := http.Cookie{
				Name:       "session-id",
				Value:      sid,
				Path:       "/",
				Domain:     "/",
				Expires:    expire,
				RawExpires: expire.Format(time.UnixDate),
				MaxAge:     86400,
				Secure:     true,
				HttpOnly:   true,
				Raw:        "session-id=" + sid,
				Unparsed:   []string{"session-id=" + sid},
			}

			http.SetCookie(w, &newCookie)
			cookie = &newCookie

		}

		data := client.HGetAll(context.Background(), cookie.Value).Val()
		fmt.Println(data)
		ctx := context.WithValue(r.Context(), "email", data["email"])
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
