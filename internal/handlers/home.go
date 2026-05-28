package handlers

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/raffy-io/02grocerybuddy/internal/db"
	"github.com/raffy-io/02grocerybuddy/ui/layout"
	"github.com/raffy-io/02grocerybuddy/ui/pages"
)

type HomeHandler struct {
	Queries *db.Queries
}

func (h *HomeHandler) ListTodo(w http.ResponseWriter, r *http.Request){

	ctx := r.Context() 
	data,err := h.Queries.ListTodos(ctx)
	if err != nil {
		log.Printf("Data not found: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	component := pages.Home(data)
	pageLayout := layout.Base("Welcome",component)
	templ.Handler(pageLayout).ServeHTTP(w,r)
}