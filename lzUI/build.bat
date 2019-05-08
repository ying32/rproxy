@echo off
res2go -outmain false -outpath "../" -pause "ew"
echo // +build gui>> $
echo.>> $
type "..\MainForm.go" >> $
move $ "..\MainForm.go"