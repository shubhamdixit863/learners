package middlewares

import (
	"fmt"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		// print out the request
		fmt.Println(r.URL.Path)

		if r.URL.Query().Get("token") == "" {

			// You can directly send a response

			w.Write([]byte("Please provide the token"))
			return

		}

		// this should be called when everything is right

		next.ServeHTTP(w, r)
	})
}

func LoggMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		// print out the request
		fmt.Println("Just printing out the logs", r.URL.Path)

		// this should be called when everything is right

		next.ServeHTTP(w, r)
	})
}
