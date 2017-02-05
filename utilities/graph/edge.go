package graph

type Edge struct {
    Source *Vertex
    Destination *Vertex
    Cost int
}

func NewEdge(source *Vertex, destination *Vertex,  cost int) Edge {
    var e Edge

    e.Source = source
    e.Destination = destination
    e.Cost = cost

    return e
}
