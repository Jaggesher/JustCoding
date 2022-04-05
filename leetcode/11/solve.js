/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function (height) {
    let ans = 0, st = 0, end = height.length - 1
    while (st < end) {
        ans = Math.max(ans, (Math.min(height[st], height[end]) * (end - st)))
        height[st] < height[end] ? st++ : end--
    }
    return ans
};

(function () {
    console.log("Case 1: ", maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]))
    console.log("Case 2: ", maxArea([1, 1]))
})()