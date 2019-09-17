package heap

//Heap defines binary heap datastructure
type Heap struct {
	MaxHeap bool
	Elems   []int32
}

//Init intialize an int32 heap
func Init(input []int32, maxHeap bool) Heap {
	inLen := len(input)
	heap := Heap{MaxHeap: maxHeap, Elems: make([]int32, 0, (inLen+1)*2)}

	for i := range input {
		heap.Insert(input[i])
	}
	return heap
}

func swapElems(elems []int32, i1, i2 int) {
	tmp := elems[i1]
	elems[i1] = elems[i2]
	elems[i2] = tmp
}

//Fixup moves the last elem of the heap up if needed
func (h *Heap) Fixup() {
	index := len(h.Elems) - 1
	if index < 1 {
		return
	}
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.Elems[index] < h.Elems[parentIndex] && h.MaxHeap ||
			h.Elems[parentIndex] < h.Elems[index] && !h.MaxHeap {
			break
		}
		swapElems(h.Elems, index, parentIndex)
		index = parentIndex
	}
}

//Insert inserts a new node(integer) into the heap
func (h *Heap) Insert(node int32) {
	h.Elems = append(h.Elems, node)
	h.Fixup()
}

//Fixdown deletes the last element in the array and reshape the heap
func (h *Heap) Fixdown() {
	length := len(h.Elems)
	if length == 0 {
		return
	}
	h.Elems = h.Elems[:len(h.Elems)-1]
	length--
	i := 0
	// (i+1)*2 ensures current node has two children
	for (i+1)*2 < length {
		var childIndex int
		if h.MaxHeap {
			if h.Elems[i] >= h.Elems[2*i+1] &&
				h.Elems[i] >= h.Elems[2*i+2] {
				return //done
			}
			if h.Elems[2*i+1] > h.Elems[2*i+2] {
				childIndex = 2*i + 1
			} else {
				childIndex = 2*i + 2
			}
		} else {
			if h.Elems[i] <= h.Elems[2*i+1] &&
				h.Elems[i] <= h.Elems[2*i+2] {
				return //done
			}
			if h.Elems[2*i+1] < h.Elems[2*i+2] {
				childIndex = 2*i + 1
			} else {
				childIndex = 2*i + 2
			}
		}
		swapElems(h.Elems, i, childIndex)
		i = childIndex
	}

	//This check is for one child case in the last possible swap
	if (i*2 + 1) < length {
		if h.Elems[i] < h.Elems[i*2+1] && h.MaxHeap ||
			h.Elems[i] > h.Elems[i*2+1] && !h.MaxHeap {
			swapElems(h.Elems, i, i*2+1)
		}
	}
}

//Remove removes the top element on the heap and reshapes the heap
func (h *Heap) Remove() {
	length := len(h.Elems)
	if length == 0 {
		return
	}
	swapElems(h.Elems, 0, length-1)
	h.Fixdown()
}
