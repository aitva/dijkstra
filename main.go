package main

import (
	"fmt"
	"math"
)

const nodeCount = 5

var board = [][]int{
	{0, 1, 5, 2, 10},
	{1, 0, 3, 1, 5},
	{5, 3, 0, 6, 6},
	{2, 1, 6, 0, 5},
	{10, 5, 6, 5, 0},
}

var nodes = map[int]*node{
	0: {ID: 0},
	1: {ID: 1},
	2: {ID: 2},
	3: {ID: 3},
	4: {ID: 4},
}

type node struct {
	ID int
}

func dijkstra(src int) (dist, prev []int) {
	unvisited := make(map[int]*node)
	for k, v := range nodes {
		unvisited[k] = v
	}
	dist = make([]int, nodeCount)
	prev = make([]int, nodeCount)
	for i := 0; i < nodeCount; i++ {
		dist[i] = math.MaxInt32
		prev[i] = -1
	}
	dist[src] = 0
	for len(unvisited) > 0 {
		min := -1
		for i := range unvisited {
			if min == -1 || dist[i] < dist[min] {
				min = i
			}
		}
		u := unvisited[min]
		delete(unvisited, min)

		// fmt.Fprintln(os.Stderr, "min:", min, "u:", u)
		for v := range board[u.ID] {
			alt := dist[u.ID] + board[u.ID][v]
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u.ID
			}
		}
	}
	return
}

func main() {
	fmt.Println("Here is a graf:")
	fmt.Println(board)
	dist, prev := dijkstra(0)
	fmt.Println("Here are the distances from node 0:")
	fmt.Println(dist)
	fmt.Println("Here are the previous to node 0:")
	fmt.Println(prev)
}
