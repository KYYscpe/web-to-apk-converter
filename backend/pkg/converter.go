package pkg

import (
    "archive/zip"
    "fmt"
    "io"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

type APKBuilder struct {
    URL      string
    AppName  string
    OutputDir string
}

func NewAPKBuilder(url, appName string) *APKBuilder {
    return &APKBuilder{
        URL:       url,
        AppName:   appName,
        OutputDir: "generated_apks",
    }
}

func (a *APKBuilder) Build() (string, error) {
    // Buat folder output
    os.MkdirAll(a.OutputDir, 0755)

    // 1. Download web resources
    webDir := filepath.Join(a.OutputDir, "web_content")
    os.MkdirAll(webDir, 0755)
    
    err := a.downloadWebPage(webDir)
    if err != nil {
        return "", fmt.Errorf("download failed: %v", err)
    }

    // 2. Create AndroidManifest.xml
    manifestContent := a.generateManifest()
    manifestPath := filepath.Join(webDir, "AndroidManifest.xml")
    os.WriteFile(manifestPath, []byte(manifestContent), 0644)

    // 3. Create config.json for PWA
    configContent := fmt.Sprintf(`{
        "name": "%s",
        "start_url": "%s",
        "display": "standalone",
        "background_color": "#000000",
        "theme_color": "#ff0033"
    }`, a.AppName, a.URL)
    configPath := filepath.Join(webDir, "config.json")
    os.WriteFile(configPath, []byte(configContent), 0644)

    // 4. Build APK using external tool (simulated)
    apkPath := filepath.Join(a.OutputDir, fmt.Sprintf("%s.apk", strings.ReplaceAll(a.AppName, " ", "_")))
    
    // Real implementation would call:
    // - pwa-builder-cli
    // - android-sdk build tools
    // - or bubblewrap
    
    // Simulate build process
    cmd := exec.Command("sh", "-c", fmt.Sprintf(
        "cd %s && echo 'Building APK for %s...' && sleep 2 && touch temp.apk",
        webDir, a.AppName,
    ))
    if err := cmd.Run(); err != nil {
        // Fallback: create dummy APK for demo
        a.createDummyAPK(apkPath)
    } else {
        // Move generated APK
        os.Rename(filepath.Join(webDir, "temp.apk"), apkPath)
    }

    return apkPath, nil
}

func (a *APKBuilder) downloadWebPage(dir string) error {
    // Create index.html with WebView wrapper
    htmlContent := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>%s</title>
    <style>
        body { margin: 0; padding: 0; }
        iframe { width: 100%%; height: 100vh; border: none; }
    </style>
</head>
<body>
    <iframe src="%s" allow="camera; microphone; geolocation"></iframe>
</body>
</html>`, a.AppName, a.URL)

    return os.WriteFile(filepath.Join(dir, "index.html"), []byte(htmlContent), 0644)
}

func (a *APKBuilder) generateManifest() string {
    return fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<manifest xmlns:android="http://schemas.android.com/apk/res/android"
    package="com.jawir.%s"
    android:versionCode="1"
    android:versionName="1.0">

    <uses-permission android:name="android.permission.INTERNET" />
    <uses-permission android:name="android.permission.ACCESS_NETWORK_STATE" />

    <application
        android:icon="@mipmap/ic_launcher"
        android:label="%s"
        android:theme="@android:style/Theme.DeviceDefault.NoActionBar">
        
        <activity android:name=".MainActivity"
            android:exported="true">
            <intent-filter>
                <action android:name="android.intent.action.MAIN" />
                <category android:name="android.intent.category.LAUNCHER" />
            </intent-filter>
        </activity>
    </application>
</manifest>`, strings.ToLower(strings.ReplaceAll(a.AppName, " ", "")), a.AppName)
}

func (a *APKBuilder) createDummyAPK(path string) error {
    // Create a dummy ZIP file (simulating APK)
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := zip.NewWriter(file)
    defer writer.Close()

    // Add dummy files
    files := map[string]string{
        "AndroidManifest.xml": a.generateManifest(),
        "res/drawable/icon.png": "[DUMMY_IMAGE_DATA]",
        "classes.dex": "[DUMMY_DEX]",
    }

    for name, content := range files {
        f, _ := writer.Create(name)
        f.Write([]byte(content))
    }

    return nil
}