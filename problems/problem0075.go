package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number"
    "strconv"
)

type problem0075 struct {
    id int
    description string
}

func (p problem0075) ID() int {
    return p.id
}

func (p problem0075) Description() string {
    return p.description
}

func (p problem0075) Solve() string {
    const limit = 1500000
    count := 0
    lengths := make([]int, limit+1)

    for n := 1; n < limit; n++ {
        for m := n + 1; m <= limit; m += 2 {
            length := 2*m*m + 2*m*n

            if length >= limit {
                break
            }

            if number.GCD(uint64(n), uint64(m)) == 1 {
                
                for k := 1; k <= limit / length; k++ {
                    lengths[k * length]++
                }
            }
        }
    }

    for i := 0; i < len(lengths); i++ {
        if lengths[i] == 1 {
            count++
        }
    }

    return strconv.Itoa(count)
}

func init() {
    Registry[75] = problem0075{75, "Given that L is the length of the wire, for how many values of L <= 1,500,000 can exactly one integer sided right angle triangle be formed?"}
}
