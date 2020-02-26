// +build gui

// 由res2go自动生成，不要编辑。
package main

import (
    "github.com/ying32/govcl/vcl"
)

type TMainForm struct {
    *vcl.TForm
    StatusBar1         *vcl.TStatusBar
    Panel1             *vcl.TPanel
    Panel2             *vcl.TPanel
    BtnStart           *vcl.TButton
    BtnStop            *vcl.TButton
    Label12            *vcl.TLabel
    PageControl1       *vcl.TPageControl
    TabSheet1          *vcl.TTabSheet
    RGMode             *vcl.TRadioGroup
    GBBase             *vcl.TGroupBox
    Label1             *vcl.TLabel
    Label2             *vcl.TLabel
    Label3             *vcl.TLabel
    Label7             *vcl.TLabel
    SpinTCPPort        *vcl.TSpinEdit
    SpinCliHTTPPort    *vcl.TSpinEdit
    ChkIsZip           *vcl.TCheckBox
    EditVerifyKey      *vcl.TEdit
    BtnRandKey         *vcl.TButton
    BtnSaveCfg         *vcl.TButton
    BtnLoadCfg         *vcl.TButton
    BtnNewCfg          *vcl.TButton
    EditSvrAddr        *vcl.TEdit
    ChkIsHttps         *vcl.TCheckBox
    Label8             *vcl.TLabel
    SpinSvrHTTPPort    *vcl.TSpinEdit
    GBTLS              *vcl.TGroupBox
    Panel3             *vcl.TPanel
    Label4             *vcl.TLabel
    EditTLSCAFile      *vcl.TEdit
    BtnCAOpen          *vcl.TButton
    PageControl2       *vcl.TPageControl
    TabSheet3          *vcl.TTabSheet
    Panel4             *vcl.TPanel
    Label6             *vcl.TLabel
    EditTLSCliKeyFile  *vcl.TEdit
    Label5             *vcl.TLabel
    EditTLSCliCertFile *vcl.TEdit
    BtnCliCertOpen     *vcl.TButton
    BtnCliKeyOpen      *vcl.TButton
    TabSheet4          *vcl.TTabSheet
    Panel5             *vcl.TPanel
    Label10            *vcl.TLabel
    EditTLSSvrCertFile *vcl.TEdit
    BtnSvrCertOpen     *vcl.TButton
    BtnSvrKeyOpen      *vcl.TButton
    EditTLSSvrKeyFile  *vcl.TEdit
    Label11            *vcl.TLabel
    GBAppSettings      *vcl.TGroupBox
    ChkAutoReconnect   *vcl.TCheckBox
    Label9             *vcl.TLabel
    SpinMaxLogLine     *vcl.TSpinEdit
    TabSheet2          *vcl.TTabSheet
    LstLogs            *vcl.TListBox
    DlgSaveCfg         *vcl.TSaveDialog
    DlgOpen            *vcl.TOpenDialog
    ActionList1        *vcl.TActionList
    ActStart           *vcl.TAction
    ActStop            *vcl.TAction
    TrayIcon1          *vcl.TTrayIcon

    //::private::
    TMainFormFields
}

var MainForm *TMainForm




// 以字节形式加载
// vcl.Application.CreateForm(&MainForm)

func NewMainForm(owner vcl.IComponent) (root *TMainForm)  {
    vcl.CreateResForm(owner, &root)
    return
}

var mainFormBytes = []byte("\x54\x50\x46\x30\x09\x54\x4D\x61\x69\x6E\x46\x6F\x72\x6D\x08\x4D\x61\x69\x6E\x46\x6F\x72\x6D\x04\x4C\x65\x66\x74\x03\x89\x03\x06\x48\x65\x69\x67\x68\x74\x03\x46\x02\x03\x54\x6F\x70\x03\x20\x01\x05\x57\x69\x64\x74\x68\x03\xF2\x01\x0B\x42\x6F\x72\x64\x65\x72\x49\x63\x6F\x6E\x73\x0B\x0C\x62\x69\x53\x79\x73\x74\x65\x6D\x4D\x65\x6E\x75\x0A\x62\x69\x4D\x69\x6E\x69\x6D\x69\x7A\x65\x00\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x08\x62\x73\x53\x69\x6E\x67\x6C\x65\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\x72\x70\x72\x6F\x78\x79\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x46\x02\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xF2\x01\x08\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x0E\x70\x6F\x53\x63\x72\x65\x65\x6E\x43\x65\x6E\x74\x65\x72\x0A\x4C\x43\x4C\x56\x65\x72\x73\x69\x6F\x6E\x06\x07\x32\x2E\x30\x2E\x36\x2E\x30\x00\x0A\x54\x53\x74\x61\x74\x75\x73\x42\x61\x72\x0A\x53\x74\x61\x74\x75\x73\x42\x61\x72\x31\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x03\x30\x02\x05\x57\x69\x64\x74\x68\x03\xF2\x01\x06\x50\x61\x6E\x65\x6C\x73\x0E\x00\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x00\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\x30\x02\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xF2\x01\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x0A\x42\x65\x76\x65\x6C\x4F\x75\x74\x65\x72\x07\x06\x62\x76\x4E\x6F\x6E\x65\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x30\x02\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xF2\x01\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x32\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x02\x22\x03\x54\x6F\x70\x03\x0E\x02\x05\x57\x69\x64\x74\x68\x03\xF2\x01\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x42\x6F\x74\x74\x6F\x6D\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x22\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xF2\x01\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x08\x42\x74\x6E\x53\x74\x61\x72\x74\x04\x4C\x65\x66\x74\x03\x46\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x04\x05\x57\x69\x64\x74\x68\x02\x4B\x06\x41\x63\x74\x69\x6F\x6E\x07\x08\x41\x63\x74\x53\x74\x61\x72\x74\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x74\x6E\x53\x74\x6F\x70\x04\x4C\x65\x66\x74\x03\x9A\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x04\x05\x57\x69\x64\x74\x68\x02\x4B\x06\x41\x63\x74\x69\x6F\x6E\x07\x07\x41\x63\x74\x53\x74\x6F\x70\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x07\x4C\x61\x62\x65\x6C\x31\x32\x04\x4C\x65\x66\x74\x02\x20\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x09\x05\x57\x69\x64\x74\x68\x03\x14\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x45\xE6\xB3\xA8\xEF\xBC\x9A\xE5\xA6\x82\xE6\x9C\x89\xE4\xBF\xAE\xE6\x94\xB9\xE9\x85\x8D\xE7\xBD\xAE\xEF\xBC\x8C\xE9\x9C\x80\xE8\xA6\x81\xE3\x80\x90\xE4\xBF\x9D\xE5\xAD\x98\xE9\x85\x8D\xE7\xBD\xAE\xE3\x80\x91\xE5\x90\x8E\xE6\x96\xB9\xE8\x83\xBD\xE7\x94\x9F\xE6\x95\x88\xE3\x80\x82\x0A\x46\x6F\x6E\x74\x2E\x43\x6F\x6C\x6F\x72\x07\x07\x63\x6C\x47\x72\x65\x65\x6E\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x00\x00\x00\x0C\x54\x50\x61\x67\x65\x43\x6F\x6E\x74\x72\x6F\x6C\x0C\x50\x61\x67\x65\x43\x6F\x6E\x74\x72\x6F\x6C\x31\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\x0E\x02\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xF2\x01\x0A\x41\x63\x74\x69\x76\x65\x50\x61\x67\x65\x07\x09\x54\x61\x62\x53\x68\x65\x65\x74\x31\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x49\x6E\x64\x65\x78\x02\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x08\x4F\x6E\x43\x68\x61\x6E\x67\x65\x07\x12\x50\x61\x67\x65\x43\x6F\x6E\x74\x72\x6F\x6C\x31\x43\x68\x61\x6E\x67\x65\x00\x09\x54\x54\x61\x62\x53\x68\x65\x65\x74\x09\x54\x61\x62\x53\x68\x65\x65\x74\x31\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x19\xE6\x9C\x8D\xE5\x8A\xA1\xE7\xAB\xAF\x2F\xE5\xAE\xA2\xE6\x88\xB7\xE7\xAB\xAF\xE8\xAE\xBE\xE7\xBD\xAE\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xF0\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xEA\x01\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x00\x0B\x54\x52\x61\x64\x69\x6F\x47\x72\x6F\x75\x70\x06\x52\x47\x4D\x6F\x64\x65\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x02\x3A\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xEA\x01\x05\x41\x6C\x69\x67\x6E\x07\x05\x61\x6C\x54\x6F\x70\x08\x41\x75\x74\x6F\x46\x69\x6C\x6C\x09\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE6\xA8\xA1\xE5\xBC\x8F\xE9\x80\x89\xE6\x8B\xA9\x1C\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x4C\x65\x66\x74\x52\x69\x67\x68\x74\x53\x70\x61\x63\x69\x6E\x67\x02\x06\x1D\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x45\x6E\x6C\x61\x72\x67\x65\x48\x6F\x72\x69\x7A\x6F\x6E\x74\x61\x6C\x07\x18\x63\x72\x73\x48\x6F\x6D\x6F\x67\x65\x6E\x6F\x75\x73\x43\x68\x69\x6C\x64\x52\x65\x73\x69\x7A\x65\x1B\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x45\x6E\x6C\x61\x72\x67\x65\x56\x65\x72\x74\x69\x63\x61\x6C\x07\x18\x63\x72\x73\x48\x6F\x6D\x6F\x67\x65\x6E\x6F\x75\x73\x43\x68\x69\x6C\x64\x52\x65\x73\x69\x7A\x65\x1C\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x53\x68\x72\x69\x6E\x6B\x48\x6F\x72\x69\x7A\x6F\x6E\x74\x61\x6C\x07\x0E\x63\x72\x73\x53\x63\x61\x6C\x65\x43\x68\x69\x6C\x64\x73\x1A\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x53\x68\x72\x69\x6E\x6B\x56\x65\x72\x74\x69\x63\x61\x6C\x07\x0E\x63\x72\x73\x53\x63\x61\x6C\x65\x43\x68\x69\x6C\x64\x73\x12\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x4C\x61\x79\x6F\x75\x74\x07\x1D\x63\x63\x6C\x4C\x65\x66\x74\x54\x6F\x52\x69\x67\x68\x74\x54\x68\x65\x6E\x54\x6F\x70\x54\x6F\x42\x6F\x74\x74\x6F\x6D\x1B\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x43\x6F\x6E\x74\x72\x6F\x6C\x73\x50\x65\x72\x4C\x69\x6E\x65\x02\x02\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x24\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xE6\x01\x07\x43\x6F\x6C\x75\x6D\x6E\x73\x02\x02\x09\x49\x74\x65\x6D\x49\x6E\x64\x65\x78\x02\x00\x0D\x49\x74\x65\x6D\x73\x2E\x53\x74\x72\x69\x6E\x67\x73\x01\x06\x09\xE5\xAE\xA2\xE6\x88\xB7\xE7\xAB\xAF\x06\x09\xE6\x9C\x8D\xE5\x8A\xA1\xE7\xAB\xAF\x00\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x09\x54\x47\x72\x6F\x75\x70\x42\x6F\x78\x06\x47\x42\x42\x61\x73\x65\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\xDD\x00\x03\x54\x6F\x70\x02\x6D\x05\x57\x69\x64\x74\x68\x03\xEA\x01\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x19\xE6\x9C\x8D\xE5\x8A\xA1\xE7\xAB\xAF\x2F\xE5\xAE\xA2\xE6\x88\xB7\xE7\xAB\xAF\xE8\xAE\xBE\xE7\xBD\xAE\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xC7\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xE6\x01\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x11\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x06\x05\x57\x69\x64\x74\x68\x03\xA2\x00\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x1E\xE8\xBF\x9E\xE6\x8E\xA5\xE6\x88\x96\xE7\x9B\x91\xE5\x90\xAC\xE7\x9A\x84\x54\x43\x50\xE7\xAB\xAF\xE5\x8F\xA3\xEF\xBC\x9A\x06\x4C\x61\x79\x6F\x75\x74\x07\x08\x74\x6C\x43\x65\x6E\x74\x65\x72\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x32\x04\x4C\x65\x66\x74\x02\x11\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x46\x05\x57\x69\x64\x74\x68\x03\xA2\x00\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x22\xE5\xAE\xA2\xE6\x88\xB7\xE7\xAB\xAF\xE8\xBD\xAC\xE5\x8F\x91\xE7\x9A\x84\x48\x54\x54\x50\x28\x53\x29\xE7\xAB\xAF\xE5\x8F\xA3\xEF\xBC\x9A\x06\x4C\x61\x79\x6F\x75\x74\x07\x08\x74\x6C\x43\x65\x6E\x74\x65\x72\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x33\x04\x4C\x65\x66\x74\x02\x13\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x03\x87\x00\x05\x57\x69\x64\x74\x68\x02\x46\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE9\xAA\x8C\xE8\xAF\x81\x4B\x45\x59\xEF\xBC\x9A\x06\x4C\x61\x79\x6F\x75\x74\x07\x08\x74\x6C\x43\x65\x6E\x74\x65\x72\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x37\x04\x4C\x65\x66\x74\x02\x13\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x66\x05\x57\x69\x64\x74\x68\x02\x47\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x12\xE6\x9C\x8D\xE5\x8A\xA1\xE5\x99\xA8\xE5\x9C\xB0\xE5\x9D\x80\xEF\xBC\x9A\x06\x4C\x61\x79\x6F\x75\x74\x07\x08\x74\x6C\x43\x65\x6E\x74\x65\x72\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x00\x00\x09\x54\x53\x70\x69\x6E\x45\x64\x69\x74\x0B\x53\x70\x69\x6E\x54\x43\x50\x50\x6F\x72\x74\x04\x4C\x65\x66\x74\x03\xCE\x00\x06\x48\x65\x69\x67\x68\x74\x02\x16\x04\x48\x69\x6E\x74\x06\x10\xE5\x8F\xAF\xE8\xBE\x93\xE5\x85\xA5\x31\x2D\x36\x35\x35\x33\x35\x03\x54\x6F\x70\x02\x06\x05\x57\x69\x64\x74\x68\x02\x4E\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x08\x4D\x61\x78\x56\x61\x6C\x75\x65\x04\xFF\xFF\x00\x00\x08\x4D\x69\x6E\x56\x61\x6C\x75\x65\x02\x01\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x05\x56\x61\x6C\x75\x65\x03\x5D\x20\x00\x00\x09\x54\x53\x70\x69\x6E\x45\x64\x69\x74\x0F\x53\x70\x69\x6E\x43\x6C\x69\x48\x54\x54\x50\x50\x6F\x72\x74\x04\x4C\x65\x66\x74\x03\xCE\x00\x06\x48\x65\x69\x67\x68\x74\x02\x16\x04\x48\x69\x6E\x74\x06\x10\xE5\x8F\xAF\xE8\xBE\x93\xE5\x85\xA5\x31\x2D\x36\x35\x35\x33\x35\x03\x54\x6F\x70\x02\x46\x05\x57\x69\x64\x74\x68\x02\x4E\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x08\x4D\x61\x78\x56\x61\x6C\x75\x65\x04\xFF\xFF\x00\x00\x08\x4D\x69\x6E\x56\x61\x6C\x75\x65\x02\x01\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x05\x56\x61\x6C\x75\x65\x03\x5E\x20\x00\x00\x09\x54\x43\x68\x65\x63\x6B\x42\x6F\x78\x08\x43\x68\x6B\x49\x73\x5A\x69\x70\x04\x4C\x65\x66\x74\x02\x13\x06\x48\x65\x69\x67\x68\x74\x02\x15\x03\x54\x6F\x70\x03\xAA\x00\x05\x57\x69\x64\x74\x68\x02\x58\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0F\xE5\xBC\x80\xE5\x90\xAF\x5A\x49\x50\xE5\x8E\x8B\xE7\xBC\xA9\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x05\x54\x45\x64\x69\x74\x0D\x45\x64\x69\x74\x56\x65\x72\x69\x66\x79\x4B\x65\x79\x04\x4C\x65\x66\x74\x02\x5C\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x03\x87\x00\x05\x57\x69\x64\x74\x68\x03\x95\x00\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0A\x42\x74\x6E\x52\x61\x6E\x64\x4B\x65\x79\x04\x4C\x65\x66\x74\x03\xF2\x00\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x86\x00\x05\x57\x69\x64\x74\x68\x02\x2B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE9\x9A\x8F\xE6\x9C\xBA\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x04\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0A\x42\x74\x6E\x53\x61\x76\x65\x43\x66\x67\x04\x4C\x65\x66\x74\x03\x93\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x3E\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE4\xBF\x9D\xE5\xAD\x98\xE9\x85\x8D\xE7\xBD\xAE\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0F\x42\x74\x6E\x53\x61\x76\x65\x43\x66\x67\x43\x6C\x69\x63\x6B\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x05\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0A\x42\x74\x6E\x4C\x6F\x61\x64\x43\x66\x67\x04\x4C\x65\x66\x74\x03\x93\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x06\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE8\xBD\xBD\xE5\x85\xA5\xE9\x85\x8D\xE7\xBD\xAE\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x06\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x09\x42\x74\x6E\x4E\x65\x77\x43\x66\x67\x04\x4C\x65\x66\x74\x03\x93\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x22\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE6\x96\xB0\xE5\xBB\xBA\xE9\x85\x8D\xE7\xBD\xAE\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x07\x00\x00\x05\x54\x45\x64\x69\x74\x0B\x45\x64\x69\x74\x53\x76\x72\x41\x64\x64\x72\x04\x4C\x65\x66\x74\x02\x5C\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x66\x05\x57\x69\x64\x74\x68\x03\xC1\x00\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x08\x00\x00\x09\x54\x43\x68\x65\x63\x6B\x42\x6F\x78\x0A\x43\x68\x6B\x49\x73\x48\x74\x74\x70\x73\x04\x4C\x65\x66\x74\x02\x7C\x06\x48\x65\x69\x67\x68\x74\x02\x15\x03\x54\x6F\x70\x03\xAA\x00\x05\x57\x69\x64\x74\x68\x02\x70\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x12\xE7\x9B\x91\xE5\x90\xAC\x2F\xE8\xBD\xAC\xE5\x8F\x91\x48\x54\x54\x50\x53\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x09\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x38\x04\x4C\x65\x66\x74\x02\x11\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x26\x05\x57\x69\x64\x74\x68\x03\xA2\x00\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x22\xE6\x9C\x8D\xE5\x8A\xA1\xE7\xAB\xAF\xE7\x9B\x91\xE5\x90\xAC\xE7\x9A\x84\x48\x54\x54\x50\x28\x53\x29\xE7\xAB\xAF\xE5\x8F\xA3\xEF\xBC\x9A\x06\x4C\x61\x79\x6F\x75\x74\x07\x08\x74\x6C\x43\x65\x6E\x74\x65\x72\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x00\x00\x09\x54\x53\x70\x69\x6E\x45\x64\x69\x74\x0F\x53\x70\x69\x6E\x53\x76\x72\x48\x54\x54\x50\x50\x6F\x72\x74\x04\x4C\x65\x66\x74\x03\xCE\x00\x06\x48\x65\x69\x67\x68\x74\x02\x16\x04\x48\x69\x6E\x74\x06\x10\xE5\x8F\xAF\xE8\xBE\x93\xE5\x85\xA5\x31\x2D\x36\x35\x35\x33\x35\x03\x54\x6F\x70\x02\x26\x05\x57\x69\x64\x74\x68\x02\x4E\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x08\x4D\x61\x78\x56\x61\x6C\x75\x65\x04\xFF\xFF\x00\x00\x08\x4D\x69\x6E\x56\x61\x6C\x75\x65\x02\x01\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x0A\x05\x56\x61\x6C\x75\x65\x03\x5E\x20\x00\x00\x00\x09\x54\x47\x72\x6F\x75\x70\x42\x6F\x78\x05\x47\x42\x54\x4C\x53\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\xA6\x00\x03\x54\x6F\x70\x03\x4A\x01\x05\x57\x69\x64\x74\x68\x03\xEA\x01\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x42\x6F\x74\x74\x6F\x6D\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0F\x54\x4C\x53\xE8\xAF\x81\xE4\xB9\xA6\xE8\xAE\xBE\xE7\xBD\xAE\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x90\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xE6\x01\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x33\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x02\x20\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xE6\x01\x05\x41\x6C\x69\x67\x6E\x07\x05\x61\x6C\x54\x6F\x70\x0A\x42\x65\x76\x65\x6C\x4F\x75\x74\x65\x72\x07\x06\x62\x76\x4E\x6F\x6E\x65\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x20\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xE6\x01\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x34\x04\x4C\x65\x66\x74\x02\x11\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x07\x05\x57\x69\x64\x74\x68\x02\x56\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x12\x54\x4C\x53\x20\x43\x41\xE6\xA0\xB9\xE8\xAF\x81\xE4\xB9\xA6\xEF\xBC\x9A\x06\x4C\x61\x79\x6F\x75\x74\x07\x08\x74\x6C\x43\x65\x6E\x74\x65\x72\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x00\x00\x05\x54\x45\x64\x69\x74\x0D\x45\x64\x69\x74\x54\x4C\x53\x43\x41\x46\x69\x6C\x65\x04\x4C\x65\x66\x74\x02\x70\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x03\x05\x57\x69\x64\x74\x68\x03\x4A\x01\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x09\x42\x74\x6E\x43\x41\x4F\x70\x65\x6E\x04\x4C\x65\x66\x74\x03\xBC\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x03\x05\x57\x69\x64\x74\x68\x02\x1B\x08\x42\x69\x64\x69\x4D\x6F\x64\x65\x07\x0D\x62\x64\x52\x69\x67\x68\x74\x54\x6F\x4C\x65\x66\x74\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\x2E\x2E\x2E\x0E\x50\x61\x72\x65\x6E\x74\x42\x69\x64\x69\x4D\x6F\x64\x65\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x00\x0C\x54\x50\x61\x67\x65\x43\x6F\x6E\x74\x72\x6F\x6C\x0C\x50\x61\x67\x65\x43\x6F\x6E\x74\x72\x6F\x6C\x32\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x02\x70\x03\x54\x6F\x70\x02\x20\x05\x57\x69\x64\x74\x68\x03\xE6\x01\x0A\x41\x63\x74\x69\x76\x65\x50\x61\x67\x65\x07\x09\x54\x61\x62\x53\x68\x65\x65\x74\x33\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x49\x6E\x64\x65\x78\x02\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x0B\x54\x61\x62\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x08\x74\x70\x42\x6F\x74\x74\x6F\x6D\x00\x09\x54\x54\x61\x62\x53\x68\x65\x65\x74\x09\x54\x61\x62\x53\x68\x65\x65\x74\x33\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE5\xAE\xA2\xE6\x88\xB7\xE7\xAB\xAF\x54\x4C\x53\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x52\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xDE\x01\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x34\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x02\x52\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xDE\x01\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x0A\x42\x65\x76\x65\x6C\x4F\x75\x74\x65\x72\x07\x06\x62\x76\x4E\x6F\x6E\x65\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x52\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xDE\x01\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x36\x04\x4C\x65\x66\x74\x02\x0D\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x33\x05\x57\x69\x64\x74\x68\x02\x56\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x10\x54\x4C\x53\x20\x4B\x65\x79\xE6\x96\x87\xE4\xBB\xB6\xEF\xBC\x9A\x06\x4C\x61\x79\x6F\x75\x74\x07\x08\x74\x6C\x43\x65\x6E\x74\x65\x72\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x00\x00\x05\x54\x45\x64\x69\x74\x11\x45\x64\x69\x74\x54\x4C\x53\x43\x6C\x69\x4B\x65\x79\x46\x69\x6C\x65\x04\x4C\x65\x66\x74\x02\x6C\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x31\x05\x57\x69\x64\x74\x68\x03\x4A\x01\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x35\x04\x4C\x65\x66\x74\x02\x0D\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x0F\x05\x57\x69\x64\x74\x68\x02\x51\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x11\x54\x4C\x53\x20\x43\x65\x72\x74\xE6\x96\x87\xE4\xBB\xB6\xEF\xBC\x9A\x06\x4C\x61\x79\x6F\x75\x74\x07\x08\x74\x6C\x43\x65\x6E\x74\x65\x72\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x00\x00\x05\x54\x45\x64\x69\x74\x12\x45\x64\x69\x74\x54\x4C\x53\x43\x6C\x69\x43\x65\x72\x74\x46\x69\x6C\x65\x04\x4C\x65\x66\x74\x02\x6C\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x0D\x05\x57\x69\x64\x74\x68\x03\x4A\x01\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0E\x42\x74\x6E\x43\x6C\x69\x43\x65\x72\x74\x4F\x70\x65\x6E\x04\x4C\x65\x66\x74\x03\xB8\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x0D\x05\x57\x69\x64\x74\x68\x02\x1B\x08\x42\x69\x64\x69\x4D\x6F\x64\x65\x07\x0D\x62\x64\x52\x69\x67\x68\x74\x54\x6F\x4C\x65\x66\x74\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\x2E\x2E\x2E\x0E\x50\x61\x72\x65\x6E\x74\x42\x69\x64\x69\x4D\x6F\x64\x65\x08\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x13\x42\x74\x6E\x43\x6C\x69\x43\x65\x72\x74\x4F\x70\x65\x6E\x43\x6C\x69\x63\x6B\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0D\x42\x74\x6E\x43\x6C\x69\x4B\x65\x79\x4F\x70\x65\x6E\x04\x4C\x65\x66\x74\x03\xB8\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x31\x05\x57\x69\x64\x74\x68\x02\x1B\x08\x42\x69\x64\x69\x4D\x6F\x64\x65\x07\x0D\x62\x64\x52\x69\x67\x68\x74\x54\x6F\x4C\x65\x66\x74\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\x2E\x2E\x2E\x0E\x50\x61\x72\x65\x6E\x74\x42\x69\x64\x69\x4D\x6F\x64\x65\x08\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x12\x42\x74\x6E\x43\x6C\x69\x4B\x65\x79\x4F\x70\x65\x6E\x43\x6C\x69\x63\x6B\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x00\x00\x09\x54\x54\x61\x62\x53\x68\x65\x65\x74\x09\x54\x61\x62\x53\x68\x65\x65\x74\x34\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE6\x9C\x8D\xE5\x8A\xA1\xE7\xAB\xAF\x54\x4C\x53\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x52\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xDE\x01\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x35\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x02\x52\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xDE\x01\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x0A\x42\x65\x76\x65\x6C\x4F\x75\x74\x65\x72\x07\x06\x62\x76\x4E\x6F\x6E\x65\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x52\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xDE\x01\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x07\x4C\x61\x62\x65\x6C\x31\x30\x04\x4C\x65\x66\x74\x02\x0D\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x0F\x05\x57\x69\x64\x74\x68\x02\x51\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x11\x54\x4C\x53\x20\x43\x65\x72\x74\xE6\x96\x87\xE4\xBB\xB6\xEF\xBC\x9A\x06\x4C\x61\x79\x6F\x75\x74\x07\x08\x74\x6C\x43\x65\x6E\x74\x65\x72\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x00\x00\x05\x54\x45\x64\x69\x74\x12\x45\x64\x69\x74\x54\x4C\x53\x53\x76\x72\x43\x65\x72\x74\x46\x69\x6C\x65\x04\x4C\x65\x66\x74\x02\x6C\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x0D\x05\x57\x69\x64\x74\x68\x03\x4A\x01\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0E\x42\x74\x6E\x53\x76\x72\x43\x65\x72\x74\x4F\x70\x65\x6E\x04\x4C\x65\x66\x74\x03\xB8\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x0D\x05\x57\x69\x64\x74\x68\x02\x1B\x08\x42\x69\x64\x69\x4D\x6F\x64\x65\x07\x0D\x62\x64\x52\x69\x67\x68\x74\x54\x6F\x4C\x65\x66\x74\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\x2E\x2E\x2E\x0E\x50\x61\x72\x65\x6E\x74\x42\x69\x64\x69\x4D\x6F\x64\x65\x08\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x13\x42\x74\x6E\x53\x76\x72\x43\x65\x72\x74\x4F\x70\x65\x6E\x43\x6C\x69\x63\x6B\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0D\x42\x74\x6E\x53\x76\x72\x4B\x65\x79\x4F\x70\x65\x6E\x04\x4C\x65\x66\x74\x03\xB8\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x31\x05\x57\x69\x64\x74\x68\x02\x1B\x08\x42\x69\x64\x69\x4D\x6F\x64\x65\x07\x0D\x62\x64\x52\x69\x67\x68\x74\x54\x6F\x4C\x65\x66\x74\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\x2E\x2E\x2E\x0E\x50\x61\x72\x65\x6E\x74\x42\x69\x64\x69\x4D\x6F\x64\x65\x08\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x12\x42\x74\x6E\x53\x76\x72\x4B\x65\x79\x4F\x70\x65\x6E\x43\x6C\x69\x63\x6B\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x05\x54\x45\x64\x69\x74\x11\x45\x64\x69\x74\x54\x4C\x53\x53\x76\x72\x4B\x65\x79\x46\x69\x6C\x65\x04\x4C\x65\x66\x74\x02\x6C\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x31\x05\x57\x69\x64\x74\x68\x03\x4A\x01\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x07\x4C\x61\x62\x65\x6C\x31\x31\x04\x4C\x65\x66\x74\x02\x0D\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x33\x05\x57\x69\x64\x74\x68\x02\x56\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x10\x54\x4C\x53\x20\x4B\x65\x79\xE6\x96\x87\xE4\xBB\xB6\xEF\xBC\x9A\x06\x4C\x61\x79\x6F\x75\x74\x07\x08\x74\x6C\x43\x65\x6E\x74\x65\x72\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x00\x00\x00\x00\x00\x00\x09\x54\x47\x72\x6F\x75\x70\x42\x6F\x78\x0D\x47\x42\x41\x70\x70\x53\x65\x74\x74\x69\x6E\x67\x73\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x02\x33\x03\x54\x6F\x70\x02\x3A\x05\x57\x69\x64\x74\x68\x03\xEA\x01\x05\x41\x6C\x69\x67\x6E\x07\x05\x61\x6C\x54\x6F\x70\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\x41\x50\x50\xE8\xAE\xBE\xE7\xBD\xAE\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x1D\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xE6\x01\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x09\x54\x43\x68\x65\x63\x6B\x42\x6F\x78\x10\x43\x68\x6B\x41\x75\x74\x6F\x52\x65\x63\x6F\x6E\x6E\x65\x63\x74\x04\x4C\x65\x66\x74\x02\x13\x06\x48\x65\x69\x67\x68\x74\x02\x15\x03\x54\x6F\x70\x02\x03\x05\x57\x69\x64\x74\x68\x03\xB2\x00\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x27\xE4\xB8\x8E\xE6\x9C\x8D\xE5\x8A\xA1\xE5\x99\xA8\xE8\xBF\x9E\xE6\x8E\xA5\xE6\x96\xAD\xE5\xBC\x80\xE5\x90\x8E\xE8\x87\xAA\xE5\x8A\xA8\xE9\x87\x8D\xE8\xBF\x9E\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x39\x04\x4C\x65\x66\x74\x03\xF3\x00\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x02\x5A\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x15\xE6\x9C\x80\xE5\xA4\xA7\xE6\x97\xA5\xE5\xBF\x97\xE8\xA1\x8C\xE6\x95\xB0\xEF\xBC\x9A\x06\x4C\x61\x79\x6F\x75\x74\x07\x08\x74\x6C\x43\x65\x6E\x74\x65\x72\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x00\x00\x09\x54\x53\x70\x69\x6E\x45\x64\x69\x74\x0E\x53\x70\x69\x6E\x4D\x61\x78\x4C\x6F\x67\x4C\x69\x6E\x65\x04\x4C\x65\x66\x74\x03\x46\x01\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\x86\x00\x08\x41\x75\x74\x6F\x53\x69\x7A\x65\x08\x08\x4D\x61\x78\x56\x61\x6C\x75\x65\x04\x40\x0D\x03\x00\x08\x4D\x69\x6E\x56\x61\x6C\x75\x65\x02\x01\x08\x4F\x6E\x43\x68\x61\x6E\x67\x65\x07\x14\x53\x70\x69\x6E\x4D\x61\x78\x4C\x6F\x67\x4C\x69\x6E\x65\x43\x68\x61\x6E\x67\x65\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x05\x56\x61\x6C\x75\x65\x03\x88\x13\x00\x00\x00\x00\x09\x54\x54\x61\x62\x53\x68\x65\x65\x74\x09\x54\x61\x62\x53\x68\x65\x65\x74\x32\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE6\x97\xA5\xE5\xBF\x97\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x6F\x02\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x66\x02\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x00\x08\x54\x4C\x69\x73\x74\x42\x6F\x78\x07\x4C\x73\x74\x4C\x6F\x67\x73\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\x6A\x02\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\x66\x02\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x14\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x42\x6F\x74\x74\x6F\x6D\x02\x04\x0B\x46\x6F\x6E\x74\x2E\x48\x65\x69\x67\x68\x74\x02\xF4\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x00\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x05\x53\x74\x79\x6C\x65\x07\x10\x6C\x62\x4F\x77\x6E\x65\x72\x44\x72\x61\x77\x46\x69\x78\x65\x64\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x00\x00\x00\x0B\x54\x53\x61\x76\x65\x44\x69\x61\x6C\x6F\x67\x0A\x44\x6C\x67\x53\x61\x76\x65\x43\x66\x67\x0A\x44\x65\x66\x61\x75\x6C\x74\x45\x78\x74\x06\x04\x2E\x63\x66\x67\x06\x46\x69\x6C\x74\x65\x72\x06\x12\xE9\x85\x8D\xE7\xBD\xAE\xE6\x96\x87\xE4\xBB\xB6\x7C\x2A\x2E\x63\x66\x67\x04\x6C\x65\x66\x74\x03\x8D\x01\x03\x74\x6F\x70\x03\x13\x01\x00\x00\x0B\x54\x4F\x70\x65\x6E\x44\x69\x61\x6C\x6F\x67\x07\x44\x6C\x67\x4F\x70\x65\x6E\x04\x6C\x65\x66\x74\x03\xB3\x01\x03\x74\x6F\x70\x03\x13\x01\x00\x00\x0B\x54\x41\x63\x74\x69\x6F\x6E\x4C\x69\x73\x74\x0B\x41\x63\x74\x69\x6F\x6E\x4C\x69\x73\x74\x31\x04\x6C\x65\x66\x74\x03\x66\x01\x03\x74\x6F\x70\x03\x13\x01\x00\x07\x54\x41\x63\x74\x69\x6F\x6E\x08\x41\x63\x74\x53\x74\x61\x72\x74\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x90\xAF\xE5\x8A\xA8\x00\x00\x07\x54\x41\x63\x74\x69\x6F\x6E\x07\x41\x63\x74\x53\x74\x6F\x70\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x81\x9C\xE6\xAD\xA2\x00\x00\x00\x09\x54\x54\x72\x61\x79\x49\x63\x6F\x6E\x09\x54\x72\x61\x79\x49\x63\x6F\x6E\x31\x07\x56\x69\x73\x69\x62\x6C\x65\x09\x04\x6C\x65\x66\x74\x03\x66\x01\x03\x74\x6F\x70\x03\xE6\x00\x00\x00\x00")

// 注册Form资源
var _ = vcl.RegisterFormResource(MainForm, &mainFormBytes)
