@echo off
echo Compiling 32-bit rproxy GUI...

set GOARCH=386
set GOOS=windows
set CGO_ENABLED=0
cd ../
go build -i -ldflags="-H windowsgui" -tags gui -o rproxy_GUI.exe
pause