package problems

import (
    "strconv"
)

type problem0076 struct {
    id int
    description string
}

func (p problem0076) ID() int {
    return p.id
}

func (p problem0076) Description() string {
    return p.description
}

func (p problem0076) Solve() string {
    const limit = 100

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

            if k % 2 == 0 {
                sum -= (firstPartition + secondPartition)
            } else {
                sum += (firstPartition + secondPartition)
            }
        }

        partitions[n] = sum
    }
    
    return strconv.FormatUint(partitions[limit] - 1, 10)
}

func init() {
    Registry[76] = problem0076{76, "How many different ways can one hundred be written as a sum of at least two positive integers?"}
}
