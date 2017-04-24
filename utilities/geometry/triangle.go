package geometry

import (
    "math"
)

type Triangle struct {
    Side1 LineSegment
    Side2 LineSegment
    Side3 LineSegment
}

func (t Triangle) Segments() []LineSegment {
    return []LineSegment{t.Side1, t.Side2, t.Side3}
}

func (t Triangle) BoundingBox() BoundingBox {
    minX := math.Min(t.Side1.Point1.X, math.Min(t.Side2.Point1.X, t.Side3.Point1.X))
    minY := math.Min(t.Side1.Point1.Y, math.Min(t.Side2.Point1.Y, t.Side3.Point1.Y))
    maxX := math.Max(t.Side1.Point1.X, math.Max(t.Side2.Point1.X, t.Side3.Point1.X))
    maxY := math.Max(t.Side1.Point1.Y, math.Max(t.Side2.Point1.Y, t.Side3.Point1.Y))

    return NewBoundingBoxFromCoordinates(minX, minY, maxX, maxY)
}

func NewTriangleFromPoints(point1 Point, point2 Point, point3 Point) Triangle {
    return Triangle{LineSegment{point1, point2}, LineSegment{point2, point3}, LineSegment{point3, point1}}
}

func NewTriangleFromCoordinates(x1 float64, y1 float64, x2 float64, y2 float64, x3 float64, y3 float64) Triangle {
    return NewTriangleFromPoints(Point{x1, y1}, Point{x2, y2}, Point{x3, y3})
}