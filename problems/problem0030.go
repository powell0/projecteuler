package problems

import (
    "strconv"
)

type problem0030 struct {
    id int
    description string
}

func (p problem0030) ID() int {
    return p.id
}

func (p problem0030) Description() string {
    return p.description
}

func (p problem0030) Solve() string {

    const limit = 1000000
    var result uint64 = 0

    for i := uint64(2); i < limit; i++ {
        sum := uint64(0)

        number := i

        for number > 0 {
            lastDigit := number % 10
            number /= 10

            sum += (lastDigit * lastDigit * lastDigit * lastDigit * lastDigit)
        }

        if sum == i {
            result += i
        }
    }

    return strconv.FormatUint(result, 10)
}

func init() {
    Registry[30] = problem0030{30, "Find the sum of all the numbers that can be written as the sum of fifth powers of their digits."}
}
