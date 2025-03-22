/*
Стандартная реализация поиска в глубину помещает каждую вершину (узел, node) графа в одну из двух категорий:
- Пройденные (Visited).
- Не пройденные (Not Visited).

Цель алгоритма состоит в том, чтобы пометить каждую вершину как “Пройденная”, избегая при этом циклов.

Алгоритм поиска в глубину работает следующим образом:
1) Начните с того, что поместите любую вершину графа на вершину стека.
2) Возьмите верхний элемент стека и добавьте его в список “Пройденных”.
3) Создайте список смежных вершин для этой вершины. Добавьте те вершины, которых нет в списке “Пройденных”, в верх стека.
4) Необходимо повторять шаги 2 и 3, пока стек не станет пустым.
*/

import "fmt"

// Graph представляет граф, используя список смежности
type Graph struct {
	adjList map[int][]int // Список смежности
}

// NewGraph создает новый пустой граф
func NewGraph() *Graph {
	return &Graph{adjList: make(map[int][]int)}
}

// AddEdge добавляет ребро в граф
func (g *Graph) AddEdge(v, w int) {
	g.adjList[v] = append(g.adjList[v], w) // Добавляем w в список смежности v
	g.adjList[w] = append(g.adjList[w], v) // Добавляем v в список смежности w (для неориентированного графа)
}

// DFS - поиск в глубину
func (g *Graph) DFS(start int) {
	// visited будет отслеживать посещенные вершины
	visited := make(map[int]bool)

	fmt.Printf("Starting DFS from vertex %d:\n", start)
	g.dfsHelper(start, visited)
}

// Вспомогательная функция для рекурсивного обхода DFS
// Если сосед еще не посещен, вызываем dfsHelper для этого соседа.
func (g *Graph) dfsHelper(v int, visited map[int]bool) {
	// Отмечаем текущую вершину как посещенную
	visited[v] = true
	fmt.Printf("Visited vertex %d\n", v)

	// Рекурсивно обходим соседей текущей вершины
	for _, neighbor := range g.adjList[v] {
		if !visited[neighbor] {
			fmt.Printf("Going deeper to vertex %d from vertex %d\n", neighbor, v)
			g.dfsHelper(neighbor, visited)
		}
	}
}

func main() {
	graph := NewGraph()

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(2, 6)

	graph.DFS(0)
}
