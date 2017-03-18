package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number"
    "strconv"
)

type problem0052 struct {
    id int
    description string
}

func (p problem0052) ID() int {
    return p.id
}

func (p problem0052) Description() string {
    return p.description
}

func (p problem0052) Solve() string {

    value := uint64(0)
    found := false

    for !found {
        value++
        found = true

        baseDigits := number.GetDigits(value)

        // 2x
        multipleDigits := number.GetDigits(value*2)

        if !number.IsPermutation(baseDigits, multipleDigits) {
            found = false
            continue
        }

        // 3x
        multipleDigits = number.GetDigits(value*3)

        if !number.IsPermutation(baseDigits, multipleDigits) {
            found = false
            continue
        }

        // 4x
        multipleDigits = number.GetDigits(value*4)

        if !number.IsPermutation(baseDigits, multipleDigits) {
            found = false
            continue
        }

        // 5x
        multipleDigits = number.GetDigits(value*5)

        if !number.IsPermutation(baseDigits, multipleDigits) {
            found = false
            continue
        }

        // 6x
        multipleDigits = number.GetDigits(value*6)

        if !number.IsPermutation(baseDigits, multipleDigits) {
            found = false
            continue
        }

    }

    return strconv.FormatUint(value, 10)
}

func init() {
    Registry[52] = problem0052{52, "Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits."}
}
