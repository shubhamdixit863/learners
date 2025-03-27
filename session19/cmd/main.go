package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type RequestBody struct {
	Name string `json:"name"`
}

//type CustomResponse struct {
//
//}
//
//func (c CustomResponse) Header() http.Header {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (c CustomResponse) Write(bytes []byte) (int, error) {
//
//}
//
//func (c CustomResponse) WriteHeader(statusCode int) {
//	//TODO implement me
//	panic("implement me")
//}

//func NewCustomResponse() http.ResponseWriter {
//	return &CustomResponse{}
//}

/*
package main

import "fmt"

	type Writer interface {
		Write() string
	}

	type HttpResponseWriter interface {
		Print() string
		Write() string
	}

type Data struct {
}

	func (d *Data) Write() string {
		return "hello from Data"
	}

	func (d *Data) Print() string {
		return "hello from Data"
	}

	func MyFun(w HttpResponseWriter) {
		fmt.Println(w.Write())

}

	func main() {
		var h Writer
		h = &Data{}
		MyFun(h)
	}
*/
func main() {
	http.Handle("/", http.FileServer(http.Dir("../internal/static/")))

	http.HandleFunc("/data", func(writer http.ResponseWriter, request *http.Request) {

		//read the file

		file, err := os.ReadFile("../internal/static/data.html")
		if err != nil {
			return
		}
		writer.Write(file)
	})

	http.HandleFunc("/datacustom", func(writer http.ResponseWriter, request *http.Request) {

		//read the file

		http.ServeFile(writer, request, "../internal/static/data.html")

	})

	// How to serve html
	// how to send json response

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
			//log.Println("HEllo", re.Name)
			//log.Println(reflect.TypeOf(w))
			//http.Response{} /// default response object which implements the http.ResponseWriter
			w.Header().Add("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(&re)
			if err != nil {
				return
			}
			//w.Write([]byte("hello world"))
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
