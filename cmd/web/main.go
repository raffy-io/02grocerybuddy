package main

import (
	"io/fs"
	"log"
	"net/http"
	"os"

	grocerybuddy "github.com/raffy-io/02grocerybuddy"
	"github.com/raffy-io/02grocerybuddy/internal/handlers"
)


func main() {

	// routes
	mux := http.NewServeMux()

	mux.HandleFunc("/",handlers.GetHome)


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