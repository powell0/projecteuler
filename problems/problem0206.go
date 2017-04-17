package problems

import (
    "math"
    "strconv"
)

type problem0206 struct {
    id int
    description string
}

func (p problem0206) ID() int {
    return p.id
}

func (p problem0206) Description() string {
    return p.description
}

func (p problem0206) Solve() string {
    upperLimit := uint64(math.Floor(math.Sqrt(1929394959697989990)))
    lowerLimit := uint64(math.Ceil(math.Sqrt(1020304050607080900)))

    result := uint64(0)

    for i := upperLimit; i >= lowerLimit; i-- {
        square := i * i

        if matchPattern(square) {
            result = i
            break
        }
    }


    return strconv.FormatUint(result, 10)
}

func matchPattern(n uint64) bool {

    if n % 10 != 0 {
        return false
    }

    n /= 100

    if n % 10 != 9 {
        return false
    }

    n /= 100

    if n % 10 != 8 {
        return false
    }

    n /= 100

    if n % 10 != 7 {
        return false
    }

    n /= 100

    if n % 10 != 6 {
        return false
    }

    n /= 100

    if n % 10 != 5 {
        return false
    }

    n /= 100

    if n % 10 != 4 {
        return false
    }

    n /= 100

    if n % 10 != 3 {
        return false
    }

    n /= 100

    if n % 10 != 2 {
        return false
    }

    n /= 100

    if n % 10 != 1 {
        return false
    }

    return true
}


func init() {
    Registry[206] = problem0206{206, "Find the unique positive integer whose square has the form 1_2_3_4_5_6_7_8_9_0, where each '_' is a single digit."}
}
