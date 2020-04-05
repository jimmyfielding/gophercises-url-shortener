package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	urlshortener "github.com/jimmyfielding/gophercises/url_shortener/url-shortener"
)

func main() {
	mux := defaultMux()
	yamlPath := flag.String("f", "path-to-url.yaml", "-f path-to-url.yaml, filename of yaml containing mappings")

	flag.Parse()
	yamlFile, err := ioutil.ReadFile(*yamlPath)
	if err != nil {
		log.Fatal(err)
	}

	YAMLHandler, err := urlshortener.YAMLHandler(yamlFile, mux)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting server, listening on port :8080")
	http.ListenAndServe(":8080", YAMLHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	return mux
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
