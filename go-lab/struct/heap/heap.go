package heap

import "strconv"

// Heap 堆
type Heap struct {
	heap []int
}

// Heapify 堆化
func Heapify(arr []int) *Heap {
	h := new(Heap)
	for _, v := range arr {
		h.Insert(v)
	}
	return h
}

// Len 长度
func (h *Heap) Len() int {
	return len(h.heap)
}

func (h *Heap) String() (s string) {
	t, l := 1, 1
	for k, v := range h.heap {
		s += " " + strconv.Itoa(v)
		if k == (l - 1) {
			s += "\n"
			t <<= 1
			l += t
		}
	}
	return
}

// Insert 插入
func (h *Heap) Insert(val int) *Heap {
	h.heap = append(h.heap, val)
	h.swim(h.Len() - 1)
	return h
}

// Pop 弹出
func (h *Heap) Pop() (res int) {
	res = h.heap[0]
	h.heap[0] = h.heap[h.Len()-1]
	h.heap = h.heap[:h.Len()-1]
	h.sink(0)
	return
}

// swap 交换位置
func (h *Heap) swap(p1, p2 int) *Heap {
	h.heap[p1], h.heap[p2] = h.heap[p2], h.heap[p1]
	return h
}

// Swim 上浮
func (h *Heap) swim(pos int) *Heap {
	for pos > 0 && h.heap[pos] < h.heap[parent(pos)] {
		// 如果第 k 个元素不是堆顶且比上层小
		// 将 k 换上去
		h.swap(pos, parent(pos))
		pos = parent(pos)
	}
	return h
}

// Sink 下沉
func (h *Heap) sink(pos int) *Heap {
	l := h.Len()
	// 如果沉到堆底，就沉不下去了
	for left(pos) < l {
		// 先假设左边节点较小
		t := left(pos)
		// 如果右边节点存在，比一下大小
		if right(pos) < l && h.heap[t] > h.heap[right(pos)] {
			t = right(pos)
		}
		// 结点 k 比俩孩子都小，就不必下沉了
		if h.heap[t] > h.heap[pos] {
			break
		}
		// 否则，不符合最小堆的结构，下沉 k 结点
		h.swap(pos, t)
		pos = t
	}

	return h
}

// parent 返回父节点
func parent(pos int) int {
	return (pos - 1) / 2
}

// left 返回左节点
func left(pos int) int {
	return pos*2 + 1

}

// right 返回左节点
func right(pos int) int {
	return pos*2 + 2

}
