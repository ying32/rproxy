package librp

import (
	"syscall"
)

const (
	STD_OUTPUT_HANDLE = ^uint32(11)
)

var (
	kernel32                 = syscall.NewLazyDLL("kernel32.dll")
	_GetStdHandle            = kernel32.NewProc("GetStdHandle")
	_SetConsoleTextAttribute = kernel32.NewProc("SetConsoleTextAttribute")

	windowsConsoleHandle = GetStdHandle(STD_OUTPUT_HANDLE)
)

func GetStdHandle(nStdHandle uint32) uintptr {
	ret, _, _ := _GetStdHandle.Call(uintptr(nStdHandle))
	return ret
}

func SetConsoleTextAttribute(hConsoleOutput uintptr, wAttributes uint16) bool {
	ret, _, _ := _SetConsoleTextAttribute.Call(hConsoleOutput, uintptr(wAttributes))
	return ret != 0
}

func textColor(cr uint16) {
	if windowsConsoleHandle <= 0 {
		return
	}
	SetConsoleTextAttribute(windowsConsoleHandle, cr)
}

func textRed() {
	textColor(8 | 4) // 8 | 4
}

func textYellow() {
	textColor(8 | 6) //8|6
}

func textDefault() {
	textColor(7) // 7
}
