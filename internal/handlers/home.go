package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/raffy-io/02grocerybuddy/ui/layout"
	"github.com/raffy-io/02grocerybuddy/ui/pages"
)

func GetHome(w http.ResponseWriter, r *http.Request){
	component := pages.Home()
	pageLayout := layout.Base("Welcome",component)
	templ.Handler(pageLayout).ServeHTTP(w,r)
}