package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number/pandigital"
    "math"
    "strconv"
)

type problem0104 struct {
    id int
    description string
}

func (p problem0104) ID() int {
    return p.id
}

func (p problem0104) Description() string {
    return p.description
}

func (p problem0104) Solve() string {
    const filter = 1000000000
    const squareRootOfFive = 2.2360679775

    termNMinusTwo := uint64(1)
    termNMinusOne := uint64(1)
    index := uint64(2)

    for true {
        index++

        term := (termNMinusOne + termNMinusTwo) % filter

        lastNineDigitsArePandigital := pandigital.IsPandigital(term, 9)

        if lastNineDigitsArePandigital {
            _, result := math.Modf(float64(index) * math.Log10(math.Phi) - math.Log10(squareRootOfFive))
            result += 8
            firstNineDigits := math.Pow(10, result)

            firstNineDigitsArePandigital := pandigital.IsPandigital(uint64(firstNineDigits), 9)

            if firstNineDigitsArePandigital {
                break
            }
        }

        termNMinusTwo = termNMinusOne
        termNMinusOne = term
    }

    return strconv.FormatUint(index, 10)
}

func init() {
    Registry[104] = problem0104{104, "Given that Fk is the first Fibonacci number for which the first nine digits AND the last nine digits are 1-9 pandigital, find k."}
}
