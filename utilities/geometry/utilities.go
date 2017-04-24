package geometry

import (
    "math"
)

func PointInPolygon(point Point, polygon Polygon) bool {
    count := 0

    sides := polygon.Segments()

    boundingBox := polygon.BoundingBox()

    epsilonX := (boundingBox.Maximum.X - boundingBox.Minimum.X) / 100
    epsilonY := (boundingBox.Maximum.Y - boundingBox.Minimum.Y) / 100

    ray := NewLineSegmentFromPoints(point, Point{boundingBox.Maximum.X + epsilonX, boundingBox.Maximum.Y + epsilonY})

    for _, side := range sides {
        intersectionPoint := LineSegmentIntersection(ray, side)

        if intersectionPoint != nil {
            count++
        }
    }

    return count % 2 == 1
}

func LineSegmentIntersection(segment1 LineSegment, segment2 LineSegment) *Point {
    const epsilon = 0.000001;

    line1 := NewLineFromLineSegment(segment1)
    line2 := NewLineFromLineSegment(segment2)

    intersectionPoint := LineIntersection(line1, line2)

    if intersectionPoint != nil {
        // Calculate the bounding box that contains the intersection of the two line segments
        minX := math.Max(math.Min(segment1.Point1.X, segment1.Point2.X), math.Min(segment2.Point1.X, segment2.Point2.X))
        maxX := math.Min(math.Max(segment1.Point1.X, segment1.Point2.X), math.Max(segment2.Point1.X, segment2.Point2.X))
        minY := math.Max(math.Min(segment1.Point1.Y, segment1.Point2.Y), math.Min(segment2.Point1.Y, segment2.Point2.Y))
        maxY := math.Min(math.Max(segment1.Point1.Y, segment1.Point2.Y), math.Max(segment2.Point1.Y, segment2.Point2.Y))

        boundingBox := NewBoundingBoxFromCoordinates(minX, minY, maxX, maxY)

        // If the intersection point is outside of the bounding box, the two line segments do not intersect
        if !boundingBox.ContainsPoint(*intersectionPoint) {
            intersectionPoint = nil
        }
    }

    return intersectionPoint
}

func LineIntersection(line1 Line, line2 Line) *Point {
    var intersectionPoint *Point
    det := line1.A * line2.B - line2.A * line1.B

    if det != 0 {
        intersectionPoint = new(Point)

        intersectionPoint.X = (line2.B * line1.C - line1.B * line2.C) / det
        intersectionPoint.Y = (line1.A * line2.C - line2.A * line1.C) / det
    } else {
        intersectionPoint = nil
    }

    return intersectionPoint
}
