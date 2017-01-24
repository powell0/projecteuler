package problems

import (
    "unicode"
    "unicode/utf8"
    "strconv"
)

type problem0017 struct {
    id int
    description string
}

func (p problem0017) ID() int {
    return p.id
}

func (p problem0017) Description() string {
    return p.description
}

func (p problem0017) Solve() string {
    count := 0

    ones := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    teens := []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
    tens := []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
    hundreds := []string{"", "one hundred", "two hundred", "three hundred", "four hundred", "five hundred", "six hundred", "seven hundred", "eight hundred", "nine hundred"}

    for n := 1; n < 1000; n++ {
        number := n
        text := ""

        if number >= 100 {
            index := number / 100
            number %= 100

            text += hundreds[index]

            if number > 0 {
                text += " and "
            }
        }

        if number >= 20 {
            index := number / 10
            number %= 10

            text += tens[index]
            text += " "
            text += ones[number]
        } else if number >= 10 {
            index := number % 10

            text += teens[index]
        } else if number >= 0 {
            text += ones[number]
        }

        count += countLetters(text)
    }

    count += countLetters("one thousand")

    return strconv.Itoa(count)
}

func init() {
    Registry[17] = problem0017{17, "If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?"}
}

func countLetters(text string) int {
    count := 0

    for s, w := 0, 0; s < len(text); s += w {
        runeValue, width := utf8.DecodeRuneInString(text[s:])
        w = width

        if unicode.IsLetter(runeValue) {
            count++
        }
    }

    return count
}
