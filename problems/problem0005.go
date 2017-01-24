package problems

import (
    "ian-ferguson/math/factoring"
    "ian-ferguson/math/number"
    "strconv"
)

type problem0005 struct {
    id int
    description string
}

func (p problem0005) ID() int {
    return p.id
}

func (p problem0005) Description() string {
    return p.description
}

func (p problem0005) Solve() string {

    factors := make([]int, 21)
    factors[1] = 1

    for i := 2; i < len(factors); i++ {
        tempFactors := factoring.ComputePrimeFactors(uint64(i), nil)

        for _, factor := range tempFactors {
            index := factor.Number
            factors[index] = number.MaxInt(factors[index], int(factor.Count))
        }
    } 

    result := 1

    for i := 1; i < len(factors); i++ {
        result *= number.PowInt(i, factors[i])
    }

    return strconv.Itoa(result)
}

func init() {
    Registry[5] = problem0005{5, "What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?"}
}
