@echo off
set GOARCH=amd64
set GOOS=linux
set CGO_ENABLED=0
go mod tidy
echo start building

go build -o ../bin/foundation ../main.go

echo build foundation success
pause
