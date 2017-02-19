package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number"
    "github.com/powell0/projecteuler/utilities/string/palindromes"
    "strconv"
)

type problem0055 struct {
    id int
    description string
}

func (p problem0055) ID() int {
    return p.id
}

func (p problem0055) Description() string {
    return p.description
}

func (p problem0055) Solve() string {
    const numberLimit = 10000
    const lychrelLimit = 50

    count := 0

    for n := uint64(0); n < numberLimit; n++ {
        lychrel := true

        candidate := n

        for i := 0; i < lychrelLimit; i++ {
            candidate += number.Reverse(candidate)

            text := strconv.FormatUint(candidate, 10)

            if palindromes.IsPalindrome(text) {
                lychrel = false
                break
            }
        }

        if lychrel {
            count++
        }
    }
    
    return strconv.Itoa(count)
}

func init() {
    Registry[55] = problem0055{55, "How many Lychrel numbers are there below ten thousand?"}
}
