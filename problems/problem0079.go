package problems

import (
    "strconv"
)

type problem0079 struct {
    id int
    description string
}

func (p problem0079) ID() int {
    return p.id
}

func (p problem0079) Description() string {
    return p.description
}

func (p problem0079) Solve() string {
    keycodes := []int{319, 680, 180, 690, 129, 620, 762, 689, 762, 318,
		368, 710, 720,
		710, 629, 168, 160, 689, 716, 731, 736, 729, 316, 729, 729, 710,
		769, 290, 719, 680, 318, 389, 162, 289, 162, 718, 729, 319, 790,
		680, 890, 362, 319, 760, 316, 729, 380, 319, 728, 716}

	digitRelationships := make(map[int]digitEntry)

	for _, keycode := range keycodes {
		firstDigit := keycode / 100
		secondDigit := (keycode / 10) % 10
		thirdDigit := keycode % 10

		if _, ok := digitRelationships[firstDigit]; !ok {
			digitRelationships[firstDigit] = digitEntry{make(map[int]struct{}),
				make(map[int]struct{})}
		}

		digitRelationships[firstDigit].before[secondDigit] = struct{}{}
		digitRelationships[firstDigit].before[thirdDigit] = struct{}{}

		if _, ok := digitRelationships[secondDigit]; !ok {
			digitRelationships[secondDigit] = digitEntry{make(map[int]struct{}),
				make(map[int]struct{})}
		}

		digitRelationships[secondDigit].after[firstDigit] = struct{}{}
		digitRelationships[secondDigit].before[thirdDigit] = struct{}{}

		if _, ok := digitRelationships[thirdDigit]; !ok {
			digitRelationships[thirdDigit] = digitEntry{make(map[int]struct{}),
				make(map[int]struct{})}
		}

		digitRelationships[thirdDigit].after[firstDigit] = struct{}{}
		digitRelationships[thirdDigit].after[secondDigit] = struct{}{}
	}

	digits := make([]int, 0)

	for len(digitRelationships) > 0 {
		for digit, entry := range digitRelationships {
			if len(entry.after) == 0 {
				digits = append(digits, digit)
				break
			}
		}

		lastFoundDigit := digits[len(digits)-1]

		delete(digitRelationships, lastFoundDigit)

		for _, entry := range digitRelationships {
			delete(entry.before, lastFoundDigit)
			delete(entry.after, lastFoundDigit)
		}
	}

	passcode := ""

	for _, digit := range digits {
		passcode += strconv.Itoa(digit)
	}

    return passcode
}

type digitEntry struct {
	before map[int]struct{}
	after  map[int]struct{}
}

func init() {
    Registry[79] = problem0079{79, "Given that the three characters are always asked for in order, analyse the list of keycodes so as to determine the shortest possible secret passcode of unknown length."}
}
