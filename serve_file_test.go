package golangweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "./resources/index.html")
	} else {
		http.ServeFile(w, r, "./resources/not-found.html")

	}

}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/index.html
var resourceOk string

//go:embed resources/not-found.html
var resourceNotFound string

func ServeFileGolangEmbed(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		fmt.Fprint(w, resourceOk)
	} else {
		fmt.Fprint(w, resourceNotFound)
	}
}

func TestServeFileGolangEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: http.HandlerFunc(ServeFileGolangEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
