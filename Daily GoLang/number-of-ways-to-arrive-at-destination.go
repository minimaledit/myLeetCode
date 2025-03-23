/*
https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/description/

You are in a city that consists of n intersections numbered from 0 to n - 1 with bi-directional roads between some intersections.
The inputs are generated such that you can reach any intersection from any other intersection and that there is at most one road between any two intersections.

You are given an integer n and a 2D integer array roads where roads[i] = [ui, vi, timei] means that there is a road between intersections ui and vi that takes timei minutes to travel.
You want to know in how many ways you can travel from intersection 0 to intersection n - 1 in the shortest amount of time.

Return the number of ways you can arrive at your destination in the shortest amount of time. Since the answer may be large, return it modulo 10^9 + 7.

Example 1:
Input: n = 7, roads = [[0,6,7],[0,1,2],[1,2,3],[1,3,3],[6,3,3],[3,5,1],[6,5,1],[2,5,1],[0,4,5],[4,6,2]]  Output: 4
Explanation: The shortest amount of time it takes to go from intersection 0 to intersection 6 is 7 minutes.
The four ways to get there in 7 minutes are:
- 0 ➝ 6
- 0 ➝ 4 ➝ 6
- 0 ➝ 1 ➝ 2 ➝ 5 ➝ 6
- 0 ➝ 1 ➝ 3 ➝ 5 ➝ 6
*/

import (
	"container/heap"
	"math"
)

// Graph edge
type Item struct {
	node int
	cost int
}


type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // Preventing memory leaks
	*pq = old[0 : n-1]
	return item
}

func countPaths(n int, roads [][]int) int {
	const MOD = 1_000_000_007

  // Initialization arrays for Dijkstra
	dist := make([]int, n)
	ways := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
  // Default values
	dist[0] = 0
	ways[0] = 1

  // Build the Graph
	graph := make([][][2]int, n)
	for _, road := range roads {
		u, v, veight := road[0], road[1], road[2]
		graph[u] = append(graph[u], [2]int{v, veight})
		graph[v] = append(graph[v], [2]int{u, veight})
	}

  // Algorithm start
	pq := &PriorityQueue{}

	heap.Init(pq)
	heap.Push(pq, &Item{node: 0, cost: 0})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Item)

		if curr.cost > dist[curr.node] {
			continue
		}

		for _, neighbor := range graph[curr.node] {
			nextNode, veight := neighbor[0], neighbor[1]
			newDist := curr.cost + veight

			if newDist < dist[nextNode] {
				dist[nextNode] = newDist
				ways[nextNode] = ways[curr.node]
				heap.Push(pq, &Item{node: nextNode, cost: newDist})
			} else if newDist == dist[nextNode] {
				ways[nextNode] = (ways[curr.node] + ways[nextNode]) % MOD
			}
		}
	}
	return ways[n-1] % MOD
}
