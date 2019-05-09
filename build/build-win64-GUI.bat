@echo off
echo Compiling 64-bit rproxy GUI...

set GOARCH=amd64
set GOOS=windows
set CGO_ENABLED=0
cd ../
go build -i -ldflags="-H windowsgui" -tags gui -o rproxy_GUI.exe
pause