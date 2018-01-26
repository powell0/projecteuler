package pandigital

import (
    "fmt"
    "strconv"
)

func generateAllowedDigitsString(start int, end int) map[string]struct{} {
    allowedDigits := make(map[string]struct{})

    for i := start; i <= end; i++ {
        allowedDigits[strconv.Itoa(i)] = struct{}{}
    }

    return allowedDigits
}

func IsPandigitalString(number string, firstNumber int, lastNumber int) bool {
    allowedDigits := generateAllowedDigitsString(firstNumber, lastNumber)
    isPandigital := true

    for i := 0; i < len(number); i++ {
        digit := number[i: i+1]

        _, ok := allowedDigits[digit]

        if !ok {
            isPandigital = false
            break
        }

        delete(allowedDigits, digit)
    }
    
    return isPandigital
}

func IsPandigital(number uint64, n int) bool {
    result := true

    if n < 1 || n > 9 {
        result = false
    } else {
        digits := make([]bool, n+1)

        for number > 0 {
            digit := int(number % 10)
            number /= 10
            digits[digit] = true
        }

        for i := 1; i <= n; i++ {
            if !digits[i] {
                result = false
                break
            }
        }
    }
    return result;
}

func GeneratePandigitalStrings(n int, includeZero bool) []string {
    lowerLimit := 1

    if includeZero {
        lowerLimit = 0
    }

    if n < lowerLimit || n > 9 {
        panic(fmt.Sprintf("Argument %d must be between %d and 9", n, lowerLimit))
    }

    capacity := 1

    if includeZero {
        capacity *= n + 1
    }

    for i := 1; i <= n; i++ {
        capacity *= i
    }

    numbers := make([]string, 0, capacity)
    digits := make(map[int]struct{})

    if includeZero {
        digits[0] = struct{}{}
    }

    for i := 1; i <= n; i++ {
        digits[i] = struct{}{}
    }

    pandigitalNumber := ""

    generatePandigitalNumbersRecursively(digits, pandigitalNumber, &numbers)

    return numbers
}

func GeneratePandigitalNumbers(n int, includeZero bool) []uint64 {
    lowerLimit := 1

    if includeZero {
        lowerLimit = 0
    }

    if n < lowerLimit || n > 9 {
        panic(fmt.Sprintf("Argument %d must be between %d and 9", n, lowerLimit))
    }

    capacity := 1

    if includeZero {
        capacity *= n + 1
    }

    for i := 1; i <= n; i++ {
        capacity *= i
    }

    numberStrings := make([]string, 0, capacity)
    digits := make(map[int]struct{})

    if includeZero {
        digits[0] = struct{}{}
    }

    for i := 1; i <= n; i++ {
        digits[i] = struct{}{}
    }

    pandigitalNumber := ""

    generatePandigitalNumbersRecursively(digits, pandigitalNumber, &numberStrings)

    numbers := make([]uint64, len(numberStrings))

    for i := 0; i < len(numberStrings); i++ {
        number, err := strconv.ParseUint(numberStrings[i], 10, 64)

        if err == nil {
            numbers[i] = number
        }
    }

    return numbers
}

func generatePandigitalNumbersRecursively(digits map[int]struct{}, pandigitalNumber string, numbers *[]string) {
    if len(digits) == 0 {
        *numbers = append(*numbers, pandigitalNumber)

        return
    }

    for digit, _ := range digits {
        newPandigitalNumber := pandigitalNumber + strconv.Itoa(digit)
        
        newDigits := make(map[int]struct{})
        for k, v := range digits {
            newDigits[k] = v
        }

        delete(newDigits, digit)

        generatePandigitalNumbersRecursively(newDigits, newPandigitalNumber, numbers)
    }
}

