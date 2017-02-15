package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number"
    "github.com/powell0/projecteuler/utilities/math/primes"
    "strconv"
)

type problem0049 struct {
    id int
    description string
}

func (p problem0049) ID() int {
    return p.id
}

func (p problem0049) Description() string {
    return p.description
}

func (p problem0049) Solve() string {
    const limit = 10000
    result := ""
    primesList := primes.ComputePrimes(1000)
    count := len(primesList)

    primesList = primes.ComputePrimes(limit)
    primesList = primesList[count:]

    for i := 0; i < len(primesList) - 1; i++ {
        if primesList[i] == 1487 {
            continue
        }

        for j := i + 1; j < len(primesList); j++ {
            firstDigits := number.GetDigits(primesList[i])
            secondDigits := number.GetDigits(primesList[j])

            if !number.IsPermutation(firstDigits, secondDigits) {
                continue
            }

            nextNumber := 2 * primesList[j] - primesList[i]

            if nextNumber < limit && primes.IsPrime(nextNumber) {
                thirdDigits := number.GetDigits(nextNumber)

                if number.IsPermutation(firstDigits, thirdDigits) {
                    result += strconv.FormatUint(primesList[i], 10) 
                    result += strconv.FormatUint(primesList[j], 10) 
                    result += strconv.FormatUint(nextNumber, 10) 
                }
            }
        }
    }

    return result
}

func init() {
    Registry[49] = problem0049{49, "What 12-digit number do you form by concatenating the three terms in the arithmetic sequence of four digit primes who are permutations of each other?"}
}
