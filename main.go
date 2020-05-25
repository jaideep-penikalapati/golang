package main

import (
	"fmt"

	"github.com/jaideep-penikalapati/golang/datatype/minheap"
)

func main() {
	heap := minheap.NewHeap()

	fmt.Println(heap)

	heap, err := minheap.PopulateHeap(heap, []int{10, 34, 30, 5, 9, 50})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("successful insertion")
	}

	fmt.Println(heap)
	fmt.Println(minheap.Peek(heap))

	val, err := minheap.Pop(&heap)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("successful deletion")
		fmt.Println(val, heap)
	}

	fmt.Println(heap)
}
