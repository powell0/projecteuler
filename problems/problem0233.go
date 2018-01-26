package problems

import (
    //"fmt"
    //"github.com/powell0/projecteuler/utilities/math/factoring"
    "math"
    "strconv"
)

type problem0233 struct {
    id int
    description string
}

func (p problem0233) ID() int {
    return p.id
}

func (p problem0233) Description() string {
    return p.description
}

func (p problem0233) Solve() string {
    /*
    const limit = 100000000
    n1 := 0
    n3 := 0

    for i := 0; i < limit; i++ {
        d1 := 4*i+1

        if d1 > limit {
            break
        }

        if limit % d1 == 0 {
            n1++
        }

        d3 := 4*i+3

        if limit % d3 == 0 {
            n3++
        }
    }

    r := 4*(n1 - n3)
    */

    const limit = 10000
    r := 0

    for i := uint64(0); i < limit; i++ {
        j := math.Sqrt(float64(limit*limit - i * i))

        if math.Floor(j) == j {
            r += 4
        }
    }
    return strconv.FormatUint(uint64(r), 10)
}

func init() {
    Registry[233] = problem0233{233, "Let f(N) be the number of points with integer coordinates that are on a circle passing through (0,0), (N,0),(0,N), and (N,N).  What is the sum of all positive integers N <= 10^11 such that f(N) = 420"}
}
