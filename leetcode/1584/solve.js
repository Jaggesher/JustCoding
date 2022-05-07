/**
 * @param {number[][]} points
 * @return {number}
 */
var minCostConnectPoints = function (points) {
    const parents = new Array(points.length);
    for (let index = 0; index < parents.length; index++) {
        parents[index] = index;
    }

    const edges = makeEdges(points);
    edges.sort((a, b) => a.weight - b.weight);

    return constructTree(parents, edges);
};

function constructTree(parents, edges) {
    let cost = 0;

    for (const edge of edges) {
        const u = findParent(parents, edge.i), v = findParent(parents, edge.j);
        if (u == v) continue;
        cost += edge.weight;
        parents[v] = u;
    }

    return cost;
};


function makeEdges(points) {
    const distance = (a, b) => Math.abs(a[0] - b[0]) + Math.abs(a[1] - b[1]);
    const edges = [];
    for (let i = 0; i < points.length; i++) {
        for (let j = i + 1; j < points.length; j++) {
            edges.push({ i, j, weight: distance(points[i], points[j]) });
        }
    }
    return edges;
};

function findParent(parents, node) {
    while (parents[node] != node) {
        node = parents[node];
    }
    return node;
};

(function () {
    console.log("Case 1: ", minCostConnectPoints([[0, 0], [2, 2], [3, 10], [5, 2], [7, 0]])); // 20
    console.log("Case 2: ", minCostConnectPoints([[3, 12], [-2, 5], [-4, 1]])); // 18
})();