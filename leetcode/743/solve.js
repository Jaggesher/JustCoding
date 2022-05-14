/**
 * @param {number[][]} times
 * @param {number} n
 * @param {number} k
 * @return {number}
 */
var networkDelayTime = function (times, n, k) {
    const compare = (a, b) => a.weight - b.weight;
    const pq = new MinPriorityQueue({ compare });
    pq.enqueue(new Node(k, 0));
    const graph = constructGraph(n, times), visited = new Array(n + 1).fill(-1);

    while (pq.size()) {
        const node = pq.dequeue();
        if (visited[node.node] != -1 && visited[node.node] <= node.weight) {
            continue;
        }

        visited[node.node] = node.weight;

        for (const next of graph[node.node]) {
            const nextWeight = node.weight + next.weight;
            if (visited[next.node] == -1 || visited[next.node] > nextWeight) {
                pq.enqueue(new Node(next.node, nextWeight));
            }
        }
    }

    return getMaxWeightIfPossible(visited);
};

function getMaxWeightIfPossible(visited) {
    let ans = visited[1];
    for (let node = 1; node < visited.length; node++) {
        if (visited[node] == -1) {
            return -1;
        }
        ans = Math.max(visited[node], ans);
    }
    return ans;
}

function constructGraph(n, times) {
    let graph = new Array(n + 1).fill().map(() => []);
    for (const time of times) {
        graph[time[0]].push(new Node(time[1], time[2]));
    }
    return graph;
};


function Node(node, weight) {
    this.node = node;
    this.weight = weight;
};

(function(){
    // Use Leetcode to execute.
    console.log("Case 1: ", networkDelayTime([[2,1,1],[2,3,1],[3,4,1]], 4, 2)); // => 2
    console.log("Case 2: ", networkDelayTime([[1,2,1]], 2, 1)); // => 1
    console.log("Case3: ", networkDelayTime([[1,2,1]], 2, 2)); // => -1
})();