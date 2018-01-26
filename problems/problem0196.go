package problems

import (
    "fmt"
    "github.com/powell0/projecteuler/utilities/math/primes"
    "github.com/powell0/projecteuler/utilities/math/summation"
    "math"
    "strconv"
)

type problem0196 struct {
    id int
    description string
}

func (p problem0196) ID() int {
    return p.id
}

func (p problem0196) Description() string {
    return p.description
}

func (p problem0196) Solve() string {
    /*
    const firstRow = 5678027
    const secondRow = 7208785
    */

    const firstRow = 8

    //firstRowSum := uint64(0)
    //firstRowPrevStartingNumber := RowStartingNumber(firstRow - 1)
    firstRowStartingNumber := RowStartingNumber(firstRow)
    //firstRowNextStartingNumber := RowStartingNumber(firstRow + 1)

    firstRowPrimesInATriplet := make(map[uint64]struct{})

    for i := uint64(0); i < firstRow; i++ {
        //fmt.Println(firstRowStartingNumber + i)

        // if target number is already part of a triplet, skip it
        if _, ok := firstRowPrimesInATriplet[firstRowStartingNumber + i]; ok {
            continue
        }

        // if target number is not a prime, continue
        if !primes.IsPrime(firstRowStartingNumber + i) {
            continue
        }

        //fmt.Println(firstRowStartingNumber + i, "is prime")


        /*
        if (currRowPlusOne, currentColPlusOne) is prime:
        	check (currRow, currentColPlusTwo), (currRowPlusTwo, currentCol), (currRowPlusTwo, currentColPlusOne), (currRowPlusTwo, currentColPlusTwo) 
        */ 

        //neighboringPrimes := 0
        //neighbor := uint64(0)

        // prev row,  prev col
        //neighbor = firstRowPrevStartingNumber + i - 1

    }

    //fmt.Println(firstRow, firstRowSum)


    /*
    const secondRow = 9

    secondRowSum := uint64(0)
    secondRowPrevStartingNumber := RowStartingNumber(secondRow - 1)
    secondRowStartingNumber := RowStartingNumber(secondRow)
    secondRowNextStartingNumber := RowStartingNumber(secondRow + 1)
    */

    //fmt.Println("The answer for row", 8, "is", SumOfPrimeTripletMembersInRow(8))
    //fmt.Println("The answer for row", 9, "is", SumOfPrimeTripletMembersInRow(9))
    //fmt.Println("The answer for row", 10000, "is", SumOfPrimeTripletMembersInRow(10000))
    
    //fmt.Println("The answer for row", 5678027, "is", SumOfPrimeTripletMembersInRow(5678027))

    fmt.Println("Row", 5678027, "start=", RowStartingNumber(5678027), "end<", RowStartingNumber(5678027+1))
    fmt.Println("Row", 7208785, "start=", RowStartingNumber(7208785), "end<", RowStartingNumber(7208785+1))

    primeLimit := uint64(math.Sqrt(float64(25983294192506)) + 1)

    primesList := primes.ComputePrimes(primeLimit)
    fmt.Println(len(primesList))
    count := uint64(0)
    
    for i := RowStartingNumber(7208785); i < RowStartingNumber(7208785+1); i++ {
        if IsPrime(i, primesList) {
            count++
        }
    }

    return strconv.FormatUint(count, 10)
}

func IsPrime(number uint64,  primesList []uint64) bool {
    isPrime := number > 1

    for i := 0; i < len(primesList); i++ {
        if number % primesList[i] == 0 {
            isPrime = false
            break
        }
    }

    return isPrime
}

func SumOfPrimeTripletMembersInRow(row uint64) uint64 {
    currRowMinusTwoStartingNumber := RowStartingNumber(row - 2)
    currRowMinusOneStartingNumber := RowStartingNumber(row - 1)
    currRowStartingNumber := RowStartingNumber(row)
    currRowPlusOneStartingNumber := RowStartingNumber(row + 1)
    currRowPlusTwoStartingNumber := RowStartingNumber(row + 2)
    currRowPlusThreeStartingNumber := RowStartingNumber(row + 3)

    currRowPrimesInATriplet := make(map[uint64]struct{})

    

    for i := uint64(0); i < row; i++ {
        // if target number is already part of a triplet, skip it
        if _, ok := currRowPrimesInATriplet[currRowStartingNumber + i]; ok {
            continue
        }

        // if target number is not a prime, continue
        if !primes.IsPrime(currRowStartingNumber + i) {
            continue
        }

        neighbor := uint64(0)

        // neighbor: current row - 1, current col - 1
        neighbor = currRowMinusOneStartingNumber + i - 1

        if neighbor >= currRowMinusOneStartingNumber {
            if primes.IsPrime(neighbor) {
                // check remaining neighbors
                neighbors := make([]uint64, 0, 4)

                if currRowMinusTwoStartingNumber + i - 2 >= currRowMinusTwoStartingNumber {
                    neighbors = append(neighbors, currRowMinusTwoStartingNumber + i - 2)
                }

                if currRowMinusTwoStartingNumber + i - 1 >= currRowMinusTwoStartingNumber {
                    neighbors = append(neighbors, currRowMinusTwoStartingNumber + i - 1)
                }

                neighbors = append(neighbors, currRowMinusTwoStartingNumber + i)

                if currRowStartingNumber + i - 2 >= currRowStartingNumber {
                    neighbors = append(neighbors, currRowStartingNumber + i - 2)
                }

                if AreAnyNeighborsPrime(neighbors) {
                    currRowPrimesInATriplet[currRowStartingNumber + i] = struct{}{}
                    continue
                }
            }
        }

        // neighbor: current row - 1, current col
        neighbor = currRowMinusOneStartingNumber + i

        if primes.IsPrime(neighbor) {
            // check remaining neighbors
            neighbors := make([]uint64, 0, 3)

            if currRowMinusTwoStartingNumber + i - 1 >= currRowMinusTwoStartingNumber {
                neighbors = append(neighbors, currRowMinusTwoStartingNumber + i - 1)
            }

            neighbors = append(neighbors, currRowMinusTwoStartingNumber + i)

            if currRowMinusTwoStartingNumber + i + 1 < currRowMinusOneStartingNumber {
                neighbors = append(neighbors, currRowMinusTwoStartingNumber + i + 1)
            }

            if AreAnyNeighborsPrime(neighbors) {
                currRowPrimesInATriplet[currRowStartingNumber + i] = struct{}{}
                continue
            }
        }

        // neighbor: current row - 1, current col + 1
        neighbor = currRowMinusOneStartingNumber + i + 1

        if neighbor < currRowStartingNumber {
            if primes.IsPrime(neighbor) {
                // check remaining neighbors (currRow, currentColPlusTwo)
                neighbors := make([]uint64, 0, 4)

                neighbors = append(neighbors, currRowMinusTwoStartingNumber + i)

                if currRowMinusTwoStartingNumber + i + 1 < currRowMinusOneStartingNumber {
                    neighbors = append(neighbors, currRowMinusTwoStartingNumber + i + 1)
                }

                if currRowMinusTwoStartingNumber + i + 2 < currRowMinusOneStartingNumber {
                    neighbors = append(neighbors, currRowMinusTwoStartingNumber + i + 2)
                }

                if currRowStartingNumber + i + 2 < currRowPlusOneStartingNumber {
                    neighbors = append(neighbors, currRowStartingNumber + i + 2)
                }

                if AreAnyNeighborsPrime(neighbors) {
                    currRowPrimesInATriplet[currRowStartingNumber + i] = struct{}{}
                    continue
                }
            }
        }

        // neighbor: current row + 1, current col - 1
        neighbor = currRowPlusOneStartingNumber + i - 1

        if neighbor < currRowPlusOneStartingNumber {
            if primes.IsPrime(neighbor) {
                // check remaining neighbors
                neighbors := make([]uint64, 0, 4)

                if currRowStartingNumber + i - 2 >= currRowStartingNumber {
                    neighbors = append(neighbors, currRowStartingNumber + i - 2)
                }

                if currRowPlusTwoStartingNumber + i - 2 >= currRowPlusTwoStartingNumber {
                    neighbors = append(neighbors, currRowPlusTwoStartingNumber + i - 2)
                }

                if currRowPlusTwoStartingNumber + i - 1 >= currRowPlusTwoStartingNumber {
                    neighbors = append(neighbors, currRowPlusTwoStartingNumber + i - 1)
                }

                neighbors = append(neighbors, currRowPlusTwoStartingNumber + i)

                if AreAnyNeighborsPrime(neighbors) {
                    currRowPrimesInATriplet[currRowStartingNumber + i] = struct{}{}
                    continue
                }
            }
        }

        // neighbor: current row + 1, current col
        neighbor = currRowPlusOneStartingNumber + i

        if primes.IsPrime(neighbor) {
            // check remaining neighbors
             neighbors := make([]uint64, 0, 3)

                if currRowPlusTwoStartingNumber + i - 1 >= currRowPlusTwoStartingNumber {
                    neighbors = append(neighbors, currRowPlusTwoStartingNumber + i - 1)
                }

                neighbors = append(neighbors, currRowPlusTwoStartingNumber + i)

                if currRowPlusTwoStartingNumber + i + 1 < currRowPlusThreeStartingNumber {
                    neighbors = append(neighbors, currRowPlusTwoStartingNumber + i + 1)
                }

                if AreAnyNeighborsPrime(neighbors) {
                    currRowPrimesInATriplet[currRowStartingNumber + i] = struct{}{}
                    continue
                }
        }

        // neighbor: current row + 1, current col + 1
        neighbor = currRowPlusOneStartingNumber + i + 1

        if neighbor < currRowPlusThreeStartingNumber {
            if primes.IsPrime(neighbor) {
                // check remaining neighbors (currRowPlusTwo, currentColPlusTwo) 
                neighbors := make([]uint64, 0, 4)
               
                if currRowStartingNumber + i + 2 < currRowPlusOneStartingNumber {
                    neighbors = append(neighbors, currRowStartingNumber + i + 2)
                }

                neighbors = append(neighbors, currRowPlusTwoStartingNumber + i)

                if currRowPlusTwoStartingNumber + i + 1 < currRowPlusThreeStartingNumber {
                    neighbors = append(neighbors, currRowPlusTwoStartingNumber + i + 1)
                }

                if currRowPlusTwoStartingNumber + i + 2 < currRowPlusThreeStartingNumber {
                    neighbors = append(neighbors, currRowPlusTwoStartingNumber + i + 2)
                }

                if AreAnyNeighborsPrime(neighbors) {
                    currRowPrimesInATriplet[currRowStartingNumber + i] = struct{}{}
                    continue
                }
            }
        }
    }

    sum := uint64(0)

    for prime, _ := range currRowPrimesInATriplet {
        sum += prime
    }

    return sum
}

func RowStartingNumber(row uint64) uint64 {
    startingNumber := summation.Sum1ToN(row - 1) + 1

    return startingNumber
}

func AreAnyNeighborsPrime(neighbors []uint64) bool {
    result := false

    for _, neighbor := range neighbors {
        if primes.IsPrime(neighbor) {
            result = true
            break
        }
    }

    return result
}

func init() {
    Registry[196] = problem0196{196, "Define S(n) as the sum of the primes in row n which are elements of any prime triplet.  Find  S(5678027) + S(7208785)."}
}
