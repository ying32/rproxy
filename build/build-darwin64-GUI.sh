echo ±‡“Î64Œªrproxy_GUI÷–...

export GOARCH=amd64
export GOOS=darwin
export CGO_ENABLED=1
cd ../
go build -tags gui -o rproxy_GUI
pause