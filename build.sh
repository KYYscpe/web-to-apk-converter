mkdir jawir-converter && cd jawir-converter
mkdir api

# Buat file convert.go
cat > api/convert.go << 'EOF'
[PASTE SEMUA CODE HANDLER DI ATAS DI SINI]
EOF

# Buat go.mod
echo 'module jawir-web2apk\ngo 1.21' > go.mod

# Buat vercel.json
cat > vercel.json << 'EOF'
{
  "builds": [
    {
      "src": "api/*.go",
      "use": "@vercel/go"
    }
  ],
  "routes": [
    {
      "src": "/(.*)",
      "dest": "/api/convert.go"
    }
  ]
}
EOF

# Deploy
vercel --prod