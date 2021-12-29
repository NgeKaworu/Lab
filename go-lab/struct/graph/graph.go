package graph

import (
	"fmt"
	"strconv"
)

// Graph 图结构
type Graph struct {
	start *Vertex
}

// Vertex 顶点
type Vertex struct {
	key int             // 键
	in  map[*Vertex]int // 入度
	out map[*Vertex]int // 出度
}

// NewVertex 工厂方法
func NewVertex(k int) *Vertex {
	return &Vertex{key: k, in: map[*Vertex]int{}, out: map[*Vertex]int{}}
}

// BFS 打印
func (g *Graph) String() (res string) {
	// 标记去过的地方
	memo := map[*Vertex]bool{g.start: true}
	queue := []*Vertex{g.start}
	for len(queue) != 0 {
		dequeue := queue[0]
		queue = queue[1:]
		for vertex, weight := range dequeue.out {
			if memo[vertex] == false {
				queue = append(queue, vertex)
				memo[vertex] = true
			}
			res += fmt.Sprintf(" %v --(%v)--> %v ", dequeue.key, weight, vertex.key)
		}
		res += "\n"
	}
	return
}

// FindAllPath 返回所有路径 回溯实现
func (v *Vertex) FindAllPath(end int) []string {
	res := make([]string, 0)

	var backtrack func(v *Vertex, memo map[*Vertex]bool)
	backtrack = func(v *Vertex, memo map[*Vertex]bool) {
		if v.key == end {
			var tmp string
			var pre *Vertex
			for v := range memo {
				if pre != nil {
					tmp += fmt.Sprintf(" --(%v)--> ", pre.out[v])
				}
				pre = v
				tmp += strconv.Itoa(v.key)
			}
			if pre != nil {
				tmp += fmt.Sprintf(" --(%v)--> %v", pre.out[v], v.key)
			}
			res = append(res, tmp)
			return
		}

		for k := range v.out {
			if memo[k] {
				continue
			}
			memo[v] = true
			backtrack(k, memo)
			delete(memo, v)
		}

	}

	m := make(map[*Vertex]bool)
	backtrack(v, m)
	return res
}

// ShortestPath 返回所有路径
func (v *Vertex) ShortestPath(end int) string {
	// 这个用FindAllPath方法算出最小路径即可, 但是会存在重复子问题
	// 这个时候需要dp一下, dp算法还在想, 想好更新
	return ""
}

/*
Connect 连接两个顶点
degree 度 —— 0: 出度; 1: 入度; 2:双向
weight 权重
*/
func (v *Vertex) Connect(n *Vertex, degree, weight int) *Vertex {
	switch degree {
	case 0:
		v.out[n] = weight
		n.in[v] = weight
		break
	case 1:
		v.in[n] = weight
		n.out[v] = weight
		break
	case 2:
		v.out[n] = weight
		v.in[n] = weight
		n.out[v] = weight
		n.in[v] = weight
		break

	}
	return v
}

/*
Disconnect 取消连接两个顶点
degree 度 —— 0: 出度; 1: 入度; 2:双向
*/
func (v *Vertex) Disconnect(n *Vertex, degree int) *Vertex {
	switch degree {
	case 0:
		delete(v.out, n)
		delete(n.in, v)
		break
	case 1:
		delete(v.in, n)
		delete(n.out, v)
		break
	case 2:
		delete(v.out, n)
		delete(n.out, v)
		delete(v.in, n)
		delete(n.in, v)
		break

	}
	return v
}
