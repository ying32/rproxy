### 构建rproxy命令

如果要编译带界面的客户端需要加上 -tags gui 如下面的   
go build -i -ldflags="-H windowsgui" -tags gui -o rproxy_GUI.exe    

构建带UI的客户端/服务端需要用到https://github.com/ying32/govcl 这个库。  

自签HTTPS证书可使用go标准库中已经有的工具  
go\src\crypto\tls\generate_cert.go  

最后面下面下载对应的govcl二进制：  
https://github.com/ying32/govcl/releases/download/v1.2.2/Librarys-1.2.2.zip   

注：现rproxy中使用的二进制为govcl dev分支中的代码，相应的二进制需要自己编译。