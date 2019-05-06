package librp

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
}

var Log Logger

var std = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)

func (l Logger) println(calldepth int, flag string, v ...interface{}) {
	std.Output(calldepth, fmt.Sprintf("[%s]: %s", flag, fmt.Sprint(v...)))
}

// 警告
func (l Logger) W(v ...interface{}) {
	l.println(3, "WARNING", v...)
}

// 错误
func (l Logger) E(v ...interface{}) {
	l.println(3, "ERROR", v...)
}

func (l Logger) EF(v ...interface{}) {
	l.println(3, "ERROR", v...)
	os.Exit(1)
}

// 信息
func (l Logger) I(v ...interface{}) {
	l.println(3, "INFO", v...)
}
