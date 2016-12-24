package cursor

import(
	"syscall"
	"image"
)

var(
	user32 = syscall.MustLoadDLL("user32.dll")
	setCursorPos = user32.MustFindProc("SetCursorPos")
)


func moveTo(p image.Point) bool {
	r, _, _ := setCursorPos.Call(uintptr(p.X), uintptr(p.Y))
	return r == 1
}
