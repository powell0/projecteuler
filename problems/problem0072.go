package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number"
    "strconv"
)

type problem0072 struct {
    id int
    description string
}

func (p problem0072) ID() int {
    return p.id
}

func (p problem0072) Description() string {
    return p.description
}

func (p problem0072) Solve() string {
    const limit = 1000000

    count := uint64(0)

    for d := uint64(2); d <= limit; d++ {
        count += number.EulerTotient(d)
    }

    return strconv.FormatUint(count, 10)
}

func init() {
    Registry[72] = problem0072{72, "How many elements would be contained in the set of reduced proper fractions for d <= 1,000,000?"}
}
