/**
 * @param {number[]} stones
 * @return {number}
 */
var lastStoneWeight = function (stones) {
    while (stones.length > 1) {
        stones.sort((a, b) => b - a)
        stones[1] = stones[0] - stones[1]
        stones.shift()
    }
    return stones[0]
};

(function () {
    console.log("Case 1: ", lastStoneWeight([2, 7, 4, 1, 8, 1]))
    console.log("Case 2: ", lastStoneWeight([1]))
})();