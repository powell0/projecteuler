package main

import (
    "flag"
    "fmt"
    "github.com/powell0/projecteuler/problems"
    "sort"
    "strconv"
    "strings"
    "time"
)

var displayHelp bool
var displayProblems bool
var displayProblemsFull bool
var problemsToSolve problemSetFlag

func init() {
    
    flag.BoolVar(&displayHelp, "h", false, "Display usage information ")
    flag.BoolVar(&displayProblems, "d", false, "Display problem numbers")
    flag.BoolVar(&displayProblemsFull, "D", false, "Display problem numbers and descriptions")
    flag.Var(&problemsToSolve, "p", "Only run the given problem numbers (e.g. \"1, 2, 5-9, 10\")")

    flag.Parse()
}


func main() {
    // Create an index of the problems in ascending order
    problemList := make([]int, 0, len(problems.Registry))

    if len(problemsToSolve.problems) > 0 {
        for i := 0; i < len(problemsToSolve.problems); i++ {
            problemList = append(problemList, problemsToSolve.problems[i])
        }
    } else {
        for key, _ := range problems.Registry {
            problemList = append(problemList, key)
        }
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

type problemSetFlag struct {
    problems []int
}

func (p *problemSetFlag) String() string {
    var text string

    if len(p.problems) > 0 {

        for i := 0; i < len(p.problems); i++ {
            count := 0

            for j := i + 1; j < len(p.problems); j++ {
                if p.problems[j-1] + 1 != p.problems[j] {
                    break
                }

                count++
            }

            text += strconv.Itoa(p.problems[i])

            if count > 1 {
                i += count
                text += "-" + strconv.Itoa(p.problems[i])
            }

            text += ", "
        }
    }

    if len(text) > 0 {
        text = text[:len(text)-2]
    }

    return text
}

func (p *problemSetFlag) Set(value string) error {
    if len(p.problems) > 0 {
    	return fmt.Errorf("The problems flag is already set")
    }
    
    problems := strings.Split(value, ",")
    for _, problem := range problems {
    	problemRange := strings.Split(problem, "-")

        if len(problemRange) == 1 {
            problemNumber, err := strconv.Atoi(strings.TrimSpace(problemRange[0]))

            if err == nil {
                p.problems = append(p.problems, problemNumber)
            } else {
                return err
            }

        } else if len(problemRange) == 2 {
            problemNumberStart, err := strconv.Atoi(strings.TrimSpace(problemRange[0]))

            if err != nil {
                return err
            }

            problemNumberEnd, err := strconv.Atoi(strings.TrimSpace(problemRange[1]))

            if err != nil {
                return err
            }

            for i := problemNumberStart; i <= problemNumberEnd; i++ {
                p.problems = append(p.problems, i)
            }
        } else {
            return fmt.Errorf("Invalid problem set ", problem)
        }
    }
    return nil
}
