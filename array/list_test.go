package array

import (
	"container/list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	dataList := list.New()

	dataList.PushBack("adit")
	dataList.PushBack("diqi")
	dataList.PushBack("ichsan")

	n := 1
	for i := dataList.Front(); i != nil; i = i.Next() {
		fmt.Printf("no antrian ke %d, atas nama %s\n", n, i.Value)
		n++
	}
}
