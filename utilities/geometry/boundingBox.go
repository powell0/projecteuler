package geometry

type BoundingBox struct {
    Maximum Point
    Minimum Point
}

func (b BoundingBox) Segments() []LineSegment {
    side1 := NewLineSegmentFromCoordinates(b.Minimum.X, b.Minimum.Y, b.Maximum.X, b.Minimum.Y)
    side2 := NewLineSegmentFromCoordinates(b.Maximum.X, b.Minimum.Y, b.Maximum.X, b.Maximum.Y)
    side3 := NewLineSegmentFromCoordinates(b.Maximum.X, b.Maximum.Y, b.Minimum.X, b.Maximum.Y)
    side4 := NewLineSegmentFromCoordinates(b.Minimum.X, b.Maximum.Y, b.Minimum.X, b.Minimum.Y)

    return []LineSegment{side1, side2, side3, side4}
}

func (b BoundingBox) BoundingBox() BoundingBox {
    return b
}

func (b BoundingBox) ContainsPoint(point Point) bool {
    contains := point.X >= b.Minimum.X && point.X <= b.Maximum.X && point.Y >= b.Minimum.Y && point.X <= b.Maximum.X

    return contains
}

func NewBoundingBoxFromPoints(minimum Point, maximum Point) BoundingBox {
    return BoundingBox{maximum, minimum}
}

func NewBoundingBoxFromCoordinates(minX float64, minY float64, maxX float64, maxY float64) BoundingBox {
    return NewBoundingBoxFromPoints(Point{minX, minY}, Point{maxX, maxY})
}
