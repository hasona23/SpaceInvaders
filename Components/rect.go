package Components

type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
	bottom int
	right  int
	left   int
	top    int
}

func (r Rect) Init(x, y, width, height int) Rect {
	return Rect{X: x, Y: y,
		Width:  width,
		Height: height,

		bottom: y + height,
		top:    y,
		right:  x + width,
		left:   x,
	}
}

func (r *Rect) Intersect(r2 Rect) bool {
	if r.bottom >= r2.top && r.top <= r2.bottom && r.left <= r2.right && r.right >= r2.left {
		return true
	}
	return false
}

func (r *Rect) Move(x, y int) {
	r.top = y
	r.bottom = y + r.Height
	r.right = x + r.Width
	r.left = x
}
