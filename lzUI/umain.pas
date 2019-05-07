unit uMain;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, ActnList, ComCtrls,
  ExtCtrls, StdCtrls, Spin, Buttons;

type

  { TMainForm }

  TMainForm = class(TForm)
    ActionList1: TActionList;
    ActStart: TAction;
    ActStop: TAction;
    BtnCAOpen: TSpeedButton;
    BtnCertOpen: TSpeedButton;
    BtnKeyOpen: TSpeedButton;
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
    GPBase: TGroupBox;
    GPTLS: TGroupBox;
    Label1: TLabel;
    Label2: TLabel;
    Label3: TLabel;
    Label4: TLabel;
    Label5: TLabel;
    Label6: TLabel;
    Label7: TLabel;
    Panel1: TPanel;
    Panel2: TPanel;
    SpinHTTPPort: TSpinEdit;
    SpinTCPPort: TSpinEdit;
    StatusBar1: TStatusBar;
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
    procedure ChkAutoReconnectChange(Sender: TObject);
    procedure ChkAutoReconnectClick(Sender: TObject);
    procedure GPTLSClick(Sender: TObject);
  private

  public

  end;

var
  MainForm: TMainForm;

implementation

{$R *.lfm}

{ TMainForm }

procedure TMainForm.BtnLoadCfgClick(Sender: TObject);
begin
  //
end;

procedure TMainForm.BtnNewCfgClick(Sender: TObject);
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

end;

procedure TMainForm.BtnKeyOpenClick(Sender: TObject);
begin
  //
end;

procedure TMainForm.BtnNewCfgClick(Sender: TObject);
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



end.

