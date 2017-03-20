package graph

import (
    "errors"
    "math"
    "sort"
    "strconv"
)

type DijkstraSearch struct {
    graph Graph
    costs map[int]map[int]int
}

func NewDijkstraSearch(graph Graph) DijkstraSearch {
    var d DijkstraSearch

    d.graph = graph

    d.costs = make(map[int]map[int]int)

    for sourceVertexId, _ := range graph.Vertices {
        d.costs[sourceVertexId] = make(map[int]int)

        for destinationVertexId, _ := range graph.Vertices {
            cost := math.MaxInt32

            if sourceVertexId == destinationVertexId {
                cost = 0
            }

            d.costs[sourceVertexId][destinationVertexId] = cost
        }
    }

    return d
}

func (d *DijkstraSearch) ShortestPathCost(sourceVertexId int, destinationVertexId int) (int, error) {
    if !d.graph.ContainsVertex(sourceVertexId) {
        return math.MaxInt32, errors.New("Vertex " + strconv.Itoa(sourceVertexId) + "does not exist")
    }

    if !d.graph.ContainsVertex(destinationVertexId) {
        return math.MaxInt32, errors.New("Vertex " + strconv.Itoa(destinationVertexId) + "does not exist")
    }

    // get the cost of traveling from the source to the destination
    cost := d.costs[sourceVertexId][destinationVertexId]

    // If the cost is the maximum value, we have not computed the cost of traveling
    // from the source to the destination
    if cost == math.MaxInt32 {
        verticesToVisit := make([]vertexCost, 0, len(d.graph.Vertices))
        unvisitedVertices := make(map[int]struct{})

        for _, vertex := range d.graph.Vertices {
            unvisitedVertices[vertex.Id] = struct{}{}
        }

        currentVertex, _ := d.graph.GetVertex(sourceVertexId)

        for len(unvisitedVertices) > 0 {
            // mark current vertex as visited
            delete(unvisitedVertices, currentVertex.Id)

            // if we have reached our destination vertex, stop searching
            if currentVertex.Id == destinationVertexId {
                break
            }

            for _, edge := range currentVertex.Edges {
                if _, unvisited := unvisitedVertices[edge.Destination.Id]; unvisited {
                    totalCost := d.costs[sourceVertexId][currentVertex.Id] + edge.Cost

                    if totalCost < d.costs[sourceVertexId][edge.Destination.Id] {
                        d.costs[sourceVertexId][edge.Destination.Id] = totalCost
                        
                        index := -1

                        for i, v := range verticesToVisit {
                            if v.vertex.Id == edge.Destination.Id {
                                index = i
                                break
                            }
                        }
                        
                        if index > -1 {
                            verticesToVisit[index].cost = totalCost
                        } else {
                            verticesToVisit = append(verticesToVisit, vertexCost{*edge.Destination, totalCost})
                        }
                    }
                }
            }

            // sort the vertices by cost
            sort.Slice(verticesToVisit, func(i, j int) bool { return verticesToVisit[i].cost < verticesToVisit[j].cost})

            // get next cheapest vertex
            currentVertex = verticesToVisit[0].vertex

            verticesToVisit = verticesToVisit[1:]
        }

        // get the cost of traveling from the source to the destination
        cost = d.costs[sourceVertexId][destinationVertexId]
    }

    return cost, nil
}

type vertexCost struct {
    vertex Vertex
    cost int
}

