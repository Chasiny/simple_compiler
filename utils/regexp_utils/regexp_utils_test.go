package regexp_utils

import (
	"fmt"
	"testing"
)

func TestEraseBrackets(t *testing.T) {
	tCase := []string{
		"((hello)(hi))",
		"((hellohi))",
		"(((this)(is(test))))",
	}

	tExcept := []string{
		"(hello)(hi)",
		"hellohi",
		"(this)(is(test))",
	}

	for i := range tCase {
		if res := EraseBrackets(tCase[i]); res != tExcept[i] {
			t.Log(res)
			t.Fail()
		}
	}
}

func TestSpitAnd(t *testing.T) {
	tCase := []string{
		"(hello)(hi)",
		"this(is(test))",
		"(th)i(s)(is(test))",
		"th*is(is(test))(hello)*",
		"(hello)(hi)this",
	}

	tExcept := [][]string{
		{
			"hello",
			"hi",
		}, {
			"t",
			"h",
			"i",
			"s",
			"is(test)",
		}, {
			"th",
			"i",
			"s",
			"is(test)",
		}, {
			"t",
			"h*",
			"i",
			"s",
			"is(test)",
			"(hello)*",
		}, {
			"hello",
			"hi",
			"t",
			"h",
			"i",
			"s",
		},
	}

	for i := range tCase {
		res := SpitAnd(tCase[i])
		if len(res) != len(tExcept[i]) {
			t.Log(fmt.Sprintf("%+v %d", res, len(res)))
			t.Fail()
			return
		}

		for j := range res {
			if res[j] != tExcept[i][j] {
				t.Log(fmt.Sprintf("%+v %d", res, len(res)))
				t.Fail()
				return
			}
		}
	}
}

func TestSpitOr(t *testing.T) {
	tCase := []string{
		"(hello)|(hi)",
		"(hello)*|hello(hi)*",
		"this|is|(me|you|(it)*)|him",
	}

	tExcept := [][]string{
		{
			"hello",
			"hi",
		},
		{
			"(hello)*",
			"hello(hi)*",
		},
		{
			"this",
			"is",
			"me|you|(it)*",
			"him",
		},
	}

	for i := range tCase {
		res := SpitOr(tCase[i])
		if len(res) != len(tExcept[i]) {
			t.Log(fmt.Sprintf("%+v %d", res, len(res)))
			t.Fail()
			return
		}

		for j := range res {
			if res[j] != tExcept[i][j] {
				t.Log(fmt.Sprintf("%+v %d", res, len(res)))
				t.Fail()
				return
			}
		}
	}
}

func TestSpitClosure(t *testing.T) {
	tCase := []string{
		"((hello)(hi))*",
		"((hellohi))*",
		"this*",
		"s*",
		"(hello)(hi)*",
	}

	tExcept := []string{
		"(hello)(hi)",
		"hellohi",
		"this*",
		"s",
		"(hello)(hi)*",
	}

	for i := range tCase {
		if res, _ := SpitClosure(tCase[i]); res != tExcept[i] {
			t.Log(res)
			t.Fail()
		}
	}
}
