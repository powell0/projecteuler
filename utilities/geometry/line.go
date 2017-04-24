package geometry

// A line in the form Ax + By = C
type Line struct {
    A float64
    B float64
    C float64
}

func NewLineFromCoordinates(x1 float64, y1 float64, x2 float64, y2 float64) Line {
    A := y1 - y2
    B := x2 - x1
    C := A*x1 + B*y1

    return Line{A, B, C}
}

func NewLineFromPoints(point1 Point, point2 Point) Line {
    return NewLineFromCoordinates(point1.X, point1.Y, point2.X, point2.Y)
}

func NewLineFromLineSegment(segment LineSegment) Line {
    return NewLineFromPoints(segment.Point1, segment.Point2)
}

func (l Line) IsVertical() bool {
    return l.B == 0
}

func (l Line) IsHorizontal() bool {
    return l.A == 0
}

func (l Line) IsOblique() bool {
    return !l.IsVertical() && !l.IsHorizontal()
}