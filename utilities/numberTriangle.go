package utilities

type NumberTriangle struct {
    Data [][]int64
    Levels int
}

func (t *NumberTriangle) Initialize(levels int) {
    t.levels = levels
    t.data = make([][]int64, levels)
}

func (t NumberTriangle) GetParentsValues(level int, index int) (first int64,  second int64) {
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
