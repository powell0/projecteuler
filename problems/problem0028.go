package problems

import (
    "strconv"
)

const (
    east = iota
    south = iota
    west = iota
    north = iota
)

type problem0028 struct {
    id int
    description string
}

func (p problem0028) ID() int {
    return p.id
}

func (p problem0028) Description() string {
    return p.description
}

func (p problem0028) Solve() string {
    const size = 1001

    grid := buildGrid(size)

    sum := 0

    for i := 0; i < size; i++ {
        sum += grid[i][i]
        sum += grid[i][size - 1 - i];
    }

    sum--;

    return strconv.Itoa(sum)
}

func buildGrid (size int) [][]int {
    grid := make([][]int, size, size)

    for i := 0; i < size; i++ {
        grid[i] = make([]int, size, size)
    }

    number := 1
    currentRow := size / 2
    currentCol := size / 2

    grid[currentRow][currentCol] = number
    number++

    previousDistance := 0
    previousDirection := -1
    currentDirection := -1

    for number <= size * size {
        previousDirection = currentDirection
        currentDirection = nextDirection(previousDirection)

        distance := nextDistance(previousDirection, currentDirection, previousDistance)
        previousDistance = distance

        for distance > 0 && number <= (size * size) {

            if currentDirection == east {
                currentCol++
            } else if currentDirection == south {
                currentRow++
            } else if currentDirection == west {
                currentCol--
            } else if currentDirection == north {
                currentRow--
            }

            grid[currentRow][currentCol] = number
            distance--
            number++
        }
    }

    return grid
}

func nextDirection (currentDirection int) int {
    return (currentDirection + 1) % 4
}

func nextDistance (previousDirection int, currentDirection int, previousDistance int) int {
    nextDistance := -1

    if previousDirection == -1 && currentDirection == east {
        nextDistance = 1
    } else if previousDirection == north && currentDirection == east {
        nextDistance = previousDistance + 1
    } else if previousDirection == east && currentDirection == south {
        nextDistance = previousDistance;
    } else if previousDirection == south && currentDirection == west {
        nextDistance = previousDistance + 1
    } else if previousDirection == west && currentDirection == north {
        nextDistance = previousDistance
    }

    return nextDistance
}

func init() {
    Registry[28] = problem0028{28, "What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?"}
}
