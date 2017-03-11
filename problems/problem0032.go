package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number"
    "github.com/powell0/projecteuler/utilities/math/number/pandigital"
    "strconv"
)

type problem0032 struct {
    id int
    description string
}

func (p problem0032) ID() int {
    return p.id
}

func (p problem0032) Description() string {
    return p.description
}

func (p problem0032) Solve() string {
    const maxDigits = 9
    const minMultiplicanAndMultiplierDigits = 4
    const maxMultiplicanAndMultiplierDigits = 5

    maxMultiplicand := number.PowUint64(10, minMultiplicanAndMultiplierDigits)
    products := make(map[uint64]struct{})

    for multiplicand := uint64(2); multiplicand < maxMultiplicand; multiplicand++ {
        if !number.HasUniqueDigits(multiplicand) {
            continue
        }

        multiplicandDigitCount := uint64(number.CountDigits(multiplicand))

        maxMultiplierDigits := maxMultiplicanAndMultiplierDigits - multiplicandDigitCount
        minMultiplierDigits := minMultiplicanAndMultiplierDigits - multiplicandDigitCount

        maxMultiplier := number.PowUint64(10, maxMultiplierDigits)
        minMultiplier := number.PowUint64(10, minMultiplierDigits)

        for multiplier := minMultiplier; multiplier < maxMultiplier; multiplier++ {
            if !number.HasUniqueDigits(multiplicand) {
                continue
            }

            if !number.NumbersAreUnique(multiplicand, multiplier) {
                continue
            }

            product := multiplicand * multiplier

            pandigitalNumberString := strconv.FormatUint(multiplicand, 10) + strconv.FormatUint(multiplier, 10) + strconv.FormatUint(product, 10)

            if len(pandigitalNumberString) == 9 {
                pandigitalNumber, _ := strconv.ParseUint(pandigitalNumberString, 10, 64)

                if pandigital.IsPandigital(pandigitalNumber, 9) {
                    products[product] = struct{}{}
                }
            }
        }
    }

    sum := uint64(0)
   
    for k, _ := range products {
        sum += k
    }

    return strconv.FormatUint(sum, 10)
}

func init() {
    Registry[32] = problem0032{32, "Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital."}
}
