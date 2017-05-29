package graph

import (
    "sort"
)

func ComputePrimMinimumSpanningTree(graph Graph) Graph {
    g := NewGraph()

    unusedVertices := make(map[int]struct{})

    firstVertexPicked := false

    // copy existing vertices from original graph
    // pick the first vertex to use in the spanning tree
    for k, _ := range graph.Vertices {
        g.Vertices[k] = NewVertex(k, 2)

        if firstVertexPicked {
            unusedVertices[k] = struct{}{}
        } else {
            firstVertexPicked = true
        }
    }

    // build the tree while there are vertices that need to be added
    for len(unusedVertices) > 0 {
        // create a list of possible edges to add to the minimum spanning tree
        // an edge is added if its source is already in the tree and its destination is not
        edgeCandidates := make([]Edge, 0, len(g.Edges))

        for _, edge := range graph.Edges {
            _, sourceNotInTree := unusedVertices[edge.Source.Id]
            _, destinationNotInTree := unusedVertices[edge.Destination.Id]

            if !sourceNotInTree && destinationNotInTree {
                edgeCandidates = append(edgeCandidates, edge)
            }
        }

        // sort candidate edges by cost
        sort.Slice(edgeCandidates, func(i, j int) bool { return edgeCandidates[i].Cost < edgeCandidates[j].Cost})

        // Get the edge with the lowest cost
        bestEdge := edgeCandidates[0]

        source := g.Vertices[bestEdge.Source.Id]
        destination := g.Vertices[bestEdge.Destination.Id]

        newEdge1 := NewEdge(&source, &destination, bestEdge.Cost)
        newEdge2 := NewEdge(&destination, &source, bestEdge.Cost)

        source.Edges = append(source.Edges, &newEdge1)
        destination.Edges = append(destination.Edges, &newEdge2)

        g.Edges = append(g.Edges, newEdge1)
        g.Edges = append(g.Edges, newEdge2)

        // Add the destination vertex to the minimum spanning tree
        delete(unusedVertices, destination.Id)
    }

    return g
}


