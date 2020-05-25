package minheap

import (
	"errors"
)

// NewHeap returns new empty heap
func NewHeap() []int {
	return make([]int, 0)
}

// PopulateHeap extends heap from slice of ints
func PopulateHeap(heap, arr []int) ([]int, error) {
	if heap == nil {
		return []int{}, errors.New("input heap in nil")
	}

	if len(arr) == 0 {
		return heap, nil
	}

	for _, val := range arr {
		heap = append(heap, val)
		heapifyUp(heap)
	}
	return heap, nil
}

// Insert element in the heap
func Insert(heap []int, element int) ([]int, error) {
	if heap == nil {
		return []int{}, errors.New("input heap in nil")
	}
	heap = append(heap, element)
	heapifyUp(heap)
	return heap, nil
}

// Peek the min element of heap
func Peek(heap []int) (int, error) {
	if heap == nil || len(heap) == 0 {
		return 0, errors.New("heap is empty")
	}
	return heap[0], nil
}

// Pop the min element from heap
func Pop(heap *[]int) (int, error) {
	if heap == nil || len(*heap) == 0 {
		return 0, errors.New("heap is empty")
	}
	val := (*heap)[0]
	(*heap)[0] = (*heap)[len(*heap)-1]
	*heap = (*heap)[:len(*heap)-1]
	heapifyDown(*heap)
	return val, nil
}

func heapifyDown(heap []int) {
	index := 0
	len := len(heap)
	for {
		if !hasLeftChild(index, len) {
			break
		} else if !hasRightChild(index, len) {
			if val, pos := getLeftChild(heap, index); val < heap[index] {
				swap(heap, pos, index)
				index = pos
			} else {
				break
			}
		} else {
			val, pos, _ := getSamllerChildWithIndex(heap, index)
			if val < heap[index] {
				swap(heap, pos, index)
				index = pos
			} else {
				break
			}
		}
	}
}

func heapifyUp(heap []int) {
	index := len(heap) - 1
	for {
		if !hasParent(index) {
			break
		} else if val, pos := getParent(heap, index); val > heap[index] {
			swap(heap, pos, index)
			index = pos
		} else {
			break
		}
	}
}

func getParent(heap []int, index int) (int, int) {
	return heap[(index-1)/2], (index - 1) / 2
}

func getLeftChild(heap []int, index int) (int, int) {
	return heap[2*index+1], 2*index + 1
}

func getRightChild(heap []int, index int) (int, int) {
	return heap[2*index+2], 2*index + 2
}

func swap(heap []int, i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func getSamllerChildWithIndex(heap []int, index int) (int, int, error) {
	len := len(heap)
	if !hasLeftChild(index, len) {
		return 0, 0, errors.New("No children present")
	} else if !hasRightChild(index, len) {
		val, pos := getLeftChild(heap, index)
		return val, pos, nil
	} else {
		lval, lpos := getLeftChild(heap, index)
		rval, rpos := getRightChild(heap, index)
		if lval < rval {
			return lval, lpos, nil
		}
		return rval, rpos, nil
	}
}

func hasParent(index int) bool {
	return (index-1)/2 >= 0
}

func hasLeftChild(index, len int) bool {
	return (2*index)+1 < len
}

func hasRightChild(index, len int) bool {
	return (2*index)+2 < len
}
