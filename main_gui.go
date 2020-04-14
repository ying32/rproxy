// +build gui

package main

import (
	"github.com/ying32/govcl/vcl"
)

// GUI模式下暂时只能是客户端，服务端依然使用命令行。
// 另外如果更新了MainForm.go或者MainFormImpl.go需要在前面补充自定义标题 // +build gui
func main() {

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&MainForm)
	vcl.Application.Run()

}
