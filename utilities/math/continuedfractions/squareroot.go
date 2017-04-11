package continuedfractions

import (
    "math"
)

func ComputeSquareRootRepresenation(n int64) []int64 {
    representation := make([]int64, 0, 2)

    a := int64(math.Floor(math.Sqrt(float64(n))))
    a0 := a
    representation = append(representation, a)

    // If n is a square number, the continued fraction representation is trivial
    if a*a == n {
        return representation
    }

    m := int64(0)
    d := int64(1)

    for a != 2 * a0 {
        m = d*a - m
        d = (n - (m*m)) / d

        a = int64(math.Floor((float64(a0) + float64(m)) / float64(d)))

        representation = append(representation, a)
    }

    return representation
}
