package problems

import (
    "fmt"
    "strconv"
)

type problem0078 struct {
    id int
    description string
}

func (p problem0078) ID() int {
    return p.id
}

func (p problem0078) Description() string {
    return p.description
}

func (p problem0078) Solve() string {
    const limit = 1000000
    const target = 100

    result := uint64(0)

    partitions := make([]uint64, limit+1)
    partitions[0] = 1
    partitions[1] = 1

    for n := 2; n <= limit; n++ {
        sum := uint64(0)

        for k := 1; k <= n; k++ {
            firstPartition := uint64(0)
            firstIndex := n - (k * (3*k - 1) / 2)

            if firstIndex >= 0 {
                firstPartition = partitions[firstIndex]
            }

            secondPartition := uint64(0)
            secondIndex := n - (k * (3*k + 1) / 2)

            if secondIndex >= 0 {
                secondPartition = partitions[secondIndex]
            }

            if n == 74 {
                fT := uint64(0)
                sT := uint64(0)

                if firstIndex >= 0 {
                    fT = partitions[firstIndex]
                }

                if secondIndex >= 0 {
                    sT = partitions[secondIndex]
                }

                fmt.Println(firstIndex, fT, secondIndex, sT)
            }

            tempSum := (firstPartition + secondPartition)

            if k % 2 == 0 {
                sum -= tempSum
            } else {
                sum += tempSum
            }

            sum %= 1000000

            if firstIndex < 0 && secondIndex < 0 {
                break
            }
        }

        if sum % target == 0 {
            result = uint64(n)
            fmt.Println(n, sum)
            //break
        }

        if n == 74 {
            break
        }

        partitions[n] = sum % 1000000
    }

    for i := 0; i < 75; i++ {
        fmt.Println(i, partitions[i])
    }
    
    return strconv.FormatUint(result, 10)
}

func init() {
    Registry[78] = problem0078{78, "Find the least value of n for which p(n) is divisible by one million."}
}
