package stringutil

func CalculateAlphabeticalScore (text string) uint64 {
    score := 0

    for _, letter := range text {
        score += int(letter) - 96
    }

    return uint64(score)
}
