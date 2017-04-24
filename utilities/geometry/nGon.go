package geometry

import (
    "math"
)

type NGon struct {
    Sides []LineSegment
}

func (n NGon) Segments() []LineSegment {
    return n.Sides
}

func (n NGon) BoundingBox() BoundingBox {
    minX := math.MaxFloat64
    minY := math.MaxFloat64
    maxX := -math.MaxFloat64
    maxY := -math.MaxFloat64

    for _, side := range n.Sides {
        minX = math.Min(minX, math.Min(side.Point1.X, side.Point2.X))
        minY = math.Min(minY, math.Min(side.Point1.Y, side.Point2.Y))
        maxX = math.Max(maxX, math.Max(side.Point1.X, side.Point2.X))
        maxY = math.Max(maxY, math.Max(side.Point1.Y, side.Point2.Y))
    }

    return NewBoundingBoxFromCoordinates(minX, minY, maxX, maxY)
}

func NewNGonFromPoints(points []Point) NGon {
    sides := make([]LineSegment, 0, len(points))

    for i := 1; i < len(points); i++ {
        side := NewLineSegmentFromPoints(points[i-1], points[i])
        sides = append(sides, side)
    }

    side := NewLineSegmentFromPoints(points[len(points)-1], points[0])
    sides = append(sides, side)

    return NGon{sides}
}

func NewNGonFromCoordinates(xCoords []float64, yCoords []float64) NGon {
    if len(xCoords) != len(yCoords) {
        panic("Mismatched numbers of coordinates")
    }

    points := make([]Point, 0, len(xCoords))

    for i := 0; i < len(xCoords); i++ {
        point := NewPointFromCoordinates(xCoords[i], yCoords[i])

        points = append(points, point)
    }

    return NewNGonFromPoints(points)
}
