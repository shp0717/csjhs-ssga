#!/bin/bash

# Build the Go server
echo "Building the server..."

GOOS=linux GOARCH=amd64 go build -o bin/server-linux-amd64 . && echo -e "\033[32mServer built successfully for Linux AMD64\033[0m" || echo -e "\033[31mFailed to build server for Linux AMD64\033[0m"
GOOS=darwin GOARCH=amd64 go build -o bin/server-darwin-amd64 . && echo -e "\033[32mServer built successfully for macOS AMD64\033[0m" || echo -e "\033[31mFailed to build server for macOS AMD64\033[0m"
GOOS=windows GOARCH=amd64 go build -o bin/server-windows-amd64.exe . && echo -e "\033[32mServer built successfully for Windows AMD64\033[0m" || echo -e "\033[31mFailed to build server for Windows AMD64\033[0m"

GOOS=linux GOARCH=arm64 go build -o bin/server-linux-arm64 . && echo -e "\033[32mServer built successfully for Linux ARM64\033[0m" || echo -e "\033[31mFailed to build server for Linux ARM64\033[0m"
GOOS=darwin GOARCH=arm64 go build -o bin/server-darwin-arm64 . && echo -e "\033[32mServer built successfully for macOS ARM64\033[0m" || echo -e "\033[31mFailed to build server for macOS ARM64\033[0m"
GOOS=windows GOARCH=arm64 go build -o bin/server-windows-arm64.exe . && echo -e "\033[32mServer built successfully for Windows ARM64\033[0m" || echo -e "\033[31mFailed to build server for Windows ARM64\033[0m"

GOOS=linux GOARCH=arm go build -o bin/server-linux-arm . && echo -e "\033[32mServer built successfully for Linux ARM\033[0m" || echo -e "\033[31mFailed to build server for Linux ARM\033[0m"
GOOS=windows GOARCH=386 go build -o bin/server-windows-386.exe . && echo -e "\033[32mServer built successfully for Windows 386\033[0m" || echo -e "\033[31mFailed to build server for Windows 386\033[0m"

GOOS=ios GOARCH=arm64 CGO_ENABLED=1 go build -o bin/server-ios-arm64 . && echo -e "\033[32mServer built successfully for iOS ARM64\033[0m" || echo -e "\033[31mFailed to build server for iOS ARM64\033[0m"
GOOS=android GOARCH=arm64 go build -o bin/server-android-arm64 . && echo -e "\033[32mServer built successfully for Android ARM64\033[0m" || echo -e "\033[31mFailed to build server for Android ARM64\033[0m"

echo "Build process completed."
