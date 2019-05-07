echo ±‡“Î32Œªrproxy_GUI÷–...

export GOARCH=386
export GOOS=linux
export CGO_ENABLED=1
cd ../
go build -tags gui -o rproxy_GUI
pause