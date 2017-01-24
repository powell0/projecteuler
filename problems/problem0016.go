package problems

import (
    "math/big"
    "strconv"
)

type problem0016 struct {
    id int
    description string
}

func (p problem0016) ID() int {
    return p.id
}

func (p problem0016) Description() string {
    return p.description
}

func (p problem0016) Solve() string {
    sum := uint64(0)
    var number big.Int

    number.Exp(big.NewInt(2), big.NewInt(1000), nil)

    result := number.String()

    for i := 0; i < len(result); i++ {
        digit, _ := strconv.ParseUint(result[i:i+1], 10, 64)
        sum += digit
    }
    
    return strconv.FormatUint(sum, 10)
}

func init() {
    Registry[16] = problem0016{16, "What is the sum of the digits of the number 2**1000?"}
}
