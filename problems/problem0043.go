package problems

import (
    "github.com/powell0/projecteuler/utilities/math/number/pandigital"
    "strconv"
)

type problem0043 struct {
    id int
    description string
}

func (p problem0043) ID() int {
    return p.id
}

func (p problem0043) Description() string {
    return p.description
}

func (p problem0043) Solve() string {
    pandigitalNumbers := pandigital.GeneratePandigitalStrings(9, true)

    sum := uint64(0)

    for _, pandigitalNumber := range pandigitalNumbers {
        substring := pandigitalNumber[1:4]

        number, err := strconv.Atoi(substring)

        if err == nil {
            if number % 2 != 0 {
                continue
            }
        } else {
            continue
        }

        substring = pandigitalNumber[2:5]

        number, err = strconv.Atoi(substring)

        if err == nil {
            if number % 3 != 0 {
                continue
            }
        } else {
            continue
        }

        substring = pandigitalNumber[3:6]

        number, err = strconv.Atoi(substring)

        if err == nil {
            if number % 5 != 0 {
                continue
            }
        } else {
            continue
        }

        substring = pandigitalNumber[4:7]

        number, err = strconv.Atoi(substring)

        if err == nil {
            if number % 7 != 0 {
                continue
            }
        } else {
            continue
        }

        substring = pandigitalNumber[5:8]

        number, err = strconv.Atoi(substring)

        if err == nil {
            if number % 11 != 0 {
                continue
            }
        } else {
            continue
        }

        substring = pandigitalNumber[6:9]

        number, err = strconv.Atoi(substring)

        if err == nil {
            if number % 13 != 0 {
                continue
            }
        } else {
            continue
        }

        substring = pandigitalNumber[7:10]

        number, err = strconv.Atoi(substring)

        if err == nil {
            if number % 17 != 0 {
                continue
            }
        } else {
            continue
        }

        result, err := strconv.ParseUint(pandigitalNumber, 10, 64)

        if err == nil {
            sum += result
        }
    }

    return strconv.FormatUint(sum, 10)
}

func init() {
    Registry[43] = problem0043{43, "Find the sum of all 0 to 9 pandigital numbers with the given substring divisibility properties"}
}
