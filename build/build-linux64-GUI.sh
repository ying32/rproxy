echo ±‡“Î64Œªrproxy_GUI÷–...

export GOARCH=amd64
export GOOS=linux
export CGO_ENABLED=1
cd ../
go build -tags gui -o rproxy_GUI
pause