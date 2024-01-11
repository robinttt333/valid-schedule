package models

type AdjacencyList struct {
	list [][]int
}

func (adj *AdjacencyList) AddEdge(a, b int) {
	adj.list[a] = append(adj.list[a], b)
}

func (adj *AdjacencyList) GetSize() int {
	return len(adj.list)
}

func (adj *AdjacencyList) Get(node int) []int {
	return adj.list[node]
}

func NewAdjacencyList(size int) *AdjacencyList {
	l := [][]int{}
	for i := 0; i < size; i++ {
		l = append(l, []int{})
	}

	return &AdjacencyList{l}
}
