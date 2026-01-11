package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "web-to-apk-converter/backend/pkg"
    "github.com/gorilla/mux"
)

type ConvertRequest struct {
    URL     string `json:"url"`
    AppName string `json:"appName"`
}

type ConvertResponse struct {
    Success      bool   `json:"success"`
    DownloadLink string `json:"downloadLink,omitempty"`
    Error        string `json:"error,omitempty"`
}

func main() {
    r := mux.NewRouter()
    
    // API routes
    r.HandleFunc("/api/convert", convertHandler).Methods("POST")
    r.HandleFunc("/download/{filename}", downloadHandler).Methods("GET")
    
    // Serve frontend
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend")))
    
    // Create necessary directories
    os.MkdirAll("generated_apks", 0755)
    
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
    fmt.Printf("🔥 Jawir's APK Converter running on port %s\n", port)
    fmt.Println("😈 Mode: HACKER SUPER AKTIF!")
    log.Fatal(http.ListenAndServe(":"+port, r))
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    var req ConvertRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        json.NewEncoder(w).Encode(ConvertResponse{
            Success: false,
            Error:   "Invalid JSON format",
        })
        return
    }
    
    // Validate URL
    if !strings.HasPrefix(req.URL, "http") {
        req.URL = "https://" + req.URL
    }
    
    // Build APK
    builder := pkg.NewAPKBuilder(req.URL, req.AppName)
    apkPath, err := builder.Build()
    
    if err != nil {
        json.NewEncoder(w).Encode(ConvertResponse{
            Success: false,
            Error:   fmt.Sprintf("Build failed: %v", err),
        })
        return
    }
    
    filename := filepath.Base(apkPath)
    json.NewEncoder(w).Encode(ConvertResponse{
        Success:      true,
        DownloadLink: "/download/" + filename,
    })
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    filename := vars["filename"]
    
    apkPath := filepath.Join("generated_apks", filename)
    
    if _, err := os.Stat(apkPath); os.IsNotExist(err) {
        http.Error(w, "APK not found", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/vnd.android.package-archive")
    w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
    http.ServeFile(w, r, apkPath)
}