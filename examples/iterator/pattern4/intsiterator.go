package main

// IntsIterator is a statefule iterator for int slices
type IntsIterator struct {
	index int
	data  []int
}

// NewIntsIterator returns a stateful iterator based on an int slice
func NewIntsIterator(data []int) *IntsIterator {
	return &IntsIterator{data: data, index: -1}
}

// Value returns the value of the current iteration
func (it *IntsIterator) Value() int {
	return it.data[it.index]
}

// Next advances iterator and returns true if there are more iterations
func (it *IntsIterator) Next() bool {
	it.index++
	return it.index < len(it.data)
}
