package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	grocerybuddy "github.com/raffy-io/02grocerybuddy"
)


func main() {

	// routes
	mux := http.NewServeMux()

	mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"<h1>Hello World</h1>")
	})


	// static assets
	staticFS,err := fs.Sub(grocerybuddy.EmbeddedAssets,"static")
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	// server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on clean port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}