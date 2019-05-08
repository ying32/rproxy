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
    BtnCertOpen: TButton;
    BtnKeyOpen: TButton;
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
    procedure ActStartExecute(Sender: TObject);
    procedure ActStartUpdate(Sender: TObject);
    procedure ActStopExecute(Sender: TObject);
    procedure ActStopUpdate(Sender: TObject);
    procedure BtnCAOpenClick(Sender: TObject);
    procedure BtnCertOpenClick(Sender: TObject);
    procedure BtnKeyOpenClick(Sender: TObject);
    procedure BtnLoadCfgClick(Sender: TObject);
    procedure BtnNewCfgClick(Sender: TObject);
    procedure BtnRandKeyClick(Sender: TObject);
    procedure BtnSaveCfgClick(Sender: TObject);
    procedure ChkAutoReconnectClick(Sender: TObject);

    procedure GBBaseClick(Sender: TObject);

    procedure LstLogsDrawItem(Control: TWinControl; Index: Integer;
      ARect: TRect; State: TOwnerDrawState);
    procedure RGModeClick(Sender: TObject);
  private

  public

  end;

var
  MainForm: TMainForm;

implementation

{$R *.lfm}

{ TMainForm }





procedure TMainForm.RGModeClick(Sender: TObject);
begin
  //
end;

procedure TMainForm.BtnRandKeyClick(Sender: TObject);
begin
  //
end;

procedure TMainForm.BtnSaveCfgClick(Sender: TObject);
begin
  //
end;

procedure TMainForm.ChkAutoReconnectClick(Sender: TObject);
begin
  //
end;



procedure TMainForm.GBBaseClick(Sender: TObject);
begin

end;


procedure TMainForm.LstLogsDrawItem(Control: TWinControl; Index: Integer;
  ARect: TRect; State: TOwnerDrawState);
begin
  //
end;

procedure TMainForm.BtnCAOpenClick(Sender: TObject);
begin
  //
end;

procedure TMainForm.ActStartExecute(Sender: TObject);
begin
  //
end;

procedure TMainForm.ActStartUpdate(Sender: TObject);
begin
  //
end;

procedure TMainForm.ActStopExecute(Sender: TObject);
begin
  //
end;

procedure TMainForm.ActStopUpdate(Sender: TObject);
begin
  //
end;

procedure TMainForm.BtnCertOpenClick(Sender: TObject);
begin
  //
end;

procedure TMainForm.BtnKeyOpenClick(Sender: TObject);
begin
  //
end;

procedure TMainForm.BtnLoadCfgClick(Sender: TObject);
begin
  //
end;

procedure TMainForm.BtnNewCfgClick(Sender: TObject);
begin
  //
end;

procedure TMainForm.BtnCAOpenClick(Sender: TObject);
begin
  //
end;

procedure TMainForm.BtnNewCfgClick(Sender: TObject);
begin
  //
end;

end.

