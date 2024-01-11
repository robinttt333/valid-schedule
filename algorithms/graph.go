package algorithms

import (
	"validSchedule/models"
)

func isACyclic(list *models.AdjacencyList, currentNode int, visited []int, stack *[]int) bool {
	visited[currentNode] = 1
	for i := 0; i < len(list.Get(currentNode)); i++ {
		child := list.Get(currentNode)[i]
		// child in current path
		if visited[child] == 1 {
			return false
		}
		// child already visited in some previous path
		if visited[child] == 2 {
			continue
		}
		if isACyclic(list, child, visited, stack) == false {
			return false
		}
	}
	visited[currentNode] = 2
	*stack = append(*stack, currentNode)
	return true
}

func ScheduleWithDFS(list *models.AdjacencyList) []int {
	sz := list.GetSize()
	visited := make([]int, sz)
	for i := 0; i < sz; i++ {
		visited[i] = 0
	}

	st := make([]int, 0)
	for i := 0; i < sz; i++ {
		if visited[i] == 2 {
			continue
		}
		if isACyclic(list, i, visited, &st) == false {
			return []int{}
		}
	}
	return st
}
