/**
 * @param {number[][]} graph
 * @return {boolean}
 */
var isBipartite = function (graph) {
    const colors = new Array(graph.length).fill(null);
    for (const index in colors) {
        if (colors[index] != null) continue;
        if (!tryColoring(true, index, colors, graph)) return false;
    }
    return true;
};

function tryColoring(color, node, colors, graph) {
    if (colors[node] != null) return colors[node] == color;
    colors[node] = color;
    let flag = true;
    for (const vl of graph[node]) {
        flag &= tryColoring(!color, vl, colors, graph);
        if (!flag) break;
    }
    return flag;
}
/**
 * Time: O(N+E)
 * Space: O(N)
 */
(function () {
    console.log("Case 1: ", isBipartite([[1, 2, 3], [0, 2], [0, 1, 3], [0, 2]]));
    console.log("Case 2: ", isBipartite([[1, 3], [0, 2], [1, 3], [0, 2]]));
})();