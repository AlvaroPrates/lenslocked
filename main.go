package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/AlvaroPrates/lenslocked/controllers"
	"github.com/AlvaroPrates/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	homeTpl := views.Must(views.Parse(filepath.Join("templates", "home.html")))
	r.Get("/", controllers.StaticHandler(homeTpl))

	contactTpl := views.Must(views.Parse(filepath.Join("templates", "contact.html")))
	r.Get("/contact", controllers.StaticHandler(contactTpl))

	faqTpl := views.Must(views.Parse(filepath.Join("templates", "faq.html")))
	r.Get("/faq", controllers.StaticHandler(faqTpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
