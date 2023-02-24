cmd /v /c "set GOOS=js&& set GOARCH=wasm&& go build cmd/main.go -o web/app.wasm"
echo "WASM built"
cmd /c "go build -o cmd/main.go main.exe"