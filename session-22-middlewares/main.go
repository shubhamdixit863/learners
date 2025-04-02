package main

import (
	"net/http"
	"session-22-middlewares/middlewares"
)

func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte(message))
	})
}
func main() {

	http.Handle("/create", middlewares.LoggMiddleware(middlewares.Middleware(messageHandler("hello world"))))
	http.ListenAndServe("localhost:8080", nil)

}
