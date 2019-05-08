unit uMain;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, ActnList, ComCtrls,
  ExtCtrls, StdCtrls, Spin, Buttons, Types;

type

  { TMainForm }

  TMainForm = class(TForm)
    ActionList1: TActionList;
    ActStart: TAction;
    ActStop: TAction;
    BtnCertOpen: TButton;
    BtnKeyOpen: TButton;
    BtnLoadCfg: TButton;
    BtnNewCfg: TButton;
    BtnRandKey: TButton;
    BtnSaveCfg: TButton;
    BtnStart: TButton;
    BtnStop: TButton;
    BtnCAOpen: TButton;
    ChkAutoReconnect: TCheckBox;
    ChkIsHttps: TCheckBox;
    ChkIsZip: TCheckBox;
    DlgOpen: TOpenDialog;
    DlgSaveCfg: TSaveDialog;
    EditSvrAddr: TEdit;
    EditTLSCAFile: TEdit;
    EditTLSCertFile: TEdit;
    EditTLSKeyFile: TEdit;
    EditVerifyKey: TEdit;
    GBBase: TGroupBox;
    GBTLS: TGroupBox;
    GBAppSettings: TGroupBox;
    Label1: TLabel;
    Label2: TLabel;
    Label3: TLabel;
    Label4: TLabel;
    Label5: TLabel;
    Label6: TLabel;
    Label7: TLabel;
    Label8: TLabel;
    Label9: TLabel;
    LstLogs: TListBox;
    PageControl1: TPageControl;
    Panel1: TPanel;
    Panel2: TPanel;
    RGMode: TRadioGroup;
    SpinCliHTTPPort: TSpinEdit;
    SpinMaxLogLine: TSpinEdit;
    SpinSvrHTTPPort: TSpinEdit;
    SpinTCPPort: TSpinEdit;
    StatusBar1: TStatusBar;
    TabSheet1: TTabSheet;
    TabSheet2: TTabSheet;
    TrayIcon1: TTrayIcon;
  private

  public

  end;

var
  MainForm: TMainForm;

implementation

{$R *.lfm}

end.

