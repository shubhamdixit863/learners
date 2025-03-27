package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type ServerHandler struct {
}

var Routes = map[string][]string{
	"data": {
		"create",
		"update",
	},
}

func (s *ServerHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Raw path", request.URL.Path)
	// Split the path
	pathSlice := strings.Split(request.URL.Path, "/")

	// Our own custom routes
	if _, ok := Routes[pathSlice[1]]; ok {
		// v would be the slice of child route
		if len(pathSlice) > 1 {

			// child routes are present

			switch pathSlice[2] {
			case "create":
				writer.Write([]byte("This is a create route"))
				return

			default:
				writer.Write([]byte("Not a valid route"))
				return

			}

		}

		// it is a data route

		// Handle all the methods post ,put ,get method
		switch request.Method {
		case "GET":
			writer.Write([]byte("Get route of the data"))

		case "POST":
			writer.Write([]byte("Get route of the data"))

		default:
			writer.WriteHeader(404)
		}

	}

}

func NewServerHandler() http.Handler {
	return &ServerHandler{}
}

func main() {

	customHandler := NewServerHandler()
	log.Println("Server Running on port 8080")
	err := http.ListenAndServe("localhost:8080", customHandler)
	if err != nil {
		return
	}

}
