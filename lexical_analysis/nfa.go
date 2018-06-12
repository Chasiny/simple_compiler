package lexical_analysis

import (
	"Compiler/utils/data_structure/queue"
	"Compiler/utils/regexp_utils"
	"container/list"
	"fmt"
)

type Line struct {
	Pass     string
	NextNode *Node
}

type NFA struct {
	count   int
	Start   *Node
	CharSet []string
	IDMap   map[int]*Node
}

func NewNFA(regexp string) (res *NFA) {
	res = &NFA{
		IDMap: make(map[int]*Node),
	}

	startNode := res.createNode()
	endNode := res.createNode()

	res.Start = startNode
	res.Start.Next.PushBack(&Line{
		Pass:     regexp,
		NextNode: endNode,
	})

	res.extend()
	res.initCharSet()

	return
}

func (n *NFA) extend() {
	q := queue.NewQueue()
	q.PushBack(n.Start)

	for q.Len() > 0 {
		curNode := q.Pop().(*Node)

		for e := curNode.Next.Front(); e != nil; e = e.Next() {
			pass := e.Value.(*Line).Pass
			nextNode := e.Value.(*Line).NextNode

			if closure, ok := regexp_utils.SpitClosure(pass); ok {
				curNode.Next.Remove(e)

				newNode := n.createNode()
				curNode.Next.PushBack(&Line{
					Pass:     "@",
					NextNode: newNode,
				})
				newNode.Next.PushBack(&Line{
					Pass:     "@",
					NextNode: nextNode,
				})
				newNode.Next.PushBack(&Line{
					Pass:     closure,
					NextNode: newNode,
				})

				q.PushBack(curNode)
				q.PushBack(newNode)
			} else if or := regexp_utils.SpitOr(pass); len(or) > 1 {
				curNode.Next.Remove(e)

				for i := range or {
					curNode.Next.PushBack(&Line{
						Pass:     or[i],
						NextNode: nextNode,
					})
				}

				q.PushBack(curNode)

			} else if and := regexp_utils.SpitAnd(pass); len(and) > 1 {
				curNode.Next.Remove(e)

				preNode := curNode
				q.PushBack(curNode)
				for i := 0; i < len(and)-1; i++ {
					newNode := n.createNode()
					preNode.Next.PushBack(&Line{
						Pass:     and[i],
						NextNode: newNode,
					})

					q.PushBack(newNode)
					preNode = newNode
				}

				preNode.Next.PushBack(&Line{
					Pass:     and[len(and)-1],
					NextNode: nextNode,
				})
			}
		}
	}
}

func (n *NFA) createNode() *Node {
	n.count++
	return &Node{
		ID:   n.count,
		Next: list.New(),
	}
}

func (n *NFA) String() (res []string) {
	q := queue.NewQueue()
	q.PushBack(n.Start)

	visit := make(map[int]struct{})
	visit[n.Start.ID] = struct{}{}

	for q.Len() > 0 {
		curNode := q.Pop().(*Node)

		for e := curNode.Next.Front(); e != nil; e = e.Next() {
			res = append(res, fmt.Sprintf("%d -> %s -> %d", curNode.ID, e.Value.(*Line).Pass, e.Value.(*Line).NextNode.ID)+"\n")

			if _, ok := visit[e.Value.(*Line).NextNode.ID]; !ok {
				visit[e.Value.(*Line).NextNode.ID] = struct{}{}
				q.PushBack(e.Value.(*Line).NextNode)
			}
		}
	}

	return
}

func (n *NFA) initCharSet() {
	q := queue.NewQueue()
	q.PushBack(n.Start)

	visit := make(map[int]struct{})
	visit[n.Start.ID] = struct{}{}
	n.IDMap[n.Start.ID] = n.Start

	charVisit := make(map[string]struct{})

	for q.Len() > 0 {
		curNode := q.Pop().(*Node)

		for e := curNode.Next.Front(); e != nil; e = e.Next() {

			if _, ok := charVisit[e.Value.(*Line).Pass]; !ok && e.Value.(*Line).Pass != "@" {
				n.CharSet = append(n.CharSet, e.Value.(*Line).Pass)
				charVisit[e.Value.(*Line).Pass] = struct{}{}
			}

			if _, ok := visit[e.Value.(*Line).NextNode.ID]; !ok {
				visit[e.Value.(*Line).NextNode.ID] = struct{}{}
				n.IDMap[e.Value.(*Line).NextNode.ID] = e.Value.(*Line).NextNode
				q.PushBack(e.Value.(*Line).NextNode)
			}
		}
	}

	return
}
