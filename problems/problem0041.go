package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number"
    "github.com/powell0/projecteuler/utilities/math/number/pandigital"
    "github.com/powell0/projecteuler/utilities/math/primes"
    "github.com/powell0/projecteuler/utilities/math/summation"
    "strconv"
)

type problem0041 struct {
    id int
    description string
}

func (p problem0041) ID() int {
    return p.id
}

func (p problem0041) Description() string {
    return p.description
}

func (p problem0041) Solve() string {
    const limit = 1000000

    result := uint64(0)

    for i := 2; i <= 9; i++ {
        sum := summation.Sum1ToN(uint64(i))

        if sum % 3 != 0 {
            pandigitalNumbers := pandigital.GeneratePandigitalNumbers(i, false)

            for _, pandigitalNumber := range pandigitalNumbers {
                if primes.IsPrime(pandigitalNumber) {
                    result = number.MaxUint64(result, pandigitalNumber)
                }
            }
        }
    }

    return strconv.FormatUint(result, 10)
}

func init() {
    Registry[41] = problem0041{41, "What is the largest n-digit pandigital prime that exists?"}
}
