package geometry

type Polygon interface {
    Segments() []LineSegment
    BoundingBox() BoundingBox
}
