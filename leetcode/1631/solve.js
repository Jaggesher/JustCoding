/**
 * @param {number[][]} heights
 * @return {number}
 */
var minimumEffortPath = function (heights) {
    const rows = heights.length, columns = heights[0].length;
    const dirs = [[1, 0], [0, 1], [-1, 0], [0, -1]];
    let visited = new Array(rows).fill().map(() => new Array(columns).fill(1_000_000));

    const compare = (a, b) => a.val - b.val;
    const pq = new MinPriorityQueue({ compare }); // Leetcode provides from a package.
    visited[0][0] = 0;
    pq.enqueue({ val: 0, i: 0, j: 0 });

    while (pq.size() > 0) {
        const node = pq.dequeue();
        if (node.val > visited[node.i][node.j]) continue;

        for (const dir of dirs) {
            const u = dir[0] + node.i, v = dir[1] + node.j;
            if (u < 0 || v < 0 || u >= rows || v >= columns) continue;
            const effort = Math.max(Math.abs(heights[u][v] - heights[node.i][node.j]), node.val);
            if (effort < visited[u][v]) {
                visited[u][v] = effort;
                pq.enqueue({ val: effort, i: u, j: v })
            }
        }
    }

    return visited[rows - 1][columns - 1];
};

(function () {
    console.log("Case 1: ", minimumEffortPath([[1, 2, 2], [3, 8, 2], [5, 3, 5]]))
    console.log("Case 2: ", minimumEffortPath([[1, 2, 3], [3, 8, 4], [5, 3, 5]]))
    console.log("Case 3: ", minimumEffortPath([[1, 2, 1, 1, 1], [1, 2, 1, 2, 1], [1, 2, 1, 2, 1], [1, 2, 1, 2, 1], [1, 1, 1, 2, 1]]))
})();