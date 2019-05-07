@echo off
echo ±‡“Î32Œªrproxy_GUI.exe÷–...

set GOARCH=386
set GOOS=windows
set CGO_ENABLED=0
cd ../
go build -i -ldflags="-H windowsgui" -tags gui -o rproxy_GUI.exe
pause