package problems

import (
    "ian-ferguson/math/number"
    "math/big"
    "strconv"
)

type problem0026 struct {
    id int
    description string
}

func (p problem0026) ID() int {
    return p.id
}

func (p problem0026) Description() string {
    return p.description
}

func (p problem0026) Solve() string {
    const NUMBER_LIMIT = 1000
    const PRECISION_LIMIT = 5000
    cycleLength := 0
    bestD := 0

    for d := 2; d < NUMBER_LIMIT; d++ {
        value := big.NewRat(1, int64(d))
        representation := trimTrailingZeroes(value.FloatString(PRECISION_LIMIT)[2:])

        tempCycleLength := 0

        for patternLength := 1; patternLength < len(representation) / 2; patternLength++ {
            pattern := representation[:patternLength]
            
            repeats := true

            for position := patternLength; position < len(representation) - patternLength; position += patternLength {
                subString := representation[position:position + patternLength]

                if subString != pattern {
                    repeats = false
                    break;
                }
            }

            if repeats {
                tempCycleLength = number.MaxInt(tempCycleLength, patternLength)
                break
            }
        }

        if tempCycleLength > cycleLength {
            cycleLength = tempCycleLength
            bestD = d
        }
    }
   
    return strconv.Itoa(bestD)
}

func init() {
    Registry[26] = problem0026{26, "Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part."}
}

func trimTrailingZeroes(s string) string {
    lastIndex := len(s)

    for lastIndex > 0 {
        if s[lastIndex - 1] == '0' {
            lastIndex--
        } else {
            break
        }
    }

    return s[:lastIndex]
 }
