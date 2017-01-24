package problems

import (
    "strconv"
)

type problem0024 struct {
    id int
    description string
}

func (p problem0024) ID() int {
    return p.id
}

func (p problem0024) Description() string {
    return p.description
}

func (p problem0024) Solve() string {
    
    var numbers [10]bool
    var permutation string

    target := 999999

    for place := 10; place > 0; place-- {
        // calculate the number of permutations for the remaining digits
        permutations := computeFactorial(place - 1)
        
        // get the remaining available digits that can be used to finish the permutation
        availableDigits := make([]int,  0, place)

        for i := 0; i < len(numbers); i++ {
            if !numbers[i] {
                availableDigits = append(availableDigits, i)
            }
        }

        index := target / permutations
        target %= permutations

        digit := availableDigits[index]

        permutation += strconv.Itoa(digit)
        numbers[digit] = true
    }

    return permutation
}

func init() {
    Registry[24] = problem0024{24, "What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?"}
}

func computeFactorial(n int) int {
    if n > 0 {
        result := n * computeFactorial(n - 1)
        return result
    }

    return 1
}

