package lexical_analysis

import (
	"fmt"
	"testing"
)

func TestNewNFA(t *testing.T) {
	n := NewNFA("(ab)*(a*|b*)(ba)*")
	fmt.Println(fmt.Sprintf("%+v\n%+v", n.Start.FindClosure(), n.CharSet))

	t.Log(n.String())
}
