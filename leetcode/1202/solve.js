/**
 * @param {string} s
 * @param {number[][]} pairs
 * @return {string}
 */
var smallestStringWithSwaps = function (s, pairs) {
    const data = [...s]
    const graph = new Map();
    const visited = new Array(s.length).fill(false);

    for (const [u, v] of pairs) {
        if (!graph.has(u)) graph.set(u, []);
        if (!graph.has(v)) graph.set(v, []);

        graph.get(u).push(v)
        graph.get(v).push(u)
    }

    for (let node = 0; node < visited.length; node++) {
        if (!visited[node] && graph.has(node)) {
            const series = dfs(graph, data, visited, node, { data: [], position: [] })
            series.data.sort((a, b) => a < b ? -1 : a > b ? 1 : 0);
            series.position.sort((a, b) => a - b);
            for (let index = 0; index < series.data.length; index++) {
                data[series.position[index]] = series.data[index];
            }
        }
    }

    return data.join("")
};

function dfs(graph, data, visited, node, resp) {
    if (visited[node]) return resp;
    visited[node] = true;
    resp.data.push(data[node]);
    resp.position.push(node);

    const others = graph.get(node);
    for (const tm of others) {
        dfs(graph, data, visited, tm, resp)
    }
    return resp
}

(function(){
    console.log("Case 1: ", smallestStringWithSwaps("dcab", [[0,3],[1,2]]));
    console.log("Case 2: ", smallestStringWithSwaps("dcab", [[0,3],[1,2],[0,2]]));
    console.log("Case 3: ", smallestStringWithSwaps("cba", [[0,1],[1,2]]));
})()