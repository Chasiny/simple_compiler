package vector

import "testing"

func TestNewVector(t *testing.T) {
	v := NewVector()

	v.PushBack(1)
	v.PushBack(4)
	v.PushBack(2)
	v.PushBack(5)
	v.PushBack(4)

	except := []int{1, 4, 2, 5, 4}

	if v.Len() != 5 {
		t.Fail()
	}

	for i := 0; i < v.Len(); i++ {
		t.Log(v.Get(i), except[i])
		if v.Get(i) != except[i] {
			t.Fail()
		}
	}
}
func TestNewVector1(t *testing.T) {
	v := NewVector()

	v.PushBack(1)
	v.PushBack(4)
	v.PushBack(2)
	v.PushBack(5)
	v.PushBack(4)
	v.PushBack(1)
	v.PushBack(4)
	v.PushBack(2)
	v.PushBack(5)
	v.PushBack(4)
	v.PushBack(1)
	v.PushBack(4)
	v.PushBack(2)
	v.PushBack(5)
	v.PushBack(4)

	except := []int{1, 4, 2, 5, 4, 1, 4, 2, 5, 4, 1, 4, 2, 5, 4}

	if v.Len() != 15 {
		t.Fail()
	}

	for i := 0; i < v.Len(); i++ {
		t.Log(v.Get(i), except[i])
		if v.Get(i) != except[i] {
			t.Fail()
		}
	}
}

func TestVector_Remove(t *testing.T) {
	v := NewVector()

	v.PushBack(1)
	v.PushBack(4)
	v.PushBack(2)
	v.PushBack(5)
	v.PushBack(4)

	except := []int{1, 4, 5, 4}

	v.Remove(2)

	if v.Len() != 4 {
		t.Fail()
	}

	for i := 0; i < v.Len(); i++ {
		t.Log(v.Get(i), except[i])
		if v.Get(i) != except[i] {
			t.Fail()
		}
	}
}

func TestVector_Equal(t *testing.T) {
	v1 := NewVector()
	v1.PushBack(1)
	v1.PushBack(4)
	v1.PushBack(2)
	v1.PushBack(5)
	v1.PushBack(4)

	v2 := NewVector()
	v2.PushBack(2)
	v2.PushBack(5)
	v2.PushBack(1)
	v2.PushBack(4)
	v2.PushBack(4)

	if v1.String() != v2.String() {
		t.Fatalf("fail at %+v and %+v", v1, v2)
		return
	}

	v1.Remove(0)
	v1.PushBack(1)
	if v1.String() != v2.String() {
		t.Fatalf("fail at %+v and %+v", v1, v2)
		return
	}

	v1.Remove(3)
	v1.PushBack(1)

	if v1.String() == v2.String() {
		t.Fatalf("fail at %+v and %+v", v1, v2)
		return
	}
}

func TestVectorMap(t *testing.T) {
	v1 := NewVector()
	v1.PushBack(1)
	v1.PushBack(4)
	v1.PushBack(2)
	v1.PushBack(5)
	v1.PushBack(4)

	v2 := NewVector()
	v2.PushBack(1)
	v2.PushBack(4)
	v2.PushBack(2)
	v2.PushBack(5)
	v2.PushBack(4)

	m := make(map[string]struct{})
	m[v1.String()] = struct{}{}
	m[v2.String()] = struct{}{}

	for k, v := range m {
		t.Log(k, v)
	}
}
