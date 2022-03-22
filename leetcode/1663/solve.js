/**
 * @param {number} n
 * @param {number} k
 * @return {string}
 */
var getSmallestString = function (n, k) {
    let ans = new Array(n)
    for (let index = 0; index < n; index++) {
        let mn = k - (26 * (n - index - 1))
        if (mn <= 0) {
            ans[index] = 'a'
            k--
        } else {
            ans[index] = String.fromCharCode(mn + 96)
            k -= mn
        }
    }
    return ans.join("")
};

(function () {
    console.log("Case 1: ", getSmallestString(3, 27))
    console.log("Case 2: ", getSmallestString(5, 73))
})()