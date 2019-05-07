package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
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
		certPool := x509.NewCertPool()
		caCertFileName := "../cert/ca.pem"
		caCrt, err := ioutil.ReadFile(caCertFileName)
		if err != nil {
			fmt.Println("读TLS根证书错误", err)
			return
		}
		ok := certPool.AppendCertsFromPEM(caCrt)
		if !ok {
			fmt.Println("CA根证书添加失败。")
			return
		}
		httpsSvr := &http.Server{
			Addr:    ":8089",
			Handler: nil,
			TLSConfig: &tls.Config{
				ClientCAs: certPool,
				//ClientAuth: tls.RequireAndVerifyClientCert,
			},
		}
		err = httpsSvr.ListenAndServeTLS("../cert/server.pem", "../cert/server.key")
		if err != nil {
			fmt.Println("开启HTTPS服务端失败：", err)
		}
	}()

	fmt.Println("监听HTTP 8085端口中...")
	if err := http.ListenAndServe(":8085", nil); err != nil {
		fmt.Println("开启HTTP服血端失败：", err)
	}
}
