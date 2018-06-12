package queue

import "testing"

func TestQueue_Len(t *testing.T) {
	q := NewQueue()
	q.PushBack(1)
	q.PushBack(1)
	q.PushBack(1)
	q.PushBack(1)

	if q.Len() != 4 {
		t.Log("len", q.Len())
		t.Fail()
	}
}

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.PushBack(1)
	q.PushBack(2)
	q.PushBack(3)

	except := []int{
		1, 2, 3,
	}

	i := 0
	for q.Len() > 0 {
		ans := q.Pop().(int)
		t.Log(except[i], ans)
		if ans != except[i] {
			t.Fail()
			return
		}

		i++
	}
}
