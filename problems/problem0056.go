package problems

import (
    "math/big"
    "strconv"
)

type problem0056 struct {
    id int
    description string
}

func (p problem0056) ID() int {
    return p.id
}

func (p problem0056) Description() string {
    return p.description
}

func (p problem0056) Solve() string {
    const limit = 100
    
    max := uint64(0)

    for a := int64(2); a < limit; a++ {
        for b := int64(2); b < limit; b++ {
            number := new(big.Int).Exp(big.NewInt(a), big.NewInt(b), nil)
            sum := sumDigits(number.String())

            if sum > max {
                max = sum
            }
        }
    }
    
    return strconv.FormatUint(max, 10)
}

func sumDigits(number string) uint64 {
    sum := uint64(0)

    for i := 0; i < len(number); i++ {
        digit, _ := strconv.ParseUint(number[i:i+1], 10, 64)
        sum += digit
    }

    return sum
}

func init() {
    Registry[56] = problem0056{56, "Considering natural numbers of the form, ab, where a, b < 100, what is the maximum digital sum?"}
}
