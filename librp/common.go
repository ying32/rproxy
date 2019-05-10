package librp

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"net"
	"path/filepath"
	"strings"
)

var (
	// 全局配置文件
	conf = initConfig()
)

const (
	CLIENT    = "client"
	SERVER    = "server"
	errorHTML = `
<!DOCTYPE html>
<html>
<head>
<title>(*>﹏<*)错误啦</title>
<style>
    body { width: 40em; margin: 0 auto; }
</style>
</head>
<body>
<h2>请求错误啦</h2>
<p>(*>﹏<*)您要访问的页面目前无法使用<br/>
要不等一会再试试吧。。。</p>
<p>本服务由<a href="https://github.com/ying32/rproxy">rproxy</a>所提供。</p>
</body>
</html>`
)

func initConfig() *TRProxyConfig {
	cfg := new(TRProxyConfig)
	if cfg.certPool == nil {
		cfg.certPool = x509.NewCertPool()
	}
	return cfg
}

// 设置配置文件
func SetConfig(cfg *TRProxyConfig) {

	// 复制字段
	conf.TCPPort = cfg.TCPPort
	conf.VerifyKey = cfg.VerifyKey
	conf.IsHTTPS = cfg.IsHTTPS
	conf.TLSCAFile = cfg.TLSCAFile
	conf.IsZIP = cfg.IsZIP

	// server
	conf.Server.TLSCertFile = cfg.Server.TLSCertFile
	conf.Server.TLSKeyFile = cfg.Server.TLSKeyFile
	conf.Server.HTTPPort = cfg.Server.HTTPPort
	// client
	conf.Client.SvrAddr = cfg.Client.SvrAddr
	conf.Client.HTTPPort = cfg.Client.HTTPPort
	conf.Client.TLSCertFile = cfg.Client.TLSCertFile
	conf.Client.TLSKeyFile = cfg.Client.TLSKeyFile

	if conf.Client.TLSCertFile != "" && conf.Client.TLSKeyFile != "" {
		var err error
		conf.cliCert, err = tls.LoadX509KeyPair(conf.Client.TLSCertFile, conf.Client.TLSKeyFile)
		if err != nil {
			Log.E(err)
		}
	} else {
		conf.cliCert = tls.Certificate{}
	}
	// 初始KEY
	conf.verifyVal = sha1.Sum([]byte("I AM A KEY:" + conf.VerifyKey))

	addRootCert()
}

func GetConfig() *TRProxyConfig {
	return conf
}

// 从 xxx.xxx.xxx.xxx:xxx格式中取出ip地址
func IPStr(conn net.Conn) string {
	if conn == nil {
		return ""
	}
	ip := conn.RemoteAddr().String()
	i := strings.LastIndex(ip, ":")
	if i == -1 {
		return ip
	}
	return ip[:i]
}

// ExtractFilePath 提取文件名路径
func ExtractFilePath(path string) string {
	filename := filepath.Base(path)
	return path[:len(path)-len(filename)]
}

// 添加CA根证书
func addRootCert() {
	if !conf.IsHTTPS {
		return
	}
	if conf.TLSCAFile == "" {
		Log.E("CA根证书文件不存在。")
	}
	bs, err := ioutil.ReadFile(conf.TLSCAFile)
	if err != nil {
		Log.E(err)
		return
	}
	if ok := conf.certPool.AppendCertsFromPEM(bs); !ok {
		Log.E("添加CA根证书失败。")
	}
}

//   zlib解压缩
func ZlibUnCompress(input []byte) ([]byte, error) {
	var out bytes.Buffer
	r, err := zlib.NewReader(bytes.NewReader(input))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	_, err = io.Copy(&out, r)
	if err != nil {
		return nil, nil
	}
	return out.Bytes(), nil
}

//  zlib压缩
func ZlibCompress(input []byte) ([]byte, error) {
	var in bytes.Buffer
	w, err := zlib.NewWriterLevel(&in, zlib.BestCompression)
	if err != nil {
		return nil, err
	}
	_, err = w.Write(input)
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return in.Bytes(), nil
}
