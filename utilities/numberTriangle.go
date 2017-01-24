package utilities

type numberTriangle struct {
    data [][]int64
    levels int
}

func (t *numberTriangle) initialize(levels int) {
    t.levels = levels
    t.data = make([][]int64, levels)
}

func (t numberTriangle) getParentsValues(level int, index int) (first int64,  second int64) {
    var firstParent int64
    var secondParent int64

    if level > 0 {
        if index <= 0 {
            firstParent = 0
            secondParent = t.data[level-1][0]
        } else if index >= level {
            firstParent = t.data[level-1][level-1]
            secondParent = 0
        } else {
            firstParent = t.data[level-1][index-1]
            secondParent = t.data[level-1][index]
        }
    }

    return firstParent, secondParent
}
