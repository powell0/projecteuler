package palindromes

import "unicode/utf8"

func IsPalindrome(text string) bool {
    isPalindrome := true

    for s, w, e := 0, 0, len(text); s <= e; s += w {
        firstRuneValue, width := utf8.DecodeRuneInString(text[s:])
        w = width
        e -= w
        lastRuneValue, width := utf8.DecodeRuneInString(text[e:])

        if firstRuneValue != lastRuneValue {
            isPalindrome = false
            break
        }
    }

    return isPalindrome
}
