package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type RequestBody struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			w.Write([]byte("hello world"))
		}
		w.WriteHeader(405)
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			body, err := io.ReadAll(req.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Request body not supported"))
			}
			var re RequestBody
			err = json.Unmarshal(body, &re)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			log.Println("HEllo", re.Name)

			w.Write([]byte("hello world"))
			return
		} else if req.Method == http.MethodGet {
			log.Println(req.URL)
			w.WriteHeader(200)

			return

		}
		w.WriteHeader(405)
	})

	log.Println("Server runnning on PORT", 8080)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Println("Error in starting the server", err)
		return
	}
	log.Println("i will not be executed")
}

// You have write an application that is used to create user data ,register user
// /register -->registers user
// /get -->gets all user
// /get?id=1 --->particular user
// /update?id=1 ,you have to pass the body
// /delete?id=1 ,you have to delete the user
