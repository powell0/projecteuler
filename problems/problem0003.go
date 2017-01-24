package problems

import (
    "github.com/powell0/projecteuler/utilities/math/factoring"
    "strconv"
)

type problem0003 struct {
    id int
    description string
}

func (p problem0003) ID() int {
    return p.id
}

func (p problem0003) Description() string {
    return p.description
}

func (p problem0003) Solve() string {
    factors := factoring.ComputePrimeFactors(600851475143, nil)

    largestPrime := factors[len(factors)-1].Number

    return strconv.FormatUint(largestPrime, 10)
}

func init() {
    Registry[3] = problem0003{3, "What is the largest prime factor of the number 600851475143?"}
}
