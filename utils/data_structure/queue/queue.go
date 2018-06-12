package queue

type data struct {
	data interface{}
	next *data
}

type Queue struct {
	start *data
	last  *data
	len   int
}

func NewQueue() *Queue {
	return &Queue{
		start: nil,
		last:  nil,
		len:   0,
	}
}

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) PushBack(i interface{}) {
	newData := &data{
		data: i,
		next: nil,
	}
	q.len++

	if q.start == nil {
		q.start = newData
		q.last = newData

		return
	}

	q.last.next = newData
	q.last = newData
}

func (q *Queue) Pop() (i interface{}) {
	if q.Len() <= 0 {
		panic("queue len <= 0")
		return nil
	}

	q.len--

	i = q.start.data

	q.start = q.start.next
	if q.start == nil {
		q.last = nil
	}

	return
}
