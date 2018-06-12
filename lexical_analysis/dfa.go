package lexical_analysis

import (
	"Compiler/utils/data_structure/queue"
	"Compiler/utils/data_structure/vector"
)

type DFA struct {
	n *NFA

	statusTable [][]*vector.Vector

	statusIDCount int

	statusIDMap map[string]int
	matrixMap   map[int]map[string]int

	endStatus map[int]struct{}
}

func NewDFA(n *NFA) (d *DFA) {
	d = &DFA{
		n:           n,
		matrixMap:   make(map[int]map[string]int),
		statusIDMap: make(map[string]int),
		endStatus:   make(map[int]struct{}),
	}

	status := make(map[string]struct{})
	q := queue.NewQueue()

	start := n.Start.FindClosure()
	q.PushBack(start)

	for q.Len() > 0 {
		curStatus := q.Pop().(*vector.Vector)
		newRow := []*vector.Vector{curStatus}

		for i := 0; i < curStatus.Len(); i++ {
			if curStatus.Get(i) == 2 {
				d.endStatus[d.getStatusID(curStatus)] = struct{}{}
			}
		}

		for i := range n.CharSet {
			newVector := vector.NewVector()
			visit := make(map[int]struct{})

			for vi := 0; vi < curStatus.Len(); vi++ {
				nextVector := n.IDMap[curStatus.Get(vi)].FindNext(n.CharSet[i])

				for ni := 0; ni < nextVector.Len(); ni++ {
					if _, ok := visit[nextVector.Get(ni)]; !ok {
						newVector.PushBack(nextVector.Get(ni))
						visit[nextVector.Get(ni)] = struct{}{}
					}
				}
			}

			d.addStatusToMap(curStatus, newVector, n.CharSet[i])
			newRow = append(newRow, newVector)
			if _, ok := status[newVector.String()]; !ok && newVector.Len() > 0 {
				q.PushBack(newVector)
				status[newVector.String()] = struct{}{}
			}
		}

		d.statusTable = append(d.statusTable, newRow)
	}

	return
}

func (d *DFA) String() (res [][]string) {
	row := []string{""}
	for i := 0; i < len(d.n.CharSet); i++ {
		row = append(row, d.n.CharSet[i])
	}
	res = append(res, row)

	for i := 0; i < len(d.statusTable); i++ {
		newRow := []string{}
		for j := 0; j < len(d.statusTable[i]); j++ {
			newRow = append(newRow, d.statusTable[i][j].String())
		}
		res = append(res, newRow)
	}

	return
}

func (d *DFA) MatrixString() (res [][]string) {
	row := []string{""}
	for i := 0; i < len(d.n.CharSet); i++ {
		row = append(row, d.n.CharSet[i])
	}
	res = append(res, row)

	for i := 0; i < len(d.statusTable); i++ {
		newRow := []string{}
		for j := 0; j < len(d.statusTable[i]); j++ {
			newRow = append(newRow, d.statusTable[i][j].String())
		}
		res = append(res, newRow)
	}

	return
}

func (d *DFA) getStatusID(status *vector.Vector) int {
	if _, ok := d.statusIDMap[status.String()]; !ok {
		d.statusIDCount++
		d.statusIDMap[status.String()] = d.statusIDCount
	}

	return d.statusIDMap[status.String()]
}

func (d *DFA) addStatusToMap(src, dest *vector.Vector, pass string) {

	if _, ok := d.matrixMap[d.getStatusID(src)]; !ok {
		d.matrixMap[d.getStatusID(src)] = make(map[string]int)
	}

	d.matrixMap[d.getStatusID(src)][pass] = d.getStatusID(dest)

}

func (d *DFA) Accept(s string) bool {
	return d.deepFind(1, s)
}

func (d *DFA) deepFind(start int, target string) bool {
	if target == "" {
		if _, ok := d.endStatus[start]; !ok {
			return false
		}
		return true
	}

	next, ok := d.matrixMap[start][target[0:1]]
	if !ok {
		return false
	}

	return d.deepFind(next, target[1:])
}
