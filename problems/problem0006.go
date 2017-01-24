package problems

import (
    "github.com/powell0/projecteuler/utilities/math/summation"
    "strconv"
)

type problem0006 struct {
    id int
    description string
}

func (p problem0006) ID() int {
    return p.id
}

func (p problem0006) Description() string {
    return p.description
}

func (p problem0006) Solve() string {

    result := summation.Sum1ToN(100) * summation.Sum1ToN(100) - summation.SumOfSquares(100)

    return strconv.FormatUint(result, 10)
}

func init() {
    Registry[6] = problem0006{6, "Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum."}
}
