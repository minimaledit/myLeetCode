/*
https://leetcode.com/problems/count-the-number-of-complete-components/description/

You are given an integer n. There is an undirected graph with n vertices, numbered from 0 to n - 1. 
You are given a 2D integer array edges where edges[i] = [ai, bi] denotes that there exists an undirected edge connecting vertices ai and bi.

Return the number of complete connected components of the graph.

A connected component is a subgraph of a graph in which there exists a path between any two vertices, and no vertex of the subgraph shares an edge with a vertex outside of the subgraph.

A connected component is said to be complete if there exists an edge between every pair of its vertices.

Example:
Input: n = 6, edges = [[0,1],[0,2],[1,2],[3,4]]  Output: 3
*/

type Graph struct {
	adjacencyList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(firstVertex, secondVertex int) {
	g.adjacencyList[firstVertex] = append(g.adjacencyList[firstVertex], secondVertex)
	g.adjacencyList[secondVertex] = append(g.adjacencyList[secondVertex], firstVertex)
}

// DFS (Depth-First Search)
func (g *Graph) DFS(startVertex int, visited map[int]bool) (vertexCount int, edgeCount int) {
	return g.traverseComponent(startVertex, visited)
}

// traverseComponent рекурсивно обходит компоненту связности, подсчитывая вершины и ребра
func (g *Graph) traverseComponent(currentVertex int, visited map[int]bool) (vertexCount int, edgeCount int) {
	// Помечаем текущую вершину как посещенную
	visited[currentVertex] = true
	vertexCount = 1

	// Обходим всех соседей текущей вершины
	for _, neighbor := range g.adjacencyList[currentVertex] {
		edgeCount++ // Учитываем каждое направление ребра
		
		if !visited[neighbor] {
			// Рекурсивно обходим непосещенных соседей
			childVertices, childEdges := g.traverseComponent(neighbor, visited)
			vertexCount += childVertices
			edgeCount += childEdges
		}
	}
	
	return vertexCount, edgeCount
}

// Проверка на то, является ли компонента полным графом
func isCompleteComponent(vertices, edges int) bool {
	// Для полного графа количество ребер = n*(n-1)/2
	maxEdges := vertices * (vertices - 1) / 2
	return edges == maxEdges
}

func countCompleteComponents(totalVertices int, edges [][]int) int {
	graph := NewGraph()
	visited := make(map[int]bool)
	
	// Добавляем все ребра в граф
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1])
	}

	completeComponents := 0
	for vertex := 0; vertex < totalVertices; vertex++ {
		if !visited[vertex] {
			vertices, edges := graph.DFS(vertex, visited)
			edges /= 2 // Корректируем счетчик ребер (каждое ребро считалось дважды)
			
			if isCompleteComponent(vertices, edges) {
				completeComponents++
			}
		}
	}
	
	return completeComponents
}
