package problems

import (
    "container/list"
    "github.com/powell0/projecteuler/utilities/graph"
    "strconv"
)

type problem0015 struct {
    id int
    description string
}

func (p problem0015) ID() int {
    return p.id
}

func (p problem0015) Description() string {
    return p.description
}

func (p problem0015) Solve() string {
    const rows = 21
    const columns = 21

    pathCount := make([]int, rows * columns)
    pathCount[len(pathCount) - 1] = 1

    unvisitedVertices := make(map[int]struct{})

    vertices := make([][]graph.Vertex, rows, rows)

    for i := 0; i < rows; i++ {
        vertices[i] = make([]graph.Vertex, columns, columns)

        for j := 0; j < columns; j++ {
            id := i*columns + j
            vertices[i][j] = graph.NewVertex(id, 5)
            unvisitedVertices[id] = struct{}{}
        }
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < columns; j++ {
            source := &vertices[i][j] 

            if i < rows - 1 {
                destination := &vertices[i + 1][j]
                sourceToDestination := graph.NewEdge(source, destination, 0)
                destinationToSource := graph.NewEdge(destination, source, 0)
                source.Edges = append(source.Edges, &sourceToDestination)
                destination.Edges = append(destination.Edges, &destinationToSource)
            }
            
            if j < columns - 1 {
                destination := &vertices[i][j + 1]
                sourceToDestination := graph.NewEdge(source, destination, 0)
                destinationToSource := graph.NewEdge(destination, source, 0)
                source.Edges = append(source.Edges, &sourceToDestination)
                destination.Edges = append(destination.Edges, &destinationToSource)
            }
        }
    }

    nextVertex := vertices[rows - 1][columns - 1]
    verticesToVisitNext := list.New()

    for len(unvisitedVertices) > 0 {
        delete(unvisitedVertices, nextVertex.Id)

        for _, edge := range nextVertex.Edges {
            if nextVertex.Id > edge.Destination.Id {
               _, unvisited := unvisitedVertices[edge.Destination.Id]
               contains := false

               for e := verticesToVisitNext.Front(); e != nil; e = e.Next() {
                   if e.Value.(graph.Vertex).Id == edge.Destination.Id {
                       contains = true
                       break
                   }
               }

               if unvisited && !contains {
                   verticesToVisitNext.PushBack(*(edge.Destination))
               }
            } else {
                pathCount[nextVertex.Id] += pathCount[edge.Destination.Id];
            }
        }

        if verticesToVisitNext.Len() > 0 {
            e := verticesToVisitNext.Front()
            nextVertex = e.Value.(graph.Vertex)
            verticesToVisitNext.Remove(e)
        } else {
            break
        }
    }

    return strconv.Itoa(pathCount[0])
}

func init() {
    Registry[15] = problem0015{15, "Find the sum of all the primes below two million."}
}
