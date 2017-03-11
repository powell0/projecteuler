package problems

import (
    "strconv"
)

type problem0061 struct {
    id int
    description string
}

func (p problem0061) ID() int {
    return p.id
}

func (p problem0061) Description() string {
    return p.description
}

func (p problem0061) Solve() string {
    sum := 0
    figurates := make(map[int]map[int][]figurateCandidate)
    figurates[3] = make(map[int][]figurateCandidate)
    figurates[4] = make(map[int][]figurateCandidate)
    figurates[5] = make(map[int][]figurateCandidate)
    figurates[6] = make(map[int][]figurateCandidate)
    figurates[7] = make(map[int][]figurateCandidate)
    figurates[8] = make(map[int][]figurateCandidate)

    n := 1
    done := false

    // Generate four digit triangle numbers
    for !done {
        number := n*(n+1)/2

        if number > 999 && number < 10000 {
            figurates[3][number] = make([]figurateCandidate, 0, 20) 
        } else if number >= 10000 {
            done = true
        }

        n++
    }

    n = 1
    done = false

    // Generate four digit square numbers
    for !done {
        number := n*(3*n-1)/2

        if number > 999 && number < 10000 {
            figurates[4][number] = make([]figurateCandidate, 0, 20) 
        } else if number >= 10000 {
            done = true
        }

        n++
    }

    n = 1
    done = false

    // Generate four digit pentagonal numbers
    for !done {
        number := n*n

        if number > 999 && number < 10000 {
            figurates[5][number] = make([]figurateCandidate, 0, 20) 
        } else if number >= 10000 {
            done = true
        }

        n++
    }

    n = 1
    done = false

    // Generate four digit hexagonal numbers
    for !done {
        number := n*(2*n-1)

        if number > 999 && number < 10000 {
            figurates[6][number] = make([]figurateCandidate, 0, 20) 
        } else if number >= 10000 {
            done = true
        }

        n++
    }

    n = 1
    done = false

    // Generate four digit heptagonal numbers
    for !done {
        number := n*(5*n-3)/2

        if number > 999 && number < 10000 {
            figurates[7][number] = make([]figurateCandidate, 0, 20) 
        } else if number >= 10000 {
            done = true
        }

        n++
    }

    n = 1
    done = false

    // Generate four digit octagonal numbers
    for !done {
        number := n*(3*n-2)

        if number > 999 && number < 10000 {
            figurates[8][number] = make([]figurateCandidate, 0, 20)
        } else if number >= 10000 {
            done = true
        }

        n++
    }

    // For each figurate
    for primaryFigurate, primaryFigurateNumbers := range figurates {

        // For each number of that figurate
        for primaryNumber, _ := range primaryFigurateNumbers {
            lastTwoDigits := primaryNumber % 100

            // Iterate over every other figurate
            for secondaryFigurate, secondaryFigurateNumbers := range figurates {
                if primaryFigurate == secondaryFigurate {
                    continue
                }

                // For each number of the other figurates
                for secondaryNumber, _ := range secondaryFigurateNumbers {
                    if primaryNumber == secondaryNumber {
                        continue
                    }

                    firstTwoDigits := secondaryNumber / 100

                    if firstTwoDigits == lastTwoDigits {
                        newCandidate := figurateCandidate{secondaryFigurate, secondaryNumber}
                        figurates[primaryFigurate][primaryNumber] = append(figurates[primaryFigurate][primaryNumber], newCandidate)
                    }
                }
            }

            if len(figurates[primaryFigurate][primaryNumber]) == 0 {
                delete(figurates[primaryFigurate], primaryNumber)
            }
        }
    }

    usedFigurates := make([]int, 0, 6)
    orderedSet := make([]int, 6)
    found := false

    for primaryNumber, _ := range figurates[3] {
        orderedSet[0] = primaryNumber
        findOrderedSetRecursively(figurates, 3, primaryNumber, usedFigurates, 1, primaryNumber, &orderedSet, &found)

        if found {
            for _, number := range orderedSet {
                sum += number
            }
            break
        }
    }
    
    return strconv.Itoa(sum)
}

type figurateCandidate struct {
    figurate int
    number int
}

func alreadyUsed(figurate int, usedFigurates []int) bool {
    alreadyUsed := false

    for _, used := range usedFigurates {
        if figurate == used {
            alreadyUsed = true
            break
        }
    }

    return alreadyUsed
}

func findOrderedSetRecursively(figurates map[int]map[int][]figurateCandidate, figurate int, number int, usedFigurates []int, level int, firstNumber int, orderedSet *[]int, found *bool) {
    usedFigurates = append(usedFigurates, figurate)

    candidates := figurates[figurate][number]

    for _, candidate := range candidates {
        if level == 6 {
            if candidate.figurate == usedFigurates[0] && candidate.number == firstNumber {
                *found = true
                break
            }
        } else {
            if !alreadyUsed(candidate.figurate, usedFigurates) {
                findOrderedSetRecursively(figurates, candidate.figurate, candidate.number, usedFigurates, level+1, firstNumber, orderedSet, found)

                if *found {
                    (*orderedSet)[level] = candidate.number
                    break
                }
            }
        }
    }
}

func init() {
    Registry[61] = problem0061{61, "Find the sum of the only ordered set of six cyclic 4-digit numbers for which each polygonal type: triangle, square, pentagonal, hexagonal, heptagonal, and octagonal, is represented by a different number in the set."}
}
