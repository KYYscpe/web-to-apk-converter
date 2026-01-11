#!/bin/bash

echo "🔥 JAWIR's APK BUILDER SCRIPT 🔥"
echo "================================"

# Check dependencies
command -v go >/dev/null 2>&1 || { 
    echo "❌ Go not installed!" 
    exit 1 
}

# Build backend
echo "📦 Building Go backend..."
cd backend
go mod download
go build -o ../web-to-apk-converter main.go
if [ $? -eq 0 ]; then
    echo "✅ Backend built successfully!"
else
    echo "❌ Backend build failed!"
    exit 1
fi

# Install frontend dependencies (if any)
echo "🌐 Setting up frontend..."
cd ../frontend
# Add npm install if needed

# Create production build
echo "🚀 Creating production build..."
cd ..
mkdir -p dist
cp -r frontend/* dist/
cp backend/pkg/converter.go dist/ 2>/dev/null || true

# Set permissions
chmod +x web-to-apk-converter

echo ""
echo "========================================"
echo "✅ BUILD COMPLETE!"
echo "➤ Run: ./web-to-apk-converter"
echo "➤ Or deploy to Vercel: vercel --prod"
echo "========================================"
echo "😈 JAWIR AI - MODE HACKER AKTIF! 🔥"