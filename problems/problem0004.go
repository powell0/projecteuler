package problems

import (
    "ian-ferguson/math/number"
    "ian-ferguson/string/palindromes"
    "strconv"
)

type problem0004 struct {
    id int
    description string
}

func (p problem0004) ID() int {
    return p.id
}

func (p problem0004) Description() string {
    return p.description
}

func (p problem0004) Solve() string {
    result := 0

    for i := 999; i > 99; i-- {
        for j := 999; j > 99; j-- {
            product := i * j

            if palindromes.IsPalindrome(strconv.Itoa(product)) {
                result = number.MaxInt(result, product)
            }
        }
    }

    return strconv.Itoa(result)
}

func init() {
    Registry[4] = problem0004{4, "Find the largest palindrome made from the product of two 3-digit numbers."}
}
