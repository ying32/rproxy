echo Compiling 32-bit rproxy GUI...

# 根据你的实际情况修改这里
export GOROOT=$HOME/godev/go
export GOPATH=$HOME/godev/gosrc
export PATH=$GOROOT/bin:$PATH

export GOARCH=386
export GOOS=darwin
export CGO_ENABLED=1
cd ../
go build -tags gui -o rproxy_GUI