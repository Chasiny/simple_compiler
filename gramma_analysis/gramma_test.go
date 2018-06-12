package gramma_analysis

import (
	"Compiler/utils/data_structure/queue"
	"fmt"
	"testing"
)

func TestLL1Analysis_IsAllEmpty(t *testing.T) {
	ll1, err := NewLL1Analysis(`E->TE'
		E'->+TE'|@
		T->FT'
		T'->*FT'|@
		F->(E)|i|@`)
	if err != nil {
		fmt.Println(err)
		return
	}

	ll1.Show()

	if !ll1.isAllEmpty("FT'") {
		t.Fatal()
	}
}

func TestLL1Analysis_IsAllEmpty2(t *testing.T) {
	ll1, err := NewLL1Analysis(
		`E->TE'
		E'->+TE'|@
		T->FT'
		T'->*FT'|@
		F->(E)|i`)
	if err != nil {
		fmt.Println(err)
	}

	ll1.Show()

}

func TestLL1Analysis_Analysis(t *testing.T) {
	ll1, err := NewLL1Analysis(
		`E->TE'
		E'->+TE'|@
		T->FT'
		T'->*FT'|@
		F->(E)|i`)
	if err != nil {
		fmt.Println(err)
	}

	ll1.Analysis("i+i*i#")
}

func TestLL1Analysis_Tree(t *testing.T) {
	ll1, err := NewLL1Analysis(
		`E->TE'
		E'->+TE'|@
		T->FT'
		T'->*FT'|@
		F->(E)|i`)
	if err != nil {
		fmt.Println(err)
	}

	_, tree, ok := ll1.Analysis("i+i*i#")
	if !ok {
		t.Fatalf("no ok")
		return
	}

	q := queue.NewQueue()
	q.PushBack(tree)
	for q.Len() > 0 {
		cur := q.Pop().(*AnalysisTreeNode)
		for i := range cur.Next {
			t.Log(fmt.Sprintf("%s(%d) -> %s(%d)", cur.Value, cur.ID, cur.Next[i].Value, cur.Next[i].ID))
			q.PushBack(cur.Next[i])
		}
	}

	t.Log(fmt.Sprintf("%+v", tree))
}

func TestNewLL1Analysis(t *testing.T) {
	ll1, err := NewLL1Analysis(
		`E->v=E'
E'->i`)
	if err != nil {
		fmt.Println(err)
	}

	_, _, ok := ll1.Analysis("vi#")
	if !ok {
		t.Fatalf("no ok")
	}
}
