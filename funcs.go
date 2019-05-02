package main

import (
	"fmt"
	"log"
)

// time.Now().Format("2006-01-02 15:04:05"),

// 日志打印，这里封装下，用于格式化
func logPrintln(v ...interface{}) {
	log.Println(fmt.Sprintf("[INFO]: %s", fmt.Sprint(v...)))
}

// 失败时打印，会调用os.Exit
func logFatalln(v ...interface{}) {
	log.Fatalln(fmt.Sprintf("[ERROR]: %s", fmt.Sprint(v...)))
}
