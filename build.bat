cmd /v /c "set GOOS=js&& set GOARCH=wasm&& go build -o web/app.wasm"
echo "WASM built"
cmd /c "go build -o main.exe"