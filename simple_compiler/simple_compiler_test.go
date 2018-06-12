package simple_compiler

import (
	"testing"
)

func TestSimpleCompiler_Run(t *testing.T) {
	c := NewSimpleCompiler()

	p := `int i;
int sum;
i=i+3;
sum=i+i;
int t;
t=sum+i;`

	v, err := c.Run(p)
	if err != nil {
		t.Fatalf(err.Error())
	}

	t.Log(v)

}
