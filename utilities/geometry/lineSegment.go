package geometry

type LineSegment struct {
    Point1 Point
    Point2 Point
}

func NewLineSegmentFromPoints(point1 Point, point2 Point) LineSegment {
    return NewLineSegmentFromCoordinates(point1.X, point1.Y, point2.X, point2.Y)
}

func NewLineSegmentFromCoordinates(x1 float64, y1 float64, x2 float64, y2 float64) LineSegment {
    return LineSegment{Point{x1, y1}, Point{x2, y2}}
}
