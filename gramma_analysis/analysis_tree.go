package gramma_analysis

import (
	"Compiler/utils/data_structure/queue"
)

type AnalysisTreeNode struct {
	ID        int
	Value     string
	Next      []*AnalysisTreeNode
	RealValue interface{}
}

func (node *AnalysisTreeNode) Mirror() {
	if node == nil {
		return
	}

	q := queue.NewQueue()
	q.PushBack(node)

	for q.Len() > 0 {
		cur := q.Pop().(*AnalysisTreeNode)

		for i := range cur.Next {
			q.PushBack(cur.Next[i])
		}

		for i := 0; i < len(cur.Next)/2; i++ {
			cur.Next[i], cur.Next[len(cur.Next)-1-i] = cur.Next[len(cur.Next)-1-i], cur.Next[i]
		}
	}
}
