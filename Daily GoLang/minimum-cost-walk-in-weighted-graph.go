/*
https://leetcode.com/problems/minimum-cost-walk-in-weighted-graph/description/

There is an undirected weighted graph with n vertices labeled from 0 to n - 1.

You are given the integer n and an array edges, where edges[i] = [ui, vi, wi] indicates that there is an edge between vertices ui and vi with a weight of wi.

A walk on a graph is a sequence of vertices and edges. The walk starts and ends with a vertex, and each edge connects the vertex that comes before it and the vertex that comes after it.
It's important to note that a walk may visit the same edge or vertex more than once.

The cost of a walk starting at node u and ending at node v is defined as the bitwise AND of the weights of the edges traversed during the walk.
In other words, if the sequence of edge weights encountered during the walk is w0, w1, w2, ..., wk, then the cost is calculated as w0 & w1 & w2 & ... & wk, where & denotes the bitwise AND operator.

You are also given a 2D array query, where query[i] = [si, ti].
For each query, you need to find the minimum cost of the walk starting at vertex si and ending at vertex ti. If there exists no such walk, the answer is -1.

Return the array answer, where answer[i] denotes the minimum cost of a walk for query i.

Example:
Input: n = 5, edges = [[0,1,7],[1,3,7],[1,2,1]], query = [[0,3],[3,4]]  Output: [1,-1]
Explanation:
To achieve the cost of 1 in the first query, we need to move on the following edges: 0->1 (weight 7), 1->2 (weight 1), 2->1 (weight 1), 1->3 (weight 7).
In the second query, there is no walk between nodes 3 and 4, so the answer is -1.


Hint 1.
Try using DSU (Disjoin Set Union) to group nodes together based on connectivity.

Hint 2.
Use the bitwise AND operation to merge cost values of two connected components. This ensures that only common set bits remain.

Hint 3.
For each query, if both nodes belong to the same component, output the merged cost. Otherwise, the answer is -1.

Hint 4.
Build a Union-Find structure where each node starts with its cost. For every edge, merge the sets using union-find and update the cost with bitwise AND. Then for each query, if the two nodes have the same representative, return the component’s cost; else , return -1.

We use Union-Find (Disjoint Set) to efficiently merge nodes based on edges.
Each connected component maintains a minimum cost value using bitwise AND.
For each query:
If q[0] == q[1], return 0 since it's the same node.
If they belong to different components, return -1.
Otherwise, return the minimum cost of their component.
*/

//find root for node by using DSU (Disjoint Set Union)
func find(parent []int, node int) int {
    if parent[node] < 0 {
        return node
    }
    parent[node] = find(parent, parent[node])
    return parent[node]
}

func minimumCost(n int, edges [][]int, query [][]int) []int {

    //init. DSU and MinCost
    parent := make([]int, n)
    minCost := make([]int, n)
    for i := range parent {
        parent[i] = -1
        minCost[i] = -1
    }

    //unite edge
    for _, edge := range edges {
        lRoot, rRoot := find(parent, edge[0]), find(parent, edge[1])
        if lRoot != rRoot {
            minCost[lRoot] &= minCost[rRoot]
            parent[rRoot] = lRoot
        }
        minCost[lRoot] &= edge[2]
    }

    result := make([]int, len(query))
    for i, q := range query {
        lRoot, rRoot := find(parent, q[0]), find(parent, q[1])
        if q[0] == q[1] {
            result[i] = 0
        } else if lRoot != rRoot {
            result[i] = -1
        } else {
            result[i] = minCost[lRoot]
        }
    }
    return result
}
