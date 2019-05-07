package librp

import (
	"crypto/sha1"
	"crypto/x509"
	"io/ioutil"
	"net"
	"strings"

	"github.com/pkg/errors"
)

var (
	// 验证的值
	verifyVal [20]byte

	// CA根证书池
	certPool    = x509.NewCertPool()
	tlsCertFile string
	tlsKeyFile  string
)

// 初始验证key
func InitVerifyKey(key string) {
	verifyVal = sha1.Sum([]byte("I AM A KEY:" + key))
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

func IIfStr(b bool, aTrue, aFalse string) string {
	if b {
		return aTrue
	}
	return aFalse
}

// 添加CA根证书
func AddRootCert(fileNmae string) error {
	if fileNmae == "" {
		return errors.New("CA根证书文件不存在。")
	}
	bs, err := ioutil.ReadFile(fileNmae)
	if err != nil {
		return err
	}
	if ok := certPool.AppendCertsFromPEM(bs); !ok {
		return errors.New("添加CA根证书失败。")
	}
	return nil
}

func SetTLSCertFile(fileName string) {
	tlsCertFile = fileName
}

func SetTLSKeyFile(fileName string) {
	tlsKeyFile = fileName
}
