package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number"
    "github.com/powell0/projecteuler/utilities/math/primes"
    "strconv"
)

type problem0070 struct {
    id int
    description string
}

func (p problem0070) ID() int {
    return p.id
}

func (p problem0070) Description() string {
    return p.description
}

func (p problem0070) Solve() string {
    const limit = 10000000
    const minPrime = 1000
    const maxPrime = limit / minPrime

    primesList := primes.ComputePrimes(maxPrime)
    permutations := make(map[uint64]uint64)

    for i := 0; i < len(primesList) - 1; i++ {
        n1 := primesList[i]
        phi1 := n1 - 1

        for j := i + 1; j < len(primesList); j++ {
            n2 := primesList[j]
            phi2 := n2 - 1

            n := n1*n2

            if n > limit {
                break
            }

            phi := phi1*phi2

            digitsOfN := number.GetDigits(n)
            digitsOfPhi := number.GetDigits(phi)

            if number.IsPermutation(digitsOfN, digitsOfPhi) {
                permutations[n] = phi
            }
        }
    }

    minimumRatio := float64(limit)
    bestN := uint64(limit)

    for n, phi := range permutations {
        ratio := float64(n) / float64(phi)

        if ratio < minimumRatio {
            minimumRatio = ratio
            bestN = n
        }
    }

    return strconv.FormatUint(bestN, 10)
}

func init() {
    Registry[70] = problem0070{70, "Find the value of n, 1 < n < 10^7, for which phi(n) is a permutation of n and the ratio n/phi(n) produces a minimum."}
}
