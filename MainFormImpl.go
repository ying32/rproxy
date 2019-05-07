// +build gui

package main

import (
	"math/rand"
	"time"

	"github.com/ying32/govcl/vcl/rtl"

	"github.com/ying32/rproxy/librp"

	"github.com/ying32/govcl/vcl"
)

const randStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

//::private::
type TMainFormFields struct {
	rpClient         *librp.TRPClient
	rpConfigFileName string
	started          bool
	autoReboot       bool
	appCfg           *vcl.TIniFile
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
	rand.Seed(time.Now().Unix())

	f.started = false
	f.appCfg = vcl.NewIniFile(rtl.ExtractFilePath(vcl.Application.ExeName()) + "app.conf")

	librp.IsGUI = true
	librp.LogGUICallback = f.logCallback
	f.loadAppConfig()
	// 新建客户端
	f.rpClient = librp.NewRPClient()
}

func (f *TMainForm) OnFormDestroy(sender vcl.IObject) {
	if f.appCfg != nil {
		f.appCfg.Free()
	}
	if f.rpClient != nil {
		f.rpClient.Close()
	}
}

func (f *TMainForm) loadAppConfig() {
	f.rpConfigFileName = f.appCfg.ReadString("System", "RPConfigFileName", "")
	cfg := new(librp.TRProxyConfig)
	err := librp.LoadConfig(f.rpConfigFileName, cfg)
	if err == nil {
		librp.SetConfig(cfg)
		f.updateUIConfig(cfg)
	}
	f.autoReboot = f.appCfg.ReadBool("System", "AutoReboot", true)
	f.ChkAutoReboot.SetChecked(f.autoReboot)
}

func (f *TMainForm) logCallback(msg string) {
	vcl.ThreadSync(func() {
		f.StatusBar1.SetSimpleText(msg)
	})
}

func (f *TMainForm) OnBtnRandKeyClick(sender vcl.IObject) {
	var randKey []byte
	for i := 0; i < 16; i++ {
		randKey = append(randKey, randStr[rand.Intn(len(randStr))])
	}
	f.EditVerifyKey.SetText(string(randKey))
}

func (f *TMainForm) OnBtnKeyOpenClick(sender vcl.IObject) {
	if f.DlgOpen.Execute() {
		f.EditTLSKeyFile.SetText(f.DlgOpen.FileName())
	}
}

func (f *TMainForm) OnBtnCAOpenClick(sender vcl.IObject) {
	if f.DlgOpen.Execute() {
		f.EditTLSCAFile.SetText(f.DlgOpen.FileName())
	}
}

func (f *TMainForm) OnBtnCertOpenClick(sender vcl.IObject) {
	if f.DlgOpen.Execute() {
		f.EditTLSCertFile.SetText(f.DlgOpen.FileName())
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
	f.EditTLSCertFile.SetText(cfg.TLSCertFile)
	f.EditTLSKeyFile.SetText(cfg.TLSKeyFile)

	f.SpinHTTPPort.SetValue(int32(cfg.Client.HTTPPort))
	f.EditSvrAddr.SetText(cfg.Client.SvrAddr)

}

func (f *TMainForm) saveUIConfig() {
	cfg := new(librp.TRProxyConfig)
	// 获取服务端的
	cfg.Server = librp.GetConfig().Server

	cfg.TCPPort = int(f.SpinTCPPort.Value())
	cfg.VerifyKey = f.EditVerifyKey.Text()
	cfg.IsZIP = f.ChkIsZip.Checked()
	cfg.IsHTTPS = f.ChkIsHttps.Checked()
	cfg.TLSCAFile = f.EditTLSCAFile.Text()
	cfg.TLSCertFile = f.EditTLSCertFile.Text()
	cfg.TLSKeyFile = f.EditTLSKeyFile.Text()

	cfg.Client.HTTPPort = int(f.SpinHTTPPort.Value())
	cfg.Client.SvrAddr = f.EditSvrAddr.Text()

	if !rtl.FileExists(f.rpConfigFileName) {
		if f.DlgSaveCfg.Execute() {
			f.rpConfigFileName = f.DlgSaveCfg.FileName()
		} else {
			librp.Log.I("取消保存配置")
			return
		}
	}
	librp.SetConfig(cfg)
	librp.SaveConfig(f.rpConfigFileName, cfg)
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
		}
	}
}

func (f *TMainForm) OnChkAutoRebootClick(sender vcl.IObject) {
	f.autoReboot = f.ChkAutoReboot.Checked()
	f.appCfg.WriteBool("System", "AutoReboot", f.autoReboot)
}

func (f *TMainForm) OnActStartExecute(sender vcl.IObject) {

	f.started = true
	go func() {
		for f.started {
			librp.Log.I("连接服务端...")
			err := f.rpClient.Start()
			if err != nil {
				librp.Log.I("5秒后重新连接...")
				time.Sleep(time.Second * 5)
				if !f.autoReboot {
					break
				}
			}
		}
		vcl.ThreadSync(func() {
			f.BtnStop.Click()
		})
	}()

	f.setControlState(false)
}

func (f *TMainForm) setControlState(state bool) {
	f.ChkIsHttps.SetEnabled(state)
	var i int32
	for i = 0; i < f.GPBase.ControlCount(); i++ {
		f.GPBase.Controls(i).SetEnabled(state)
	}
	for i = 0; i < f.GPTLS.ControlCount(); i++ {
		f.GPTLS.Controls(i).SetEnabled(state)
	}
}

func (f *TMainForm) OnActStartUpdate(sender vcl.IObject) {
	vcl.ActionFromObj(sender).SetEnabled(!f.started)
}

func (f *TMainForm) OnActStopExecute(sender vcl.IObject) {
	f.started = false
	f.setControlState(true)

	f.rpClient.Close()
	librp.Log.I("已停止")
}

func (f *TMainForm) OnActStopUpdate(sender vcl.IObject) {
	vcl.ActionFromObj(sender).SetEnabled(f.started)
}

func (f *TMainForm) OnBtnNewCfgClick(sender vcl.IObject) {
	librp.Log.I("新建配置")
	f.rpConfigFileName = ""
}
