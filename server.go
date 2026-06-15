package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

//go:embed pages/*
var Pages embed.FS

//go:embed static/*
var Static embed.FS

func StaticFiles() {
	staticFS, err := fs.Sub(Static, "static")
	if err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.FS(staticFS))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		content, err := Pages.ReadFile("pages/404.html")
		if err != nil {
			http.Error(w, "Could not load page", http.StatusInternalServerError)
			fmt.Printf("\033[31m[ERROR] Failed to read 404 page: %v\033[0m\n", err)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusNotFound)
		w.Write(content)
		fmt.Printf("\033[33m[WARN] 404 Not Found: %s\033[0m\n", r.URL.Path)
		return
	}
	content, err := Pages.ReadFile("pages/home.html")
	if err != nil {
		http.Error(w, "Could not load page", http.StatusInternalServerError)
		fmt.Printf("\033[31m[ERROR] Failed to read home page: %v\033[0m\n", err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(content)
	fmt.Println("[INFO] HomePage accessed")
}

func HandleRequests() {
	http.HandleFunc("/", HomePage)
	// handle 404 for any other routes
	fmt.Println("[INFO] Server starting on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	StaticFiles()
	HandleRequests()
}
