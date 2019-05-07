package librp

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"crypto/x509"
	"io"
	"io/ioutil"
	"net"
	"path/filepath"
	"strings"
)

var (
	// 全局配置文件
	conf *TRProxyConfig
)

// 设置配置文件
func SetConfig(cfg *TRProxyConfig) {
	conf = cfg
	conf.certPool = x509.NewCertPool()
	// 初始KEY
	conf.verifyVal = sha1.Sum([]byte("I AM A KEY:" + conf.VerifyKey))
	addRootCert()
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
