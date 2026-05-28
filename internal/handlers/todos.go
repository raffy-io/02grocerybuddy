package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/raffy-io/02grocerybuddy/internal/db"
	"github.com/raffy-io/02grocerybuddy/ui/layout"
	"github.com/raffy-io/02grocerybuddy/ui/pages"
)

type TodosHandler struct {
	Queries *db.Queries
}

func (h *TodosHandler) AddTodo(w http.ResponseWriter, r *http.Request){
	component := pages.AddTodo()
	pageLayout := layout.Base("Add Item",component)
	templ.Handler(pageLayout).ServeHTTP(w,r)
}

func (h *TodosHandler) InsertTodo(w http.ResponseWriter, r *http.Request) {
 
    // Use PostFormValue to guarantee it comes from the form body, not the URL query
    name := r.PostFormValue("name")

    // Simple, explicit Go validation
    if name == "" {
        http.Error(w, "name is required", http.StatusBadRequest)
        return
    }

	finalName := strings.ToLower(name)

    // Pass the request context down to SQLC
	ctx := r.Context()

	// SQLC now returns the created row (todo) and an error
	err := h.Queries.CreateTodo(ctx, finalName)
	if err != nil {
		log.Printf("Error inserting todo: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (h *TodosHandler) DeleteTodo(w http.ResponseWriter, r *http.Request){
	// Extract the hidden field value
    idValue := r.FormValue("id")
    if idValue == "" {
        http.Error(w, "Missing item ID", http.StatusBadRequest)
        return
    }

	// Convert string to integer (assuming your DB uses integer IDs)
    id, err := strconv.Atoi(idValue)
    if err != nil {
        log.Printf("Invalid ID format: %v\n", err)
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

	// 4. Execute your SQLc generated query
    ctx := r.Context()
    err = h.Queries.DeleteTodo(ctx, int64(id)) // Cast to int64 if required by your DB schema
    if err != nil {
        log.Printf("Failed to delete todo: %v\n", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    // 5. The Go Way: Redirect back to the homepage to refresh the data list
    // StatusSeeOther (303) changes the browser's original POST method into a clean GET request for the target page
    http.Redirect(w, r, "/", http.StatusSeeOther)
}