# ReverseProxy
简单的反向代理用于内网穿透  

**特别注意，此工具只适合小文件类的访问测试，用来做做数据调试**  

最近周末闲来无事，想起了做下微信公共号的开发，但微信限制只能80端口的，自己用的城中村的那种宽带，共用一个公网，没办法自己用路由做端口映射。自己的服务器在腾讯云上，每次都要编译完后用ftp上传再进行调试，非常的浪费时间。 一时间又不知道上哪找一个符合我的这种要求的工具，就索性自己构思了下，整个工作流程大致为：   


> 外部请求自己服务器上的HTTP服务端 -> 将数据传递给Socket服务器 -> Socket服务器将数据发送至已连接的Socket客户端 -> Socket客户端收到数据 -> 使用http请求本地http服务端 -> 本地http服务端处理相关后返回 -> Socket客户端将返回的数据发送至Socket服务端 -> Socket服务端解析出数据后原路返回至外部请求的HTTP  
 

> 使用方法：go get github.com/ying32/ReverseProxy  
> 1、go build   
> 2、服务端运行runsvr.bat或者runsvr.sh    
> 3、客户端运行runcli.bat或者runcli.sh    

> 命令行说明：  
>  --tcpport    Socket连接或者监听的端口   
>  --httpport   当mode为server时为服务端监听端口，当为mode为client时为转发至本地客户端的端口  
>  --mode       启动模式，可选为client、server，默认为client  
>  --svraddr    当mode为client时有效，为连接服务器的地址，不需要填写端口  


##### 支持Windows、Linux、MacOSX等，无第三方依赖库。

v0.2 释出 https://github.com/ying32/ReverseProxy/releases/tag/v0.2
