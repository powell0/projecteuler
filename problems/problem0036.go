package problems

import (
    "github.com/powell0/projecteuler/utilities/string/palindromes"
    "strconv"
)

type problem0036 struct {
    id int
    description string
}

func (p problem0036) ID() int {
    return p.id
}

func (p problem0036) Description() string {
    return p.description
}

func (p problem0036) Solve() string {
    count := 0
    
    for i := 1; i < 1000000; i++ {
        if palindromes.IsPalindrome(strconv.Itoa(i)) && palindromes.IsPalindrome(strconv.FormatInt(int64(i), 2)) {
            count += i
        }
    }

    return strconv.Itoa(count)
}

func init() {
    Registry[36] = problem0036{36, "Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2."}
}
