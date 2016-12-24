package cursor

import (
	"image"
)

func MoveTo(p image.Point) bool {
	return moveTo(p)
}
