package problems

import (
    "github.com/powell0/projecteuler/utilities/math/continuedfractions"
    "strconv"
)

type problem0064 struct {
    id int
    description string
}

func (p problem0064) ID() int {
    return p.id
}

func (p problem0064) Description() string {
    return p.description
}

func (p problem0064) Solve() string {
    const limit = 10000

    count := 0

    for i := int64(2); i <= limit; i++ {
        rep := continuedfractions.ComputeSquareRootRepresenation(i)

        period := len(rep) - 1

        if period > 0 && period % 2 == 1{
            count++
        }
    }

    return strconv.Itoa(count)
}

func init() {
    Registry[64] = problem0064{64, "How many continued fractions for N <= 10000 have an odd period?"}
}
