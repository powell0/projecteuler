package factoring

import (
    "github.com/powell0/projecteuler/utilities/math/primes"
    "math"
    "sort"
)

type uint64Slice []uint64

func (a uint64Slice) Len() int           { return len(a) }
func (a uint64Slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a uint64Slice) Less(i, j int) bool { return a[i] < a[j] }

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

func CountFactors (primeFactorization []PrimeFactor) uint {
    count := uint(1)

    for _, primeFactor := range primeFactorization {
        count *= (primeFactor.Count + 1)
    }

    return count;
}

func ComputeFactors (number uint64) []uint64 {
    maxNumberToTest := uint64(math.Sqrt(float64(number)))
    factors := make([]uint64, 0, maxNumberToTest / 10)
    
    for i := uint64(1); i <= maxNumberToTest; i++ {
        if number % i == 0 {
            factors = append(factors, i)
            
            if number / i != i {
                factors = append(factors, number / i)
            }
        }
    }

    sort.Sort(uint64Slice(factors))

    return factors 
}
