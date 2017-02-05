package graph

type Vertex struct {
    Id int
    Edges []*Edge
}

func NewVertex(id int, capacity int) Vertex {
    var v Vertex
    v.Id = id
    v.Edges = make([]*Edge, 0, capacity)

    return v
}
