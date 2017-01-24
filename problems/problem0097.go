package problems

import (
    "strconv"
)

type problem0097 struct {
    id int
    description string
}

func (p problem0097) ID() int {
    return p.id
}

func (p problem0097) Description() string {
    return p.description
}

func (p problem0097) Solve() string {

    const MODULUS uint64 = 10000000000
    const EXPONENT uint64 = 7830457;
    var result uint64 = 1

    for i := uint64(0); i < EXPONENT; i++ {
        result *= 2
        result %= MODULUS
    }

    result *= 28433
    result %= MODULUS

    result += 1
    result %= MODULUS

    return strconv.FormatUint(result, 10)
}

func init() {
    Registry[97] = problem0097{97, "Find the last ten digits of the prime number 28433*2^(7830457)+1."}
}
