/**
 * @param {number[]} nums
 * @param {number} k
 * @param {number} t
 * @return {boolean}
 */
var containsNearbyAlmostDuplicate = function (nums, k, t) {
    for (let i = 0; i < nums.length; i++) {
        let st = i - k
        st = st < 0 ? 0 : st
        while (st < i) {
            if (Math.abs(nums[st] - nums[i]) <= t) {
                return true
            }
            st++
        }
    }

    return false
};

(function () {
    console.log("Case 1: ", containsNearbyAlmostDuplicate([1, 2, 3, 1], 3, 0))
    console.log("Case 2: ", containsNearbyAlmostDuplicate([1, 0, 1, 1], 1, 2))
    console.log("Case 3: ", containsNearbyAlmostDuplicate([1, 5, 9, 1, 5, 9], 2, 3))
})();