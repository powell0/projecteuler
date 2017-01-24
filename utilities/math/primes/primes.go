package primes

import "math"

type sieveElement struct {
    number uint64
    prime bool
}

func newSieveElement(number uint64) sieveElement {
    return sieveElement{number, true}
}

func ComputePrimes(length uint64) []uint64 {
    sieve := make([]sieveElement, length)

    sieve[0] = sieveElement{0, false}
    sieve[1] = sieveElement{0, false}

    for i := uint64(2); i < length; i++ {
        sieve[i] = newSieveElement(i)
    }

    var count uint64 = 0
    var p uint64 = 2

    for p < length {
        if sieve[p].prime {
            count++
            q := p * 2

            for q < length {
                sieve[q].prime = false
                q += p
            }
        }

        p++
    }

    primes := make([]uint64, 0, count)

    for i := uint64(2); i < length; i++ {
        if sieve[i].prime {
            primes = append(primes, i)
        }
    }

    return primes
}

func IsPrime(number uint64) bool {
    isPrime := number > 1

    primeLimit := uint64(math.Sqrt(float64(number)) + 1)

    for i := uint64(2); i < primeLimit; i++ {
        if number % i == 0 {
            isPrime = false
            break
        }
    }

    return isPrime
}
