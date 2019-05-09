@echo off
echo Compiling 32-bit rproxy...

set GOARCH=386
set GOOS=windows
set CGO_ENABLED=0
cd ../
go build -o rproxy.exe
pause