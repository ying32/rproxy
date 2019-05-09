@echo off
res2go -outmain false -outpath "../" -pause "ew" -outres false
echo // +build gui>> $
echo.>> $
type "..\MainForm.go" >> $
move $ "..\MainForm.go"