package main

import (
	"context"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	grocerybuddy "github.com/raffy-io/02grocerybuddy"
	"github.com/raffy-io/02grocerybuddy/internal/db"
	"github.com/raffy-io/02grocerybuddy/internal/handlers"
)


func main() {
	ctx := context.Background()
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://postgres:secret@localhost:5431/postgres?sslmode=disable"
	}

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v\n", err)
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	// 2. Inject the queries into your handler instance
	homehandler := &handlers.HomeHandler{
		Queries: queries,
	}
	todosHandler := &handlers.TodosHandler{
		Queries: queries,
	}

	// routes
	mux := http.NewServeMux()

	mux.HandleFunc("GET /",homehandler.ListTodo)
	mux.HandleFunc("GET /add",todosHandler.AddTodo)
	mux.HandleFunc("POST /add",todosHandler.InsertTodo)
	mux.HandleFunc("POST /delete",todosHandler.DeleteTodo)


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