package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

// index
func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

// load content
func getContentHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "content.html", nil); err != nil {
		log.Println("Template rendering error:", err)
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// from submit
func submitFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	response := fmt.Sprintf("<div id='form-result'>Hello, %s!</div>", name)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(response))
}

// result
func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	results := []string{
		"Result 1: " + query,
		"Result 2: " + query,
		"Result 3: " + query,
	}
	if err := templates.ExecuteTemplate(w, "search-results.html", results); err != nil {
		log.Println("Template rendering error:", err)
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// delete
func deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("")) //
}

// load more
func getMoreItemsHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1" // 默认页码
	}
	items := []string{
		fmt.Sprintf("Item A (Page %s)", page),
		fmt.Sprintf("Item B (Page %s)", page),
		fmt.Sprintf("Item C (Page %s)", page),
	}

	var html string
	for _, item := range items {
		html += fmt.Sprintf("<div class='item'>%s</div>", item)
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/get-content", getContentHandler)
	http.HandleFunc("/submit-form", submitFormHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/delete-item/", deleteItemHandler)
	http.HandleFunc("/get-more-items", getMoreItemsHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
