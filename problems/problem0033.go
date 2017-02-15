package problems

import (
    "math/big"
    "strconv"
)

type problem0033 struct {
    id int
    description string
}

func (p problem0033) ID() int {
    return p.id
}

func (p problem0033) Description() string {
    return p.description
}

func (p problem0033) Solve() string {
    fractions := make([]big.Rat, 0, 10)

    for denominator := 10; denominator < 100; denominator++ {
        if denominator % 10 == 0 {
            continue
        }

        for numerator := 10; numerator < denominator; numerator++ {
            if numerator % 10 == 0 {
                continue
            }

            numeratorString := strconv.Itoa(numerator)
            denominatorString := strconv.Itoa(denominator)

            first := big.NewRat(int64(numerator), int64(denominator))

            if numeratorString[0] == denominatorString[0] {
                a, _ := strconv.Atoi(numeratorString[1:])
                b, _ := strconv.Atoi(denominatorString[1:])
                second := big.NewRat(int64(a), int64(b))

                if first.Cmp(second) == 0 {
                    fractions = append(fractions,  *first)
                }
            }

            if numeratorString[0] == denominatorString[1] {
                a, _ := strconv.Atoi(numeratorString[1:])
                b, _ := strconv.Atoi(denominatorString[:1])
                second := big.NewRat(int64(a), int64(b))

                if first.Cmp(second) == 0 {
                    fractions = append(fractions,  *first)
                }
            }

            if numeratorString[1] == denominatorString[0] {
                a, _ := strconv.Atoi(numeratorString[:1])
                b, _ := strconv.Atoi(denominatorString[1:])
                second := big.NewRat(int64(a), int64(b))

                if first.Cmp(second) == 0 {
                    fractions = append(fractions,  *first)
                }
            }

            if numeratorString[1] == denominatorString[1] {
                a, _ := strconv.Atoi(numeratorString[:1])
                b, _ := strconv.Atoi(denominatorString[:1])
                second := big.NewRat(int64(a), int64(b))

                if first.Cmp(second) == 0 {
                    fractions = append(fractions,  *first)
                }
            }

        }
    }

    result := big.NewRat(1,1)

    for _, fraction := range fractions {
        result.Mul(result, &fraction)
    }

    return result.Denom().String()
}

func init() {
    Registry[33] = problem0033{33, "There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.  If the product of these four fractions is given in its lowest common terms, find the value of the denominator."}
}
