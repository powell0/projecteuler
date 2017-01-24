package main

import (
    "flag"
    "fmt"
    "github.com/powell0/projecteuler/problems"
    "sort"
    "time"
)

var displayHelp bool
var displayProblems bool
var displayProblemsFull bool
var problemsToSolve []int

func init() {
    
    flag.BoolVar(&displayHelp, "h", false, "Display usage information ")
    flag.BoolVar(&displayProblems, "d", false, "Display problem numbers")
    flag.BoolVar(&displayProblemsFull, "D", false, "Display problem numbers and descriptions")

    flag.Parse()
}

func main() {
    // Create an index of the problems in ascending order
    problemList := make([]int, 0, len(problems.Registry))

    for key, _ := range problems.Registry {
        problemList = append(problemList, key)
    }

    sort.Ints(problemList)

    if displayHelp {
        flag.PrintDefaults()
    } else if displayProblems || displayProblemsFull {
        // Iterate through the problems in sorted order
        for _, problemNumber := range problemList {
            problem := problems.Registry[problemNumber]
            fmt.Printf("Problem %d", problem.ID())

            if displayProblemsFull {
                fmt.Printf(": %s", problem.Description())
            }

            fmt.Printf("\n")
        }
    } else {
        start := time.Now()

        // Iterate through the problems in sorted order
        for _, problemNumber := range problemList {
            problem := problems.Registry[problemNumber]
            results, ellapsedTime := solveProblem(problem)

            fmt.Printf("Problem %d solved in %s: %s\n", problem.ID(), ellapsedTime, results)
        }

        fmt.Printf("\n%d problems solved in %s\n", len(problemList), time.Since(start))
    }
}

func solveProblem (problem problems.Problem) (string, time.Duration) {
     start := time.Now()
     results := problem.Solve()

     return results, time.Since(start)
}

