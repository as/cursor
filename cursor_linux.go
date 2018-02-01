package cursor

import (
	"fmt"
	"image"
	"os/exec"
)

func widSelf() string {
	out, _ := exec.Command("xdotool", "getactivewindow").Output()
	return string(out)
}

func MoveRelativeTo(p image.Point) {
	id := widSelf()
	exec.Command("xdotool", "mousemove", "-w", id, fmt.Sprint(p.X), fmt.Sprint(p.Y))
}

func MoveTo(p image.Point) bool {
	return moveTo(p)
}

func moveTo(p image.Point) bool {
	return false
}
