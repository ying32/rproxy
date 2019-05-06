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
	go func() {
		fmt.Println("监听HTTPS 8089端口中...")
		err := http.ListenAndServeTLS(":8089", "./cert.pem", "./key.pem", nil)
		if err != nil {
			fmt.Println("开启HTTPS服务端失败：", err)
		}
	}()

	fmt.Println("监听HTTP 8085端口中...")
	if err := http.ListenAndServe(":8085", nil); err != nil {
		fmt.Println("开启HTTP服血端失败：", err)
	}
}
