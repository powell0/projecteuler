package problems

import (
    "strconv"
)

type problem0034 struct {
    id int
    description string
}

func (p problem0034) ID() int {
    return p.id
}

func (p problem0034) Description() string {
    return p.description
}

func (p problem0034) Solve() string {

    const limit = 100000
    factorials := make([]int, 10, 10)
    factorials[0] = 1

    for i := 1; i < len(factorials); i++ {
        factorials[i] = i * factorials[i - 1]
    }

    result := uint64(0)

    for i := 10; i < limit; i++ {
        sum := 0
        number := i

        for number > 0 {
            sum += factorials[number % 10]
            number /= 10
        }

        if i == sum {
            result += uint64(i)
        }
    }

    return strconv.FormatUint(result, 10)
}

func init() {
    Registry[34] = problem0034{34, "Find the sum of all numbers which are equal to the sum of the factorial of their digits."}
}
