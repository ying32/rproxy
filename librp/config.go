package librp

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
)

// 配置文件，一个json文件，可包含客户端和服务端配置
type TRProxyConfig struct {
	// 不导出字段
	certPool  *x509.CertPool  // 这个字段不导出的，只内部使用
	cliCert   tls.Certificate // 这个字段不导出的，只内部使用
	verifyVal [20]byte        // 验证的值

	// 导出的字段
	TCPPort   int    `json:"tcpport"`   // Socket连接或者监听的端口
	VerifyKey string `json:"vkey"`      // 用作客户端与服务端连接时的校验
	IsHTTPS   bool   `json:"ishttps"`   // httpPort端口是否只用作HTTPS监听
	TLSCAFile string `json:"tlscafile"` // 当ishttps为true时，所需的CA根证书文件
	IsZIP     bool   `json:"iszip"`     // 是否开启zip压缩
	Server    struct {
		HTTPPort    int    `json:"httpport"`    // 当mode为server时为服务端监听端口，当为mode为client时为转发至本地客户端的端口
		TLSCertFile string `json:"tlscertfile"` // 当ishttps为true时，所需求的TLS证书文件
		TLSKeyFile  string `json:"tlskeyfile"`  // 当ishttps为true时，所需求的TLS密匙文件
	} `json:"server"`
	Client struct {
		SvrAddr     string `json:"svraddr"`     // 127.0.0.1", "当mode为client时有效，为连接服务器的地址
		HTTPPort    int    `json:"httpport"`    // 当mode为server时为服务端监听端口，当为mode为client时为转发至本地客户端的端口
		TLSCertFile string `json:"tlscertfile"` // 当ishttps为true时，所需求的TLS证书文件
		TLSKeyFile  string `json:"tlskeyfile"`  // 当ishttps为true时，所需求的TLS密匙文件
		LocalAddr   string `json:"localaddr"`   // 转发至本地的地址，默认为127.0.0.1
	} `json:"client"`
}

func LoadConfig(fileName string, cfg *TRProxyConfig) error {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, cfg)
}

func SaveConfig(fileName string, cfg *TRProxyConfig) error {
	bs, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, bs, 0666)
}
