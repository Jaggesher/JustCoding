/**
 * @param {number[]} nums
 * @param {number} x
 * @return {number}
 */
var minOperations = function (nums, x) {
    let ans = -1, st = -1, end = nums.length - 1, sum = 0;
    while ((st + 1) <= end && (sum + nums[st + 1]) <= x) {
        st++
        sum += nums[st]
        if (sum == x) {
            ans = st + 1
            break
        }
    }

    while (st < end && end >= 0) {
        sum += nums[end]
        while (sum > x && st >= 0) {
            sum -= nums[st]
            st--
        }
        if (x == sum) {
            let tm = (st + 1) + (nums.length - end)
            if (ans == -1 || ans > tm) {
                ans = tm
            }
        }
        end--
    }

    return ans
};

(function () {
    console.log(minOperations([1, 1, 4, 2, 3], 5))
    console.log(minOperations([5, 6, 7, 8, 9], 4))
    console.log(minOperations([3, 2, 20, 1, 1, 3], 10))
    console.log(minOperations([3], 3))
})();