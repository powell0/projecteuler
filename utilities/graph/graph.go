package graph

import (
    "errors"
    "strconv"
)

type Graph struct {
    Edges []Edge
    Vertices map[int]Vertex
}

func NewGraph() Graph {
    var g Graph
    g.Edges = make([]Edge, 0)
    g.Vertices = make(map[int]Vertex)

    return g
}

func (g Graph) ContainsVertex(id int) bool {
    _, ok := g.Vertices[id]

    return ok
}

func (g Graph) GetVertex(id int) (Vertex, error) {
    if g.ContainsVertex(id) {
        v, _ := g.Vertices[id]

        return v, nil
    } else {
        return Vertex{}, errors.New("Vertex " + strconv.Itoa(id) + "does not exist")
    }
}

func (g *Graph) AddVertex(v Vertex) error {
    if _, ok := g.Vertices[v.Id]; ok {
        return errors.New("Vertex " + strconv.Itoa(v.Id) + "already exists")
    }

    g.Vertices[v.Id] = v

    return nil
}

func (g *Graph) AddEdge(e Edge) error {
    g.Edges = append(g.Edges, e)

    return nil
}

func (g *Graph) AddNewVertex(id int, capacity int) error {
    if _, ok := g.Vertices[id]; ok {
        return errors.New("Vertex " + strconv.Itoa(id) + "already exists")
    }

    g.Vertices[id] = NewVertex(id, capacity)

    return nil
}

func (g *Graph) AddNewEdge(sourceVertexId int, destinationVertexId int, cost int) error {
    sourceVertex, sourceOk := g.Vertices[sourceVertexId]
    if !sourceOk {
        return errors.New("Source vertex " + strconv.Itoa(sourceVertexId) + "does not exist")
    }

    destVertex, destOk := g.Vertices[destinationVertexId]

    if !destOk {
        return errors.New("Destination vertex " + strconv.Itoa(destinationVertexId) + "does not exist")
    }

    e := NewEdge(&sourceVertex, &destVertex, cost)

    g.Edges = append(g.Edges, e)
    sourceVertex.Edges = append(sourceVertex.Edges, &e)

    return nil
}