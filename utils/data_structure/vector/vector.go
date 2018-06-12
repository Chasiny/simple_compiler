package vector

import (
	"sort"
	"strconv"
)

type Vector struct {
	data []int

	first int
	last  int
	end   int
}

func NewVector() (v *Vector) {
	return &Vector{
		data:  make([]int, 10),
		first: 0,
		last:  0,
		end:   10,
	}
}

func (v *Vector) PushBack(in int) {
	if v.last == v.end {
		newData := make([]int, v.end+10)
		v.end = v.end + 10

		for i := range v.data {
			newData[i] = v.data[i]
		}
		v.data = newData
	}

	v.data[v.last] = in
	v.last++
}

func (v *Vector) Len() (res int) {
	return v.last - v.first
}

func (v *Vector) Cap() (res int) {
	return v.end - v.first
}

func (v *Vector) Get(index int) (res int) {
	if index >= v.last {
		panic("vector get error: index out of size")
		return -1
	}

	return v.data[index]
}

func (v *Vector) Remove(index int) {
	if index >= v.last {
		panic("vector get error: index out of size")
		return
	}

	for i := index; i < v.last-1; i++ {
		v.data[i] = v.data[i+1]
	}

	v.last--
}

func (v *Vector) String() (res string) {
	sort.Ints(v.data[:v.last])

	res += "{"
	for i := 0; i < v.last-1; i++ {
		res += strconv.Itoa(v.data[i]) + ","
	}

	if v.last-1 >= 0 {
		res += strconv.Itoa(v.data[v.last-1])
	}

	res += "}"
	return
}

func (v *Vector) AddVector(addV *Vector) {
	for i := 0; i < addV.Len(); i++ {
		v.PushBack(addV.Get(i))
	}
}
