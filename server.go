package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var PagesDir = filepath.Join(execDir, "pages")

func StaticFiles() {
	dir := filepath.Join(execDir, "static")
	fileServer := http.FileServer(http.Dir(dir))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	fmt.Printf("[INFO] Serving static files from %s\n", dir)
}

func HandleRequests(addr string) {
	// handle pages
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/news", NewsListPageHandler)
	http.HandleFunc("/news/", NewsPageHandler)
	http.HandleFunc("/events", EventsListPageHandler)
	http.HandleFunc("/events/", EventPageHandler)
	http.HandleFunc("/feedback", FeedbackPageHandler)
	http.HandleFunc("/feedback/new", NewFeedbackPageHandler)
	http.HandleFunc("GET /feedback/view", ViewFeedbackPageHandler)
	http.HandleFunc("/about", AboutPageHandler)
	http.HandleFunc("/contact", ContactPageHandler)
	http.HandleFunc("/secret", SecretPageHandler)

	// handle API
	http.HandleFunc("/api/", ApiNotFoundHandler) // catch-all for undefined API routes
	http.HandleFunc("/api/news", NewsApiHandler)
	http.HandleFunc("/api/events", EventsApiHandler)
	http.HandleFunc("/api/feedbacks", GetFeedbacksHandler)
	http.HandleFunc("POST /api/post-feedback", PostFeedbackHandler)
	http.HandleFunc("POST /api/post-reply", AddReplyHandler)
	http.HandleFunc("POST /api/vErIfYsEcReTpAsSwOrD", VerifySecretPasswordHandler)
	http.HandleFunc("POST /api/secret-cmd", SecretCommandHandler)
	// handle 404 for any other routes
	fmt.Printf("[INFO] Server starting on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func renderNotFoundPage(w http.ResponseWriter, requestPath string) {
	content, err := os.ReadFile(filepath.Join(PagesDir, "404.html"))
	if err != nil {
		http.Error(w, "Could not load page", http.StatusInternalServerError)
		fmt.Printf("\033[31m[ERROR] Failed to read 404 page for %s: %v\033[0m\n", requestPath, err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	w.Write(content)
	fmt.Printf("\033[33m[WARN] 404 Not Found: %s\033[0m\n", requestPath)
}

func renderFile(w http.ResponseWriter, filePathRel string) {
	filePath := filepath.Join(PagesDir, filePathRel)
	content, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Could not load page", http.StatusInternalServerError)
		fmt.Printf("\033[31m[ERROR] Failed to read file %s: %v\033[0m\n", filePath, err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(content)
	fmt.Printf("[INFO] File served: %s\n", filePath)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderNotFoundPage(w, r.URL.Path)
		return
	}

	renderFile(w, "home.html")
}

func NewsListPageHandler(w http.ResponseWriter, r *http.Request) {
	renderFile(w, "news.html")
}

func EventsListPageHandler(w http.ResponseWriter, r *http.Request) {
	renderFile(w, "events.html")
}

func NewsPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[INFO] News page requested: %s\n", r.URL.Path)
	if r.URL.Path == "/news/" {
		http.Redirect(w, r, "/news", http.StatusMovedPermanently)
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/news/")
	content, err := os.ReadFile(filepath.Join(execDir, "data", "pages", fmt.Sprintf("news-%s.html", id)))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			renderNotFoundPage(w, r.URL.Path)
			return
		}

		http.Error(w, "Could not load page", http.StatusInternalServerError)
		fmt.Printf("\033[31m[ERROR] Failed to read news page %s: %v\033[0m\n", id, err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(content)
	fmt.Printf("[INFO] News page accessed: %s\n", id)
}

func EventPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[INFO] Event page requested: %s\n", r.URL.Path)
	if r.URL.Path == "/events/" {
		http.Redirect(w, r, "/events", http.StatusMovedPermanently)
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/events/")
	content, err := os.ReadFile(filepath.Join(execDir, "data", "pages", fmt.Sprintf("event-%s.html", id)))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			renderNotFoundPage(w, r.URL.Path)
			return
		}

		http.Error(w, "Could not load page", http.StatusInternalServerError)
		fmt.Printf("\033[31m[ERROR] Failed to read event page %s: %v\033[0m\n", id, err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(content)
	fmt.Printf("[INFO] Event page accessed: %s\n", id)
}

func FeedbackPageHandler(w http.ResponseWriter, r *http.Request) {
	renderFile(w, "feedback.html")
}

func NewFeedbackPageHandler(w http.ResponseWriter, r *http.Request) {
	renderFile(w, "new-feedback.html")
}

func ViewFeedbackPageHandler(w http.ResponseWriter, r *http.Request) {
	renderFile(w, "view-feedback.html")
}

func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	renderFile(w, "about.html")
}

func ContactPageHandler(w http.ResponseWriter, r *http.Request) {
	renderFile(w, "contact.html")
}

func SecretPageHandler(w http.ResponseWriter, r *http.Request) {
	renderFile(w, "secret.html")
}

func main() {
	StaticFiles()
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		port = "8080"
	}
	HandleRequests(host + ":" + port)
}
