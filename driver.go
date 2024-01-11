package main

import (
	"fmt"
	"strconv"
	"strings"
	"validSchedule/algorithms"
	"validSchedule/models"
	"validSchedule/utils"
)

var PATH = "input.txt"

func main() {
	rows, err := utils.Read(PATH)
	if err != nil {
		panic(err)
	}
	sz, err := strconv.Atoi(rows[0])
	if err != nil || sz == 0 {
		panic(err)
	}
	list := models.NewAdjacencyList(sz)
	for i := 1; i < len(rows); i++ {
		numbers := strings.Split(rows[i], " ")
		a, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}
		list.AddEdge(a, b)
	}
	fmt.Println(algorithms.ScheduleWithDFS(list))
}
