def findShortestPath(graph, start):
    # set distance to start node itself to 0
    # set the rest of nodes to infinity
    distances = [float("inf")]*len(graph)
    distances[start] = 0

    # iterate at most n-1 times
    for i in range(1, len(graph)-1):
        # if no update occurs, that means there is no negative weight cycle
        if not relaxAndUpdateDistance(graph, distances):
            print("iteration %d: no update, stop the algorithm" % (i))
            return distances
    # for the nth iteration, detect if there is negative weight cycle in the graph
    if relaxAndUpdateDistance(graph, distances):
        print("detect a negative weight cycle in the graph")
    return distances


# visited all edges and relax them
def relaxAndUpdateDistance(graph, distances):
    updated = False
    for currentNode, edges in enumerate(graph):
        for neighborNode, edgeWeight in edges:
            newDistanceToNeighbor = distances[currentNode] + edgeWeight
            if newDistanceToNeighbor < distances[neighborNode]:
                updated = True
                distances[neighborNode] = newDistanceToNeighbor 
    return updated

graphWithoutNegativeCycle = [
    # node 0 
    [[1, 2], [3, 6]], # [node, edge_weight]
    # node 1
    [[3, 1]],
    # node 2
    [[1, 13]],
    # node3
    [[1, 3], [2, -7]]
]

graphWithNegativeCycle = [
    # node 0 
    [[1, 2], [3, 6]], # [node, edge_weight]
    # node 1
    [[3, 1]],
    # node 2
    [[1, 13]],
    # node3
    [[1, -3], [2, -7]]
]

if __name__ == "__main__":
    print("--test graphWithoutNegativeCycle--")
    distancesForgraphWithoutNegativeCycle = findShortestPath(graphWithoutNegativeCycle, 0)
    print("the shortest distance in graphWithoutNegativeCycle: distances",  distancesForgraphWithoutNegativeCycle)
    print("\n")

    print("--test graphWithNegativeCycle--")
    distancesForgraphWithNegativeCycle = findShortestPath(graphWithNegativeCycle, 0)
    print("the shortest distance in graphWithNegativeCycle: distances", distancesForgraphWithNegativeCycle)
    print("Above distances are incorrect. It is not possible to produce the shortest path if a graph contains negative weight cycles.")

"""
output:
--test graphWithoutNegativeCycle--
iteration 2: no update, stop the algorithm
('the shortest distance in graphWithoutNegativeCycle: distances', [0, 2, -4, 3])


--test graphWithNegativeCycle--
detect a negative weight cycle in the graph
('the shortest distance in graphWithNegativeCycle: distances', [0, -4, -8, -1])
Above distances are incorrect. It is not possible to produce the shortest path if a graph contains negative weight cycles.

"""