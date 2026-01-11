package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	
	"github.com/gorilla/mux"
)

// Handler untuk Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
    // Routing manual karena Vercel punya routing sendiri
    path := r.URL.Path
    
    switch {
    case path == "/api/convert" && r.Method == "POST":
        convertHandler(w, r)
    case strings.HasPrefix(path, "/download/"):
        downloadHandler(w, r)
    default:
        // Serve static files dari frontend folder
        serveStaticHandler(w, r)
    }
}

// ===== TAMBAHKIN INI DI BAWAH =====
func serveStaticHandler(w http.ResponseWriter, r *http.Request) {
    // Routing untuk file static
    staticDir := "./frontend"
    
    // Jika root, serve index.html
    if r.URL.Path == "/" {
        http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
        return
    }
    
    // Cek file yang diminta
    requestedFile := filepath.Join(staticDir, r.URL.Path)
    if _, err := os.Stat(requestedFile); err == nil {
        http.ServeFile(w, r, requestedFile)
        return
    }
    
    // Fallback ke index.html untuk client-side routing
    http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
}

// ===== UPDATE vercel.json =====
