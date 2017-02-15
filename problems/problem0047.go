package problems

import (
    "github.com/powell0/projecteuler/utilities/math/factoring"
    "github.com/powell0/projecteuler/utilities/math/primes"
    "strconv"
)

type problem0047 struct {
    id int
    description string
}

func (p problem0047) ID() int {
    return p.id
}

func (p problem0047) Description() string {
    return p.description
}

func (p problem0047) Solve() string {
    const limit = 1000000
    primesList := primes.ComputePrimes(limit)

    numberPrimeFactorizations := make([][]factoring.PrimeFactor, limit)

    for i := 2; i < len(numberPrimeFactorizations); i++ {
        numberPrimeFactorizations[i] = factoring.ComputePrimeFactors(uint64(i), primesList)
    }

    result := 0

    for i := 2; i < len(numberPrimeFactorizations)-3; i++ {
        
        if len(numberPrimeFactorizations[i]) != 4 || 
           len(numberPrimeFactorizations[i+1]) != 4 ||
           len(numberPrimeFactorizations[i+2]) != 4 ||
           len(numberPrimeFactorizations[i+3]) != 4 {
            continue
        }

        if !areDistinctPrimeFactorizations(numberPrimeFactorizations[i], numberPrimeFactorizations[i+1]) {
            continue
        }

        if !areDistinctPrimeFactorizations(numberPrimeFactorizations[i], numberPrimeFactorizations[i+2]) || 
           !areDistinctPrimeFactorizations(numberPrimeFactorizations[i+1], numberPrimeFactorizations[i+2]) {
            continue
        }

        if !areDistinctPrimeFactorizations(numberPrimeFactorizations[i], numberPrimeFactorizations[i+3]) || 
           !areDistinctPrimeFactorizations(numberPrimeFactorizations[i+1], numberPrimeFactorizations[i+3]) || 
           !areDistinctPrimeFactorizations(numberPrimeFactorizations[i+2], numberPrimeFactorizations[i+3]) {
            continue
        }

        result = i
        break
    }

    return strconv.Itoa(result)
}

func areDistinctPrimeFactorizations(first []factoring.PrimeFactor, second []factoring.PrimeFactor) bool {
    distinct := true

    for _, primeFactorFirst := range first {
        if distinct {
            for _, primeFactorSecond := range second {
                if primeFactorFirst.Number == primeFactorSecond.Number && primeFactorFirst.Count == primeFactorSecond.Count {
                    distinct = false
                    break;
                }
            }
        } else {
            break
        }
    }

    return distinct
}

func init() {
    Registry[47] = problem0047{47, "Find the first four consecutive integers to have four distinct prime factors each. What is the first of these numbers?"}
}
