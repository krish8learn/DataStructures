package src

//MaxHeap sruct which indicate to the array that holds the values of heap's elements
type Heap struct {
	holder []int
}

//Insert adds an element to the heap
func (h *Heap) Insert(element int) {
	h.holder = append(h.holder, element)
	h.BTHeapify(len(h.holder) - 1) //we are adding element at the bottom end, so heapify from bottom to top
}

//Extract the largest key and Heapify
func (h *Heap) Extract() int {
	if len(h.holder) == 0 {
		return -1
	}
	ouput := h.holder[0]
	lastIndex := len(h.holder) - 1
	h.holder[0] = h.holder[lastIndex]
	h.holder = h.holder[:lastIndex]
	h.TBHeapify(0)
	return ouput
}

//Heapify bottom to top approach
func (h *Heap) BTHeapify(index int) {
	for h.holder[Parent(index)] < h.holder[index] {
		h.swap(Parent(index), index)
		index = Parent(index)
	}
}

//Heapify top to bottom approach
func (h *Heap) TBHeapify(index int) {
	l, r := RightIndex(index), LeftIndex(index)
	comp_element := 0
	lastIndex := len(h.holder) - 1

	for l <= lastIndex {
		if l == lastIndex {
			comp_element = l
		} else if h.holder[l] > h.holder[r] {
			comp_element = l
		} else if h.holder[l] < h.holder[r] {
			comp_element = r
		}

		if h.holder[comp_element] > h.holder[index] {
			h.swap(index, comp_element)
			index = comp_element
			l, r = RightIndex(index), LeftIndex(index)
		} else {
			return
		}
	}
}

//getting parent index
func Parent(i int) int {
	return (i - 1) / 2
}

//getting child Index left
func LeftIndex(i int) int {
	return (2 * i) + 1
}

//getting child index right
func RightIndex(i int) int {
	return (2 * i) + 2
}

//swap two values
func (h *Heap) swap(i1, i2 int) {
	h.holder[i1], h.holder[i2] = h.holder[i2], h.holder[i1]
}
