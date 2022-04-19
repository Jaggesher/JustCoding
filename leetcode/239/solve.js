/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */

/***
 * Time: O(n) + O(k) => O(n)
 * Space: O(k)
 */
var maxSlidingWindow = function (nums, k) {
    let ans = [], monoQueue = [];

    for (let index = 0; index < nums.length; index++) {
        while (monoQueue.length > 0 && nums[monoQueue[monoQueue.length - 1]] <= nums[index]) monoQueue.pop();

        monoQueue.push(index);
        if (monoQueue.length && monoQueue[0] <= (index - k)) monoQueue.shift();

        if (index + 1 >= k) {
            ans.push(nums[monoQueue[0]])
        }
    }

    return ans;
};

(function () {
    console.log("Case 1: ", maxSlidingWindow([1, 3, -1, -3, 5, 3, 6, 7], 3));
    console.log("Case 2: ", maxSlidingWindow([1], 1));
})();