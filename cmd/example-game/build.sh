#!/bin/bash
GOOS=js GOARCH=wasm go build -o cmd/example-game/static/main.wasm ./cmd/example-game
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" cmd/example-game/static/