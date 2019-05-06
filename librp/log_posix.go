// +build !windows

package librp

import "fmt"

func textColor(cr uint16) {
	fmt.Print(fmt.Sprintf("\033[0;%dm", cr))
}
func textRed() {
	textColor(31)
}

func textYellow() {
	textColor(33)
}

func textDefault() {
	textColor(37)
}
