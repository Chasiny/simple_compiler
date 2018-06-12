package lexical_analysis

import (
	"fmt"
	"testing"
)

func TestNewDFA(t *testing.T) {
	n := NewNFA("abas*")
	d := NewDFA(n)
	t.Log(fmt.Sprintf("%+v", d.String()))
	t.Log(d.Accept("ab"))
}
