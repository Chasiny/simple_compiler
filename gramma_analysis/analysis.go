package gramma_analysis

import (
	"Compiler/utils/data_structure/stack"
	"fmt"
)

type node struct {
	data string
	id   int
}

func (ll1 *LL1Analysis) Analysis(input string) (res [][]string, t *AnalysisTreeNode, ok bool) {
	id := 0
	s := &stack.Stack{}

	l, r := 0, 1

	id++
	s.Push(&node{
		"#",
		id,
	})
	id++
	s.Push(&node{
		ll1.start,
		id,
	})

	tree := &AnalysisTreeNode{
		ID:    id,
		Value: ll1.start,
	}
	treeMap := make(map[int]*AnalysisTreeNode)
	treeMap[id] = tree

	for {
		if s.Len() <= 1 {
			break
		}

		if input[l:r] == `'` {
			r++
		}

		for {
		Loop:
			if s.Len() < 1 {
				break
			}
			curNode := s.GetAndPop().(*node)
			if curNode.data == input[l:r] {
				res = analysisStingTable(res, getStackString(s)+curNode.data, input[l:r], input[r:],
					fmt.Sprintf("匹配,弹出%s", curNode.data))
				break
			}

			mapStr, ok := ll1.AnalysisTable[curNode.data][input[l:r]]
			if !ok {
				return nil, nil, false
			}

			if mapStr == "@" {
				res = analysisStingTable(res, getStackString(s)+curNode.data, input[l:r], input[r:],
					fmt.Sprintf("匹配,弹出%s,%s->%s为@,不入栈", curNode.data, curNode.data, input[l:r]))
				goto Loop
			}

			res = analysisStingTable(res, getStackString(s)+curNode.data, input[l:r], input[r:],
				fmt.Sprintf("弹出%s,将%s逆序入栈", curNode.data, mapStr))

			ll, rr := len(mapStr)-1, len(mapStr)
			for {
				if ll < 0 {
					break
				}
				if mapStr[ll:rr] == "'" {
					ll--
				}

				id++
				newNode := &node{
					mapStr[ll:rr],
					id,
				}
				newTreeNode := &AnalysisTreeNode{
					ID:    id,
					Value: mapStr[ll:rr],
				}
				s.Push(newNode)
				treeMap[curNode.id].Next = append(treeMap[curNode.id].Next, newTreeNode)
				treeMap[id] = newTreeNode

				rr = ll
				ll = ll - 1
			}
		}

		l = r
		r = r + 1
	}

	if s.Len() > 1 {
		return res, tree, false
	}

	tree.Mirror()

	return res, tree, true
}

func analysisStingTable(in [][]string, s, cur, instr, caption string) (res [][]string) {
	var t []string
	t = append(t, s)
	t = append(t, cur)
	t = append(t, instr)
	t = append(t, caption)

	return append(in, t)
}

func getStackString(s *stack.Stack) (res string) {
	for i := 0; i < s.Len(); i++ {
		res += fmt.Sprintf("%+v", s.GetByIndex(i).(*node).data)
	}
	return
}
