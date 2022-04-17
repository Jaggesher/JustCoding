/**
 * @param {string} s
 * @return {number}
 */
var firstUniqChar = function (s) {
    let flag = new Map()
    for (let index = 0; index < s.length; index++) {
        flag.set(s[index], flag.has(s[index]) ? -1 : index)
    }

    let ans = -1
    for (let [, val] of flag) {
        if (val >= 0 && (ans == -1 || ans > val)) {
            ans = val
        }
    }

    return ans
};
(function () {
    console.log("Case 1: ", firstUniqChar("leetcode"))
    console.log("Case 2: ", firstUniqChar("loveleetcode"))
    console.log("Case 3: ", firstUniqChar("aabb"))
})()