echo 编译64位rproxy_GUI中...

# 根据你的实际情况修改这里
export GOROOT=$HOME/go
export PATH=$GOROOT/bin:$PATH

export GOARCH=amd64
export GOOS=darwin
export CGO_ENABLED=1
cd ../
go build -tags gui -o rproxy_GUI