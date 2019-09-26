package main

import "fmt"

type Node struct {
	Value int
}

// 用于构建结构体切片为最小堆，需要调用down函数
func Init(nodes []Node) {
	for i := (len(nodes)-1) / 2; i >= 0; i-- {
		down(nodes, i)
	}
}

// 需要down（下沉）的元素在切片中的索引为i，n为heap的长度，将该元素下沉到该元素对应的子树合适的位置，从而满足该子树为最小堆的要求
func down(nodes []Node, i int) {
	for cur, left, right, n := i, i*2 + 1, i*2 + 2, len(nodes);
		(left < n && nodes[cur].Value > nodes[left].Value) ||
		(right < n && nodes[cur].Value > nodes[right].Value);
		left, right = cur*2 + 1, cur*2 + 2 {
			if right < n && nodes[right].Value < nodes[left].Value {
				nodes[cur], nodes[right] = nodes[right], nodes[cur]
				cur = right
			} else {
				nodes[cur], nodes[left] = nodes[left], nodes[cur]
				cur = left
			}
		}
}

// 用于保证插入新元素(j为元素的索引,切片末尾插入，堆底插入)的结构体切片之后仍然是一个最小堆
func up(nodes []Node, j int) {
	for cur, parent := j, (j-1)/2;
		parent >= 0 && nodes[cur].Value < nodes[parent].Value;
		cur, parent = parent, (parent-1)/2 {
			nodes[cur], nodes[parent] = nodes[parent], nodes[cur]
		}
}

// 弹出最小元素，并保证弹出后的结构体切片仍然是一个最小堆，第一个返回值是弹出的节点的信息，第二个参数是Pop操作后得到的新的结构体切片
func Pop(nodes []Node) (Node, []Node) {
	if len(nodes) == 0 {return Node{}, nodes}
	min_node := nodes[0]
	nodes[0], nodes[len(nodes)-1] = nodes[len(nodes)-1], nodes[0]
	nodes = nodes[:len(nodes)-1]
	down(nodes, 0)
	return min_node, nodes
}

// 保证插入新元素时，结构体切片仍然是一个最小堆，需要调用up函数
func Push(node Node, nodes []Node) []Node {
	nodes = append(nodes, node)
	up(nodes, len(nodes)-1)
	return nodes
}

// 移除切片中指定索引的元素，保证移除后结构体切片仍然是一个最小堆
func Remove(nodes []Node, node Node) []Node {
	for i, n := range nodes {
		if n == node {
			nodes[i], nodes[len(nodes)-1] = nodes[len(nodes)-1], nodes[i]
			nodes = nodes[:len(nodes)-1]
			up(nodes, i)
			down(nodes, i)
			break
		}
	}
	return nodes
}

func main() {
	a := []Node{Node{1}, Node{3}, Node{2}, {0}}
	fmt.Println(a)
	Init(a)
	fmt.Println(a)
}