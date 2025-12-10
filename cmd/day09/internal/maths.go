package internal

type Point struct {
	X float64
	Y float64
}

type Line struct {
	Start Point
	End   Point
}

func (l Line) Horizontal() bool {
	return l.Start.Y == l.End.Y
}

func (l Line) Vertical() bool {
	return l.Start.X == l.End.X
}

func (l1 Line) Crosses(l2 Line) bool {
	// Normalize: ensure we have horizontal first, vertical second
	var hLine, vLine Line
	if l1.Horizontal() && l2.Vertical() {
		hLine, vLine = l1, l2
	} else if l1.Vertical() && l2.Horizontal() {
		hLine, vLine = l2, l1
	} else {
		return false
	}

	// Check X overlap: vertical line's X must be within horizontal line's X range
	vX := vLine.Start.X
	hX0 := hLine.Start.X
	hX1 := hLine.End.X
	if hX0 > hX1 {
		hX0, hX1 = hX1, hX0
	}
	if vX < hX0 || vX > hX1 {
		return false
	}

	// Check Y overlap: horizontal line's Y must be within vertical line's Y range
	hY := hLine.Start.Y
	vY0 := vLine.Start.Y
	vY1 := vLine.End.Y
	if vY0 > vY1 {
		vY0, vY1 = vY1, vY0
	}
	return hY >= vY0 && hY <= vY1
}

type Rectangle struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

func NewRectangle(minX, minY, maxX, maxY float64) Rectangle {
	if minX > maxX {
		minX, maxX = maxX, minX
	}
	if minY > maxY {
		minY, maxY = maxY, minY
	}
	return Rectangle{MinX: minX, MinY: minY, MaxX: maxX, MaxY: maxY}
}

func (r Rectangle) Area() float64 {
	width := r.MaxX - r.MinX + 1
	height := r.MaxY - r.MinY + 1
	return width * height
}

func (r Rectangle) Shrink(amount float64) Rectangle {
	return Rectangle{
		MinX: r.MinX + amount,
		MinY: r.MinY + amount,
		MaxX: r.MaxX - amount,
		MaxY: r.MaxY - amount,
	}
}
