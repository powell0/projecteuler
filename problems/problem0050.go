package problems

import (
    "github.com/powell0/projecteuler/utilities/math/primes"
    "strconv"
)

type problem0050 struct {
    id int
    description string
}

func (p problem0050) ID() int {
    return p.id
}

func (p problem0050) Description() string {
    return p.description
}

func (p problem0050) Solve() string {

    const LIMIT = 1000000
    var result uint64
    var length int

    primesList := primes.ComputePrimes(LIMIT)

    for i := 0; i < len(primesList); i++ {
        tempPrimesList := primesList[i:]

        primeSequence := make([]uint64, len(tempPrimesList), len(tempPrimesList))
        copy(primeSequence, tempPrimesList)

        for j := 1; j < len(primeSequence); j++ {
            primeSequence[j] += primeSequence[j-1]
        }

        for j := len(primeSequence) - 1; j >= 0; j-- {
            if primeSequence[j] < LIMIT && primes.IsPrime(primeSequence[j]) {
                if j + 1 > length {
                    result = primeSequence[j]
                    length = j + 1
                    break
                }
            }
        }

    }

    return strconv.FormatUint(result, 10)
}

func init() {
    Registry[50] = problem0050{50, "Which prime, below one-million, can be written as the sum of the most consecutive primes?"}
}
