package factoring

import (
    "ian-ferguson/math/primes"
    "math"
)

type PrimeFactor struct {
    Number uint64
    Count uint
}

func newPrimeFactor(number uint64, count uint) PrimeFactor {
    return PrimeFactor{number, count}
}

func ComputePrimeFactors(number uint64, primeList []uint64) []PrimeFactor {
    primeLimit := uint64(math.Sqrt(float64(number)) + 1)

    if primeList == nil || len(primeList) > 0 || primeLimit > primeList[len(primeList)-1] {
        primeList = primes.ComputePrimes(primeLimit)
    }

    primeFactors := make([]PrimeFactor, 0, len(primeList))

    for _, prime := range primeList {
        primeFactor := newPrimeFactor(prime, 0)

        for number % prime == 0 {
            number /= prime
            primeFactor.Count++
        }

        if primeFactor.Count > 0 {
            primeFactors = append(primeFactors, primeFactor)
        }
    }

    if primes.IsPrime(number) {
        primeFactors = append(primeFactors, newPrimeFactor(number, 1))
    }

    return primeFactors
}
