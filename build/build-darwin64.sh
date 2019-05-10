echo Compiling 64-bit rproxy...

# 根据你的实际情况修改这里
export GOROOT=$HOME/godev/go
export GOPATH=$HOME/godev/gosrc
export PATH=$GOROOT/bin:$PATH

export GOARCH=amd64
export GOOS=darwin
export CGO_ENABLED=0
cd ../
go build -o rproxy