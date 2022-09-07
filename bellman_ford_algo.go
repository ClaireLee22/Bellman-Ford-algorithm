package main

import (
	"fmt"
	"math"
)

func findShortestPath(graph [][][]int, start int) []int {
	// set distance to start node itself to 0
	// set the rest of nodes to infinity
	distances := make([]int, len(graph))
	for i := range distances {
		distances[i] = math.MaxInt32
	}
	distances[start] = 0

	// iterate at most n-1 times
	for i := 1; i < len(graph)-1; i++ {
		// if no update occurs, that means there is no negative weight cycle
		if !relaxAndUpdateDistance(graph, distances) {
			fmt.Printf("iteration %d: no update, stop the algorithm\n", i)
			return distances
		}
	}
	if relaxAndUpdateDistance(graph, distances) {
		fmt.Println("detect a negative weight cycle in the graph")
	}
	return distances
}

// visited all edges and relax them
func relaxAndUpdateDistance(graph [][][]int, distances []int) bool {
	update := false
	for currentNode, edges := range graph {
		for _, edge := range edges {
			neighborNode, edgeWeight := edge[0], edge[1]
			newDistanceToNeighbor := distances[currentNode] + edgeWeight
			if newDistanceToNeighbor < distances[neighborNode] {
				update = true
				distances[neighborNode] = newDistanceToNeighbor
			}
		}
	}
	return update
}

var graphWithoutNegativeCycle = [][][]int{
	// node 0
	{
		{1, 2}, {3, 6},
	},
	// node 1
	{
		{3, 1},
	},
	// node 2
	{
		{1, 13},
	},
	// node 3
	{
		{1, 3}, {2, -7},
	},
}

var graphWitNegativeCycle = [][][]int{
	// node 0
	{
		{1, 2}, {3, 6},
	},
	// node 1
	{
		{3, 1},
	},
	// node 2
	{
		{1, 13},
	},
	// node 3
	{
		{1, -3}, {2, -7},
	},
}

func main() {
	fmt.Println("--test graphWithoutNegativeCycle--")
	distancesForgraphWithoutNegativeCycle := findShortestPath(graphWithoutNegativeCycle, 0)
	fmt.Println("the shortest distance in graphWithoutNegativeCycle: distances ", distancesForgraphWithoutNegativeCycle)
	fmt.Println()

	fmt.Println("--test graphWithNegativeCycle--")
	distancesForgraphWithNegativeCycle := findShortestPath(graphWitNegativeCycle, 0)
	fmt.Println("the shortest distance in graphWithNegativeCycle: distances ", distancesForgraphWithNegativeCycle)
	fmt.Println("Above distances are incorrect. It is not possible to produce the shortest path if a graph contains negative weight cycles.")
}

/* output:
--test graphWithoutNegativeCycle--
iteration 2: no update, stop the algorithm
the shortest distance in graphWithoutNegativeCycle: distances  [0 2 -4 3]

--test graphWithNegativeCycle--
detect a negative weight cycle in the graph
the shortest distance in graphWithNegativeCycle: distances  [0 -4 -8 -1]
Above distances are incorrect. It is not possible to produce the shortest path if a graph contains negative weight cycles.

*/
