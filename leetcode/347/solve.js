/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
var topKFrequent = function (nums, k) {
    let tracker = {}
    for (let vl of nums) {
        tracker[vl] = (tracker[vl] || 0) + 1
    }
    let keys = Object.keys(tracker)
    keys.sort((i, j) => tracker[j] - tracker[i])
    return keys.slice(0, k)
};

(function () {
    console.log("Case 1: ", topKFrequent([1, 1, 1, 2, 2, 3], 2))
    console.log("Case 2: ", topKFrequent([1], 1))
})()