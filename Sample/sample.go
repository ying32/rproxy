package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("服务端请求数据了！")
		fmt.Println(r)
		w.Write([]byte("请求成功！"))
	})
	fmt.Println("监听HTTP 8085端口中...")
	if err := http.ListenAndServe(":8085", nil); err != nil {
		fmt.Println("服务器开启失败")
	}
}
