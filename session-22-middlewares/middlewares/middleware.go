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

// use net http package
// you have to create two  handlers ,which gives the file list
// and another gives the filedata as per the name passed

// You have to create ,log middleare that logs following things
// first logs all things about rquest in strcuture format ,agent ,header
// is you have to create a map with the user list and the files they can access
// if the user tries to access any  other file other than what they have access too ,
//send 405 status forbidden
