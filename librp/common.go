package librp

import (
	"crypto/sha1"
	"fmt"
	"log"
)

// 验证的值
var verifyVal [20]byte

// time.Now().Format("2006-01-02 15:04:05"),

// 日志打印，这里封装下，用于格式化
func LogPrintln(v ...interface{}) {
	log.Println(fmt.Sprintf("[INFO]: %s", fmt.Sprint(v...)))
}

// 失败时打印，会调用os.Exit
func LogFatalln(v ...interface{}) {
	log.Fatalln(fmt.Sprintf("[ERROR]: %s", fmt.Sprint(v...)))
}

// 初始验证key
func InitVerifyKey(key string) {
	verifyVal = sha1.Sum([]byte("I AM A KEY:" + key))
}
