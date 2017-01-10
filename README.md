# ReverseProxy
简单的反向代理用于内网穿透

最近周末闲来无事，想起了做下微信公共号的开发，但微信限制只能80端口的，自己用的城中村的那种宽带，共用一个公网，没办法自己用路由做端口映射。自己的服务器在腾讯云上，每次都要编译完后用ftp上传再进行调试，非常的浪费时间。 一时间又不知道上哪找一个符合我的这种要求的工具，就索性自己构思了下，整个工作流程大致为：  

> 外部请求自己服务器上的HTTP服务端 -> 将数据传递给Socket服务器 -> Socket服务器将数据发送至已连接的Socket客户端 -> Socket客户端收到数据 -> 使用http请求本地http服务端 -> 本地http服务端处理相关后返回 -> Socket客户端将返回的数据发送至Socket服务端 -> Socket服务端解析出数据后原路返回至外部请求的HTTP  

**不过目前不支持Cookie与Header的设置，只保留了Method, URL, Body段的数据**  

> 使用方法：go get github.com/ying32/ReverseProxy
> 1、分别编译**RPServer**、 **RPClient**。  
> 2、在RPServer目录下的runsvr.bat或者runsvr.sh修改为你想要的端口后启动并执行。  
> 3、在RPClient目录下的runcli_local.bat或者runcli_local.sh修改为你想要的端口后启动并执行。      

**两个程序总共不到300行的代码量。简单是简单了点，但还是能用用的**  

##### 支持Win32, Win64, Linux32, Linux64, MacOSX32, MacOSX64等，无第三方依赖库。

v0.1 释出 https://github.com/ying32/ReverseProxy/releases/tag/v0.1
