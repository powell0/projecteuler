package problems

import (
    "math"
    "strconv"
)

type problem0040 struct {
    id int
    description string
}

func (p problem0040) ID() int {
    return p.id
}

func (p problem0040) Description() string {
    return p.description
}

func (p problem0040) Solve() string {

    result := determineChampernowneConstantDigit(1)
    result *= determineChampernowneConstantDigit(10)
    result *= determineChampernowneConstantDigit(100)
    result *= determineChampernowneConstantDigit(1000)
    result *= determineChampernowneConstantDigit(10000)
    result *= determineChampernowneConstantDigit(100000)
    result *= determineChampernowneConstantDigit(1000000)

    return strconv.Itoa(result)
}

func determineChampernowneConstantDigit (index int) int {
    number := 1
    var digit int

    if index < 0 {
        digit = -1
    } else if index == 1 {
        digit = 1
    } else {
        index--

        delta := 0

        for index >= 0 {
            number++
            delta = int(math.Ceil(math.Log10(float64(number))))

            index -= delta
        }

        number--
        index += delta

        digit, _ = strconv.Atoi(strconv.Itoa(number)[index: index+1])
    }

    return digit
}

func init() {
    Registry[40] = problem0040{40, "If dn represents the nth digit of the fractional part, find the value of the following expression: d1 * d10 * d100 * d1000 * d10000 * d100000 * d1000000"}
}
