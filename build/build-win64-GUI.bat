@echo off
echo ±‡“Î64Œªrproxy_GUI.exe÷–...

set GOARCH=amd64
set GOOS=windows
set CGO_ENABLED=0
cd ../
go build -i -ldflags="-H windowsgui" -tags gui -o rproxy_GUI.exe
pause