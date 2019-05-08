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
    BtnCAOpen: TButton;
    BtnCliCertOpen: TButton;
    BtnSvrCertOpen: TButton;
    BtnCliKeyOpen: TButton;
    BtnSvrKeyOpen: TButton;
    BtnLoadCfg: TButton;
    BtnNewCfg: TButton;
    BtnRandKey: TButton;
    BtnSaveCfg: TButton;
    BtnStart: TButton;
    BtnStop: TButton;
    ChkAutoReconnect: TCheckBox;
    ChkIsHttps: TCheckBox;
    ChkIsZip: TCheckBox;
    DlgOpen: TOpenDialog;
    DlgSaveCfg: TSaveDialog;
    EditSvrAddr: TEdit;
    EditTLSCAFile: TEdit;
    EditTLSCliCertFile: TEdit;
    EditTLSSvrCertFile: TEdit;
    EditTLSCliKeyFile: TEdit;
    EditTLSSvrKeyFile: TEdit;
    EditVerifyKey: TEdit;
    GBBase: TGroupBox;
    GBTLS: TGroupBox;
    GBAppSettings: TGroupBox;
    Label1: TLabel;
    Label10: TLabel;
    Label11: TLabel;
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
    PageControl2: TPageControl;
    Panel1: TPanel;
    Panel2: TPanel;
    Panel3: TPanel;
    Panel4: TPanel;
    Panel5: TPanel;
    RGMode: TRadioGroup;
    SpinCliHTTPPort: TSpinEdit;
    SpinMaxLogLine: TSpinEdit;
    SpinSvrHTTPPort: TSpinEdit;
    SpinTCPPort: TSpinEdit;
    StatusBar1: TStatusBar;
    TabSheet1: TTabSheet;
    TabSheet2: TTabSheet;
    TabSheet3: TTabSheet;
    TabSheet4: TTabSheet;
    TrayIcon1: TTrayIcon;
    procedure BtnCliCertOpenClick(Sender: TObject);
    procedure BtnCliKeyOpenClick(Sender: TObject);
    procedure BtnSvrCertOpenClick(Sender: TObject);
    procedure BtnSvrKeyOpenClick(Sender: TObject);
    procedure PageControl1Change(Sender: TObject);
    procedure SpinMaxLogLineChange(Sender: TObject);
  private

  public

  end;

var
  MainForm: TMainForm;

implementation

{$R *.lfm}

{ TMainForm }

procedure TMainForm.PageControl1Change(Sender: TObject);
begin

end;

procedure TMainForm.SpinMaxLogLineChange(Sender: TObject);
begin

end;

procedure TMainForm.BtnCliCertOpenClick(Sender: TObject);
begin

end;

procedure TMainForm.BtnCliKeyOpenClick(Sender: TObject);
begin

end;

procedure TMainForm.BtnSvrCertOpenClick(Sender: TObject);
begin

end;

procedure TMainForm.BtnSvrKeyOpenClick(Sender: TObject);
begin

end;

end.

