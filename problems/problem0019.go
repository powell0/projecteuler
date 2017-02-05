package problems

import (
    "strconv"
)

type problem0019 struct {
    id int
    description string
}

const (
    monday = iota
    tuesday = iota
    wednesday = iota
    thursday = iota
    friday = iota
    saturday = iota
    sunday = iota
)

const (
    january = iota
    february = iota
    march = iota
    april = iota
    may = iota
    june = iota
    july = iota
    august = iota
    september = iota
    october = iota
    november = iota
    december = iota
)

func getFutureDay (startingDay int, daysInFuture int) int {
    return (startingDay + daysInFuture) % 7
}

func isLeapYear (year int) bool {
    return (year % 4 == 0) && ((year % 100 != 0) || ((year % 100 == 0) && (year % 400 == 0)))
}

func (p problem0019) ID() int {
    return p.id
}

func (p problem0019) Description() string {
    return p.description
}

func (p problem0019) Solve() string {
    count := 0
    daysToMoveForward := 365
    firstDayOfMonth := monday
    daysPerMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

    if isLeapYear(1900) {
        daysToMoveForward++
    }

    firstDayOfMonth = getFutureDay(firstDayOfMonth, daysToMoveForward);

    if firstDayOfMonth == sunday {
        count++
    }

    for year := 1901; year < 2001; year++ {
        for month := january; month <= december; month++ {
            daysToMoveForward = daysPerMonth[month]

            if isLeapYear(year) && month == february {
                daysToMoveForward++
            }

            firstDayOfMonth = getFutureDay(firstDayOfMonth, daysToMoveForward);

            if firstDayOfMonth == sunday {
                count++
            }
        }
    }

    return strconv.Itoa(count)
}

func init() {
    Registry[19] = problem0019{19, "How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?"}
}
