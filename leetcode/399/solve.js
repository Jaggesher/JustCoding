/**
 * @param {string[][]} equations
 * @param {number[]} values
 * @param {string[][]} queries
 * @return {number[]}
 */
var calcEquation = function (equations, values, queries) {
    const graph = constructGraph(equations, values);
    const ans = new Array(queries.length).fill(-1);
    for (const index in queries) {
        const [A, B] = queries[index];
        if (!graph.has(A) || !graph.has(B)) continue;
        ans[index] = resolvePath(A, 1, B, graph, new Map());
    }
    return ans
};

function constructGraph(equations, values) {
    const newNode = (node, weight) => { return { node, weight } };
    const graph = new Map();
    for (const index in equations) {
        const [A, B] = equations[index];
        if (!graph.has(A)) graph.set(A, []);
        if (!graph.has(B)) graph.set(B, []);
        graph.get(A).push(newNode(B, values[index]));
        if (A != B) graph.get(B).push(newNode(A, 1 / values[index]));
    }
    return graph;
};


function resolvePath(node, soFar, target, graph, visited) {
    if (node == target) return soFar
    visited.set(node, true)
    const nextNodes = graph.get(node);
    for (const vl of nextNodes) {
        if (!visited.has(vl.node)) {
            const tm = resolvePath(vl.node, soFar * vl.weight, target, graph, visited);
            if (tm != -1) return tm
        }
    }
    return -1;
};

(function () {
    console.log("Case 1: ", calcEquation([["a", "b"], ["b", "c"]], [2.0, 3.0], [["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x", "x"]]));
    console.log("Case 2: ", calcEquation([["a", "b"], ["b", "c"], ["bc", "cd"]], [1.5, 2.5, 5.0], [["a", "c"], ["c", "b"], ["bc", "cd"], ["cd", "bc"]]));
    console.log("Case 3: ", calcEquation([["a", "b"]], [0.5], [["a", "b"], ["b", "a"], ["a", "c"], ["x", "y"]]));
})();