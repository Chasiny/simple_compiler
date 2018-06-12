package lexical_analysis

import (
	"github.com/Chasiny/simple_compiler/utils/data_structure/queue"
	"github.com/Chasiny/simple_compiler/utils/data_structure/vector"
	"container/list"
)

type Node struct {
	ID   int
	Next *list.List
}

func (n *Node) FindClosure() (v *vector.Vector) {
	v = vector.NewVector()

	v.PushBack(n.ID)
	visit := make(map[int]struct{})
	visit[n.ID] = struct{}{}

	q := queue.NewQueue()
	q.PushBack(n)

	for q.Len() > 0 {
		curNode := q.Pop().(*Node)

		for e := curNode.Next.Front(); e != nil; e = e.Next() {
			if e.Value.(*Line).Pass == "@" {
				if _, ok := visit[e.Value.(*Line).NextNode.ID]; !ok {
					visit[e.Value.(*Line).NextNode.ID] = struct{}{}
					v.PushBack(e.Value.(*Line).NextNode.ID)

					q.PushBack(e.Value.(*Line).NextNode)
				}
			}
		}
	}

	return
}

func (n *Node) FindNext(s string) (v *vector.Vector) {
	resMap := make(map[int]struct{})
	queueMap := make(map[int]struct{})

	q := queue.NewQueue()

	q.PushBack(n)
	queueMap[n.ID] = struct{}{}

	v = vector.NewVector()

	for q.Len() > 0 {
		curNode := q.Pop().(*Node)
		for e := curNode.Next.Front(); e != nil; e = e.Next() {
			if e.Value.(*Line).Pass == s {
				if _, ok := resMap[e.Value.(*Line).NextNode.ID]; !ok {
					v.PushBack(e.Value.(*Line).NextNode.ID)
					resMap[e.Value.(*Line).NextNode.ID] = struct{}{}

					closure := e.Value.(*Line).NextNode.FindClosure()
					for i := 0; i < closure.Len(); i++ {
						if _, ok := resMap[closure.Get(i)]; !ok {
							v.PushBack(closure.Get(i))
							resMap[closure.Get(i)] = struct{}{}
						}
					}
				}
			} else if e.Value.(*Line).Pass == "@" {
				q.PushBack(e.Value.(*Line).NextNode)
				queueMap[e.Value.(*Line).NextNode.ID] = struct{}{}
			}
		}
	}

	return
}
