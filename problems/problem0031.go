package problems

import (
    "strconv"
)

type problem0031 struct {
    id int
    description string
}

func (p problem0031) ID() int {
    return p.id
}

func (p problem0031) Description() string {
    return p.description
}

func (p problem0031) Solve() string {
    count := 1
    totalPence := 200

    for oneHundredPence := 0; oneHundredPence <= totalPence / 100; oneHundredPence++ {
        remainingPence := totalPence - 100 * oneHundredPence

        for fiftyPence := 0; fiftyPence <= remainingPence / 50; fiftyPence++ {
            remainingPence -= 50 * fiftyPence

            for twentyPence := 0; twentyPence <= remainingPence / 20; twentyPence++ {
                remainingPence -= 20 * twentyPence

                for tenPence := 0; tenPence <= remainingPence / 10; tenPence++ {
                    remainingPence -= 10 * tenPence

                    for fivePence := 0; fivePence <= remainingPence / 5; fivePence++ {
                        remainingPence -= 5 * fivePence

                        for twoPence := 0; twoPence <= remainingPence / 2; twoPence++ {
                            remainingPence -= 2 * twoPence

                            count++

                            remainingPence += 2 * twoPence
                        }

                        remainingPence += 5 * fivePence
                    }

                    remainingPence += 10 * tenPence
                }

                remainingPence += 20 * twentyPence
            }

            remainingPence += 50 * fiftyPence
        }
    }
    
    return strconv.Itoa(count)
}

func init() {
    Registry[31] = problem0031{31, "How many different ways can two pounds be made using any number of coins?"}
}
