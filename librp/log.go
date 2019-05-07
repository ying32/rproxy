package librp

import (
	"fmt"
	"log"
	"os"
)

var (
	DEBUG          = true           // 调试模式 输出 Log.D 字符
	IsGUI          bool             // 是否为GUI
	LogGUICallback func(msg string) // 当IsGUI=true时，不再输出到日志，使用当前回调函数
)

type Logger struct {
}

var Log Logger

var std = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)

func (l Logger) println(calldepth int, flag string, v ...interface{}) {
	msg := fmt.Sprintf("[%s]: %s", flag, fmt.Sprint(v...))
	if IsGUI {
		if LogGUICallback != nil {
			LogGUICallback(msg)
		}
		return
	}
	std.Output(calldepth, msg)
}

// 警告
func (l Logger) W(v ...interface{}) {
	textYellow()
	l.println(3, "WARNING", v...)
}

// 错误
func (l Logger) E(v ...interface{}) {
	textRed()
	l.println(3, "ERROR", v...)
}

func (l Logger) EF(v ...interface{}) {
	textRed()
	l.println(3, "ERROR", v...)
	os.Exit(1)
}

// 信息
func (l Logger) I(v ...interface{}) {
	textDefault()
	l.println(3, "INFO", v...)
}

// 调试
func (l Logger) D(v ...interface{}) {
	if DEBUG {
		textGreen()
		l.println(3, "DEBUG", v...)
	}
}
