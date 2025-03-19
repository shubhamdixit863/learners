package main

import (
	"html"
	"io"
	"log"
	"net/http"
	"strings"
)

//https://usf-cs272-s25.github.io/top10/

type Data struct {
	UserId int    `json:"userId"`
	Title  string `json:"title"`
}

func main() {

	data, err := http.Get("https://usf-cs272-s25.github.io/top10/")
	if err != nil {
		return
	}

	body := data.Body

	all, err := io.ReadAll(body)
	if err != nil {
		return
	}

	//var data1 Data
	//
	//err = json.Unmarshal(all, &data1)
	//if err != nil {
	//	return
	//}
	//
	//parsedString := html.EscapeString(string(all))
	//log.Println(parsedString)
	//mp:=map[string] string
	//ur:="google.com"
	//str:="hello there guys how are you doing today"
	//
	//d:=strings.Split(str,"")
	//for _,v:=range d{
	//	if V,oK:=mp[v];oK{
	//
	//
	//	}
	//}

}

//  CrawlService  -- would traverse the link recursively ,downloads the data  // in backrgound as a go routine
//  ExtractService -- it will extract the words from the text  //t
// CleanService
//SearchService --it will search the data
// InvIndex map[string]map[string]int

var mp = map[string]map[string]int{
	"dracula": {
		"http://one.com": 2,
		"http://uo.co,":  78,
	},
}

/**
func clean(host string, href string) string {
	// Make sure host ends with slash
	if !strings.HasSuffix(host, "/") {
		host += "/"
	}

	// Parse href URL
	parsedHref, err := url.Parse(href)
	if err == nil && (parsedHref.Scheme == "http" || parsedHref.Scheme == "https") {
		// If href is an absolute URL (http/https), return it as is
		return parsedHref.String()
	}

	// use it with host if the href is relative URL
	if parsedHref != nil {
		return host + strings.TrimPrefix(parsedHref.Path, "/")
	}
	return host
}

*/

//
