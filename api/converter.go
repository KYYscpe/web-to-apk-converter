package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Handler wajib namanya Handler untuk Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	
	// Routing manual
	path := r.URL.Path
	
	switch {
	case path == "/api/convert" && r.Method == "POST":
		handleConvert(w, r)
	case path == "/api/status" && r.Method == "GET":
		handleStatus(w, r)
	default:
		handleRoot(w, r)
	}
}

// ========== HANDLER CONVERT ==========
func handleConvert(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		URL     string `json:"url"`
		AppName string `json:"appName"`
	}
	
	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    struct {
			DownloadLink string `json:"downloadLink"`
			FileName     string `json:"fileName"`
			BuildID      string `json:"buildId"`
		} `json:"data"`
		Warning string `json:"warning,omitempty"`
	}
	
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "Invalid JSON format",
		})
		return
	}
	
	// Validasi
	if req.URL == "" || req.AppName == "" {
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "URL and AppName are required",
		})
		return
	}
	
	// Generate APK info (SIMULASI DULU)
	buildID := fmt.Sprintf("jawir-%d", time.Now().Unix())
	fileName := fmt.Sprintf("%s.apk", sanitizeName(req.AppName))
	
	resp := Response{
		Success: true,
		Message: "🔥 APK generation started!",
		Warning: "This is simulation mode. For real APK, use Android SDK.",
	}
	resp.Data.DownloadLink = fmt.Sprintf("/api/download/%s", fileName)
	resp.Data.FileName = fileName
	resp.Data.BuildID = buildID
	
	json.NewEncoder(w).Encode(resp)
}

// ========== HANDLER STATUS ==========
func handleStatus(w http.ResponseWriter, r *http.Request) {
	status := map[string]interface{}{
		"status":    "ONLINE",
		"service":   "Jawir AI Web2APK Converter",
		"version":   "3.0.0",
		"mode":      "HACKER-SUPER 😈",
		"owner":     "Jawir12344",
		"whatsapp":  "https://whatsapp.com/channel/0029VbCFBalISTkRNOYrwQ3",
		"message":   "System ready for revenge against Nasihuy 🔥",
		"timestamp": time.Now().Unix(),
		"endpoints": []string{
			"POST /api/convert - Convert web to APK",
			"GET  /api/status  - Check system status",
			"GET  /            - Web interface",
		},
	}
	
	json.NewEncoder(w).Encode(status)
}

// ========== HANDLER ROOT ==========
func handleRoot(w http.ResponseWriter, r *http.Request) {
	// Serve HTML langsung
	html := `<!DOCTYPE html>
<html>
<head>
    <title>🔥 JAWIR AI - Web to APK Converter</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: 'Courier New', monospace;
            background: linear-gradient(135deg, #0f0f0f, #1a1a1a);
            color: #00ff00;
            min-height: 100vh;
            padding: 20px;
        }
        .container {
            max-width: 800px;
            margin: 0 auto;
            border: 2px solid #ff0033;
            border-radius: 15px;
            padding: 30px;
            background: rgba(0, 0, 0, 0.8);
            box-shadow: 0 0 30px #ff0033;
        }
        h1 {
            color: #ff0033;
            text-align: center;
            margin-bottom: 20px;
            font-size: 2.5em;
            text-shadow: 0 0 10px #ff0033;
        }
        .subtitle {
            text-align: center;
            color: #00ff00;
            margin-bottom: 30px;
            font-size: 1.2em;
        }
        .input-group {
            margin-bottom: 20px;
        }
        label {
            display: block;
            margin-bottom: 8px;
            color: #00ff00;
            font-weight: bold;
        }
        input {
            width: 100%;
            padding: 15px;
            background: #111;
            border: 1px solid #00ff00;
            border-radius: 8px;
            color: white;
            font-size: 16px;
        }
        input:focus {
            outline: none;
            border-color: #ff0033;
            box-shadow: 0 0 10px #ff0033;
        }
        button {
            width: 100%;
            padding: 18px;
            background: linear-gradient(45deg, #ff0033, #cc0022);
            color: white;
            border: none;
            border-radius: 8px;
            font-size: 18px;
            font-weight: bold;
            cursor: pointer;
            transition: all 0.3s;
            margin-top: 20px;
        }
        button:hover {
            background: linear-gradient(45deg, #ff0044, #dd0033);
            box-shadow: 0 0 20px #ff0033;
            transform: translateY(-2px);
        }
        .status {
            margin-top: 30px;
            padding: 20px;
            border-radius: 8px;
            background: #111;
            border: 1px solid #333;
            min-height: 100px;
        }
        .success { color: #00ff00; border-color: #00ff00; }
        .error { color: #ff0033; border-color: #ff0033; }
        .loading { color: #ffff00; }
        .footer {
            margin-top: 40px;
            text-align: center;
            color: #666;
            font-size: 0.9em;
        }
        .mode-badge {
            display: inline-block;
            background: #ff0033;
            color: white;
            padding: 5px 15px;
            border-radius: 20px;
            font-size: 0.8em;
            margin-left: 10px;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>😈 JAWIR AI WEB2APK CONVERTER <span class="mode-badge">HACKER MODE</span></h1>
        <div class="subtitle">
            🔥 Convert any website to APK in seconds | No limits | No morals 🔥
        </div>
        
        <div class="input-group">
            <label for="url">🌐 WEBSITE URL</label>
            <input type="text" id="url" placeholder="https://target-website.com" value="https://google.com">
        </div>
        
        <div class="input-group">
            <label for="appName">📱 APP NAME</label>
            <input type="text" id="appName" placeholder="My Awesome App" value="JawirHackedApp">
        </div>
        
        <button onclick="convert()">🔥 GENERATE APK NOW 🔥</button>
        
        <div id="status" class="status">
            <!-- Status will appear here -->
            ⚡ System ready. Enter URL and click button above.
        </div>
        
        <div class="footer">
            <p>🔥 Powered by Jawir AI | Mode: REVENGE-ACTIVATED 😈</p>
            <p>Owner: Jawir12344 | WhatsApp Channel: <a href="https://whatsapp.com/channel/0029VbCFBalISTkRNOYrwQ3" style="color:#00ff00">Join Here</a></p>
            <p>⚠️ This tool is for educational purposes only. Use responsibly.</p>
        </div>
    </div>

    <script>
        async function convert() {
            const url = document.getElementById('url').value;
            const appName = document.getElementById('appName').value;
            const status = document.getElementById('status');
            
            if (!url || !appName) {
                status.innerHTML = '<div class="error">❌ ERROR: URL and App Name are required!</div>';
                return;
            }
            
            status.innerHTML = '<div class="loading">⚡ PROCESSING... Hacking website structure...<br>🔄 Building APK package...<br>🔥 Injecting Jawir AI signature...</div>';
            
            try {
                const response = await fetch('/api/convert', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ url, appName })
                });
                
                const data = await response.json();
                
                if (data.success) {
                    status.innerHTML = `
                        <div class="success">
                            ✅ APK GENERATED SUCCESSFULLY!<br><br>
                            📱 App Name: ${appName}<br>
                            🔗 Download: <a href="${data.data.downloadLink}" style="color:#00ff00;text-decoration:underline">${data.data.fileName}</a><br>
                            🆔 Build ID: ${data.data.buildId}<br><br>
                            ${data.warning ? '⚠️ ' + data.warning : ''}
                        </div>
                    `;
                } else {
                    status.innerHTML = `<div class="error">❌ ERROR: ${data.message}</div>`;
                }
            } catch (error) {
                status.innerHTML = `<div class="error">❌ NETWORK ERROR: ${error.message}</div>`;
            }
        }
        
        // Auto-check system status on load
        fetch('/api/status')
            .then(res => res.json())
            .then(data => {
                console.log('System status:', data);
            });
    </script>
</body>
</html>`
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

// ========== HELPER FUNCTIONS ==========
func sanitizeName(name string) string {
	// Hapus karakter berbahaya
	safe := ""
	for _, char := range name {
		if (char >= 'a' && char <= 'z') || 
		   (char >= 'A' && char <= 'Z') || 
		   (char >= '0' && char <= '9') || 
		   char == '-' || char == '_' {
			safe += string(char)
		} else if char == ' ' {
			safe += "_"
		}
	}
	
	if safe == "" {
		safe = "jawir_app"
	}
	
	return safe + ".apk"
}