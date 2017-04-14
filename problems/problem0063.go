package problems

import (
    "math"
    "strconv"
)

type problem0063 struct {
    id int
    description string
}

func (p problem0063) ID() int {
    return p.id
}

func (p problem0063) Description() string {
    return p.description
}

func (p problem0063) Solve() string {
    count := 0

    for i := float64(1); i < 10; i++ {
        n := float64(0)

        for true {
            n++

            digits := math.Floor(n * math.Log10(i)) + 1

            if (digits == n) {
                count++
            } else {
                break;
            }
        }
    }

    return strconv.Itoa(count)
}

func init() {
    Registry[63] = problem0063{63, "How many n-digit positive integers exist which are also an nth power?"}
}
