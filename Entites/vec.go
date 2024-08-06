package entities

import (
	"errors"
	"log"
)

type Vec[T any] struct {
	Arr   []T
	Count int
	Size  int
}

func (lst Vec[T]) At(index int) *T {
	if index < lst.Size {
		return &lst.Arr[index]
	}
	log.Fatal("error: index out of range")
	return nil
}
func (lst Vec[T]) GetCount() int {
	return lst.Count
}
func (lst Vec[T]) GetLength() int {
	return len(lst.Arr)
}
func (lst Vec[T]) IsEmpty() bool {
	return lst.Size == 0
}
func (lst *Vec[T]) Empty() {
	lst.Size = 0
	lst.Arr = []T{}
}

func (Vec[T]) Init() Vec[T] {
	return Vec[T]{Arr: []T{}, Count: 16, Size: 0}
}
func (lst *Vec[T]) PushBack(elemnt T) {
	if lst.Count <= lst.Size {
		lst.Count *= 2
	}

	lst.Arr = append(lst.Arr, elemnt)
	lst.Size++

}
func (lst *Vec[T]) PopBack() {
	lst.Size--
	lst.Arr = lst.Arr[:lst.Size]

}

func (lst *Vec[T]) PopIndex(index int) error {
	if index < lst.GetLength() && index >= 0 {
		lst.Arr = append(lst.Arr[:index], lst.Arr[(index+1):]...)
		lst.Size--
		return nil
	}
	return errors.New("index out of range")
}

/*func (lst *Vec[T]) Contains(elemnt T) (bool, int) {
	for i, e := range lst.Arr {
		if e == elemnt {
			return true, i
		}
	}
	return false, -1
}
func (lst *Vec[T]) PopElemnt(elemnt T) error {
	if b, i := lst.Contains(elemnt); b {
		lst.PopIndex(i)
		return nil
	}
	return nil
}
*/
