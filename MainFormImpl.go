// +build gui

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"time"

	"errors"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"github.com/ying32/rproxy/librp"
)

const (
	randStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	SERVER  = 1
	CLIENT  = 0
)

var CheckConfigError = errors.New("检查配置错误")

//::private::
type TMainFormFields struct {
	rpObj            librp.IRPObject
	rpConfigFileName string
	started          bool
	autoReboot       bool
	appCfg           *vcl.TIniFile
	modeIndex        int32
	isDarwin         bool
	rpConfigLoaded   bool
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	rand.Seed(time.Now().Unix())
	f.isDarwin = runtime.GOOS == "darwin"

	f.ScreenCenter()
	f.PageControl1.SetActivePageIndex(0)
	f.PageControl2.SetActivePageIndex(0)

	f.started = false
	f.appCfg = vcl.NewIniFile(rtl.ExtractFilePath(vcl.Application.ExeName()) + "app.conf")

	if runtime.GOOS == "windows" {
		f.TrayIcon1.SetIcon(vcl.Application.Icon())
	} else {
		loadMainIconFromStream(f.TrayIcon1.Icon())
	}

	librp.IsGUI = true
	librp.LogGUICallback = f.logCallback
	f.loadAppConfig()

}

func loadMainIconFromStream(outIcon *vcl.TIcon) {
	if outIcon.IsValid() {
		mem := vcl.NewMemoryStreamFromBytes(mainIconBytes)
		defer mem.Free()
		mem.SetPosition(0)
		outIcon.LoadFromStream(mem)
	}
}

func (f *TMainForm) OnFormDestroy(sender vcl.IObject) {
	if f.appCfg != nil {
		f.appCfg.Free()
	}
	if f.rpObj != nil {
		f.rpObj.Close()
	}
}

func (f *TMainForm) loadAppConfig() {
	f.rpConfigFileName = f.appCfg.ReadString("System", "RPConfigFileName", "")
	cfg := new(librp.TRProxyConfig)
	err := librp.LoadConfig(f.rpConfigFileName, cfg)
	if err == nil {
		librp.SetConfig(cfg)
		f.updateUIConfig(cfg)
		f.rpConfigLoaded = true
	}
	f.autoReboot = f.appCfg.ReadBool("System", "AutoReboot", true)
	f.ChkAutoReconnect.SetChecked(f.autoReboot)

	f.modeIndex = f.appCfg.ReadInteger("System", "ModeIndex", 0)
	f.setRPMode(f.modeIndex)

	f.SpinMaxLogLine.SetValue(f.appCfg.ReadInteger("System", "MaxLogLine", 5000))

	f.updateCaption()
}

func (f *TMainForm) logCallback(msg string) {
	vcl.ThreadSync(func() {
		if f.LstLogs.Count() > f.SpinMaxLogLine.Value() {
			f.LstLogs.Clear()
		}
		f.LstLogs.Items().Add(msg)
	})
}

// 下面两个函数用于解决TRadioGroup在MacOS下bug问题
func (f *TMainForm) getRPMode() int32 {
	if f.isDarwin {
		var i int32
		for i = 0; i < f.RGMode.ControlCount(); i++ {
			ctl := f.RGMode.Controls(i)
			if strings.Compare(ctl.Name(), fmt.Sprintf("RadioButton%d", i)) == 0 {
				if vcl.RadioButtonFromObj(ctl).Checked() {
					return i
				}
			}
		}
		return -1
	} else {
		return f.RGMode.ItemIndex()
	}
}

func (f *TMainForm) setRPMode(idx int32) {
	if f.isDarwin {
		var i int32
		for i = 0; i < f.RGMode.ControlCount(); i++ {
			ctl := f.RGMode.Controls(i)
			if strings.Compare(ctl.Name(), fmt.Sprintf("RadioButton%d", i)) == 0 {
				if idx == i {
					vcl.RadioButtonFromObj(ctl).SetChecked(true)
				} else {
					vcl.RadioButtonFromObj(ctl).SetChecked(false)
				}
			}
		}
	} else {
		f.RGMode.SetItemIndex(idx)
	}
}

func (f *TMainForm) setTrayHint(text string) {
	if f.isDarwin {
		// darwin下出bug
		return
	}
	f.TrayIcon1.SetHint(text)
}

func (f *TMainForm) OnBtnRandKeyClick(sender vcl.IObject) {
	var randKey []byte
	for i := 0; i < 16; i++ {
		randKey = append(randKey, randStr[rand.Intn(len(randStr))])
	}
	f.EditVerifyKey.SetText(string(randKey))
}

func (f *TMainForm) OnBtnCAOpenClick(sender vcl.IObject) {
	if f.DlgOpen.Execute() {
		f.EditTLSCAFile.SetText(f.DlgOpen.FileName())
	}
}

func (f *TMainForm) OnBtnSaveCfgClick(sender vcl.IObject) {
	f.saveUIConfig()
}

func (f *TMainForm) updateUIConfig(cfg *librp.TRProxyConfig) {
	f.SpinTCPPort.SetValue(int32(cfg.TCPPort))
	f.EditVerifyKey.SetText(cfg.VerifyKey)
	f.ChkIsZip.SetChecked(cfg.IsZIP)
	f.ChkIsHttps.SetChecked(cfg.IsHTTPS)
	f.EditTLSCAFile.SetText(cfg.TLSCAFile)

	// client
	f.SpinCliHTTPPort.SetValue(int32(cfg.Client.HTTPPort))
	f.EditSvrAddr.SetText(cfg.Client.SvrAddr)
	f.EditTLSCliCertFile.SetText(cfg.Client.TLSCertFile)
	f.EditTLSCliKeyFile.SetText(cfg.Client.TLSKeyFile)

	// server
	f.SpinSvrHTTPPort.SetValue(int32(cfg.Server.HTTPPort))
	f.EditTLSSvrCertFile.SetText(cfg.Server.TLSCertFile)
	f.EditTLSSvrKeyFile.SetText(cfg.Server.TLSKeyFile)
}

func (f *TMainForm) checkConfig() error {
	if f.EditVerifyKey.Text() == "" {
		f.EditVerifyKey.SetFocus()
		vcl.ShowMessage("必须输入一个验证的key")
		return CheckConfigError
	}
	if f.EditSvrAddr.Text() == "" {
		f.EditSvrAddr.SetFocus()
		vcl.ShowMessage("要连接的服务器地址不能为空")
		return CheckConfigError
	}
	if f.ChkIsHttps.Checked() {

		if f.EditTLSCliCertFile.Text() == "" || f.EditTLSCliKeyFile.Text() == "" {
			vcl.ShowMessage("当为HTTPS时，客户端TLS证书不能为空。")
			return CheckConfigError
		}

		if f.EditTLSSvrCertFile.Text() == "" || f.EditTLSSvrKeyFile.Text() == "" {
			vcl.ShowMessage("当为HTTPS时，服务端TLS证书不能为空。")
			return CheckConfigError
		}
	}
	return nil
}

func (f *TMainForm) saveUIConfig() {
	if f.checkConfig() != nil {
		return
	}
	cfg := new(librp.TRProxyConfig)

	cfg.TCPPort = int(f.SpinTCPPort.Value())
	cfg.VerifyKey = f.EditVerifyKey.Text()
	cfg.IsZIP = f.ChkIsZip.Checked()
	cfg.IsHTTPS = f.ChkIsHttps.Checked()
	cfg.TLSCAFile = f.EditTLSCAFile.Text()

	// client
	cfg.Client.HTTPPort = int(f.SpinCliHTTPPort.Value())
	cfg.Client.SvrAddr = f.EditSvrAddr.Text()
	cfg.Client.TLSCertFile = f.EditTLSCliCertFile.Text()
	cfg.Client.TLSKeyFile = f.EditTLSCliKeyFile.Text()

	// server
	cfg.Server.HTTPPort = int(f.SpinSvrHTTPPort.Value())
	cfg.Server.TLSCertFile = f.EditTLSSvrCertFile.Text()
	cfg.Server.TLSKeyFile = f.EditTLSSvrKeyFile.Text()

	if !rtl.FileExists(f.rpConfigFileName) {
		if f.DlgSaveCfg.Execute() {
			f.rpConfigFileName = f.DlgSaveCfg.FileName()
		} else {
			librp.Log.I("取消保存配置")
			return
		}
	}
	librp.SetConfig(cfg)
	err := librp.SaveConfig(f.rpConfigFileName, cfg)
	if err != nil {
		librp.Log.I("配置保存失败")
		return
	}
	f.rpConfigLoaded = true
	librp.Log.I("配置已保存")
}

func (f *TMainForm) OnBtnLoadCfgClick(sender vcl.IObject) {
	if f.DlgOpen.Execute() {
		cfg := new(librp.TRProxyConfig)
		err := librp.LoadConfig(f.DlgOpen.FileName(), cfg)
		if err != nil {
			vcl.ShowMessage("载入配置失败：" + err.Error())
		} else {
			librp.SetConfig(cfg)
			f.rpConfigFileName = f.DlgOpen.FileName()
			f.updateUIConfig(cfg)
			// 载入即保存下当前的文件名
			f.appCfg.WriteString("System", "RPConfigFileName", f.rpConfigFileName)
			f.rpConfigLoaded = true
		}
	}
}

func (f *TMainForm) OnChkAutoReconnectClick(sender vcl.IObject) {
	f.autoReboot = f.ChkAutoReconnect.Checked()
	f.appCfg.WriteBool("System", "AutoReboot", f.autoReboot)
}

func (f *TMainForm) runSvr() {
	str := ""
	switch f.modeIndex {
	case CLIENT:
		s := fmt.Sprintln("客户端启动，连接服务器：", f.EditSvrAddr.Text(), "， 端口：", f.SpinTCPPort.Value())
		str += s
		librp.Log.I(s)
		if f.ChkIsHttps.Checked() {
			s = "转发至HTTP服务为HTTPS"
			str += s + "\r\n"
			librp.Log.I(s)
		}
		s = fmt.Sprintln("转发至本地HTTP(S)端口：", f.SpinCliHTTPPort.Value())
		str += s
		librp.Log.I(s)

	case SERVER:
		s := fmt.Sprintln("TCP服务端已启动，端口：", f.SpinTCPPort.Value())
		str += s
		librp.Log.I(s)
		if f.ChkIsHttps.Checked() {
			s = "当前HTTP服务为HTTPS"
			str += s + "\r\n"
			librp.Log.I(s)
		}
		s = fmt.Sprintln("HTTP(S)服务端已开启，端口：", f.SpinSvrHTTPPort.Value())
		str += s
		librp.Log.I(s)
	}
	f.setTrayHint(str)

	switch f.modeIndex {
	case CLIENT:
		librp.Log.I("连接服务端...")
		for f.started {
			err := f.rpObj.Start()
			if err != nil {
				librp.Log.E("连接服务器错误：", err)
				librp.Log.I("5秒后重新连接...")
				for i := 0; i < 5; i++ {
					if !f.started {
						break
					}
					time.Sleep(time.Second * 1)
				}
				if !f.autoReboot {
					break
				}
			}
		}
	case SERVER:
		if f.started {
			err := f.rpObj.Start()
			if err != nil {
				librp.Log.E("启动服务端错误：", err)
			}
		}
	}
	librp.Log.I("已停止")
	vcl.ThreadSync(func() {
		f.BtnStop.Click()
	})
}

func (f *TMainForm) OnActStartExecute(sender vcl.IObject) {

	//if f.checkConfig() != nil {
	//	return
	//}

	switch f.modeIndex {
	case CLIENT:
		f.rpObj = librp.NewRPClient()
	case SERVER:
		f.rpObj = librp.NewRPServer()
	}
	if f.rpObj == nil {
		vcl.ShowMessage("rproxy对象创建失败。")
		return
	}

	f.started = true

	f.LstLogs.Clear()

	go f.runSvr()

	f.setControlState(false)
}

func (f *TMainForm) setControlState(state bool) {
	f.ChkIsHttps.SetEnabled(state)
	f.RGMode.SetEnabled(state)

	var i int32
	for i = 0; i < f.GBBase.ControlCount(); i++ {
		f.GBBase.Controls(i).SetEnabled(state)
	}
	for i = 0; i < f.GBTLS.ControlCount(); i++ {
		f.GBTLS.Controls(i).SetEnabled(state)
	}
	for i = 0; i < f.GBAppSettings.ControlCount(); i++ {
		f.GBAppSettings.Controls(i).SetEnabled(state)
	}
}

func (f *TMainForm) OnActStartUpdate(sender vcl.IObject) {
	vcl.ActionFromObj(sender).SetEnabled(!f.started && f.rpConfigLoaded)
}

func (f *TMainForm) OnActStopExecute(sender vcl.IObject) {
	f.started = false
	f.setControlState(true)

	if f.rpObj != nil {
		f.rpObj.Close()
	}
	f.rpObj = nil
	f.updateCaption()
}

func (f *TMainForm) OnActStopUpdate(sender vcl.IObject) {
	vcl.ActionFromObj(sender).SetEnabled(f.started && f.rpConfigLoaded)
}

func (f *TMainForm) OnBtnNewCfgClick(sender vcl.IObject) {
	librp.Log.I("新建配置")
	f.rpConfigFileName = ""
	f.rpConfigLoaded = false
}

func (f *TMainForm) updateCaption() {
	if f.isDarwin {
		// 唉macOS下liblcl.dylib有内部bug
		m := []string{"客户端", "服务端"}
		f.SetCaption("rproxy[" + m[f.modeIndex] + "]")
	} else {
		f.SetCaption("rproxy[" + f.RGMode.Items().Strings(f.modeIndex) + "]")
	}
	f.setTrayHint(f.Caption())
}

func (f *TMainForm) OnRGModeClick(sender vcl.IObject) {
	f.modeIndex = f.getRPMode()
	f.appCfg.WriteInteger("System", "ModeIndex", f.modeIndex)
	f.updateCaption()
}

func (f *TMainForm) OnLstLogsDrawItem(control vcl.IWinControl, index int32, aRect types.TRect, state types.TOwnerDrawState) {
	canvas := f.LstLogs.Canvas()

	canvas.Font().SetColor(colors.ClBlack)
	s := f.LstLogs.Items().Strings(index)
	if strings.HasPrefix(s, "[ERROR]") {
		canvas.Font().SetColor(colors.ClRed)
	} else if strings.HasPrefix(s, "[DEBUG]") {
		canvas.Font().SetColor(colors.ClGreen)
	} else if strings.HasPrefix(s, "[WARNING]") {
		canvas.Font().SetColor(colors.ClYellow)
	}
	canvas.Brush().SetColor(colors.ClWhite)
	canvas.FillRect(aRect)
	if rtl.InSets(uint32(state), types.OdSelected) {
		canvas.Font().SetColor(colors.ClBlue)
	}
	canvas.TextOut(aRect.Left, aRect.Top, s)
}

func (f *TMainForm) OnBtnCliCertOpenClick(sender vcl.IObject) {
	if f.DlgOpen.Execute() {
		f.EditTLSCliCertFile.SetText(f.DlgOpen.FileName())
	}
}

func (f *TMainForm) OnBtnCliKeyOpenClick(sender vcl.IObject) {
	if f.DlgOpen.Execute() {
		f.EditTLSCliKeyFile.SetText(f.DlgOpen.FileName())
	}
}

func (f *TMainForm) OnBtnSvrCertOpenClick(sender vcl.IObject) {
	if f.DlgOpen.Execute() {
		f.EditTLSSvrCertFile.SetText(f.DlgOpen.FileName())
	}
}

func (f *TMainForm) OnBtnSvrKeyOpenClick(sender vcl.IObject) {
	if f.DlgOpen.Execute() {
		f.EditTLSSvrKeyFile.SetText(f.DlgOpen.FileName())
	}
}

func (f *TMainForm) OnSpinMaxLogLineChange(sender vcl.IObject) {
	f.appCfg.WriteInteger("System", "MaxLogLine", f.SpinMaxLogLine.Value())
}

func (f *TMainForm) OnPageControl1Change(sender vcl.IObject) {

}
