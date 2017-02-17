package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number"
    "github.com/powell0/projecteuler/utilities/math/number/pandigital"
    "strconv"
)

type problem0038 struct {
    id int
    description string
}

func (p problem0038) ID() int {
    return p.id
}

func (p problem0038) Description() string {
    return p.description
}

func (p problem0038) Solve() string {
    const limit = 10000

    largestPandigital := uint64(123456789)

    for i := 2; i < limit; i++ {
        pandigitalNumber := ""

        n := 1
        for len(pandigitalNumber) < 9 {
            product := i * n

            pandigitalNumber += strconv.Itoa(product)

            n++
        }

        if len(pandigitalNumber) == 9 {
            newNumber, err := strconv.ParseUint(pandigitalNumber, 10, 64)

            if err == nil && pandigital.IsPandigital(newNumber, 9) {
                largestPandigital = number.MaxUint64(largestPandigital, newNumber)
            }
        }

    }

    return strconv.FormatUint(largestPandigital, 10)
}

func init() {
    Registry[38] = problem0038{38, "What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2, ... , n) where n > 1?"}
}
