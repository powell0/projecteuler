package geometry

type Point struct {
    X float64
    Y float64
}

func NewPointFromCoordinates(x float64, y float64) Point {
    return Point{x, y}
}
