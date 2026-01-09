package learngoweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
	})
	mux.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Main page of images")
	})
	mux.HandleFunc("/images/Henry", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Henry's images")
	})

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
