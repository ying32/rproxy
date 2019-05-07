// +build !windows

package librp

import "fmt"

func textColor(cr uint16) {
	if IsGUI {
		return
	}
	fmt.Print(fmt.Sprintf("\033[0;%dm", cr))
}
func textRed() {
	textColor(31)
}

func textYellow() {
	textColor(33)
}

func textGreen() {
	textColor(32)
}

func textDefault() {
	textColor(37)
}
