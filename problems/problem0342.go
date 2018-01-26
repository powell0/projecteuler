package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number"
    "strconv"
)

type problem0342 struct {
    id int
    description string
}

func (p problem0342) ID() int {
    return p.id
}

func (p problem0342) Description() string {
    return p.description
}

func (p problem0342) Solve() string {
    //const limit = 10000000000
    const limit = 10000000
    sum := uint64(0)

    for n := uint64(2); n < limit; n++ {
        sum = n*number.EulerTotient(n)
    }

    return strconv.FormatUint(sum, 10)
}

func init() {
    //Registry[342] = problem0342{342, "Find the sum of all numbers n, 1 < n < 10^10 such that phi(n^2) is a cube."}
}
