/**
 * @param {string} s
 * @return {number}
 */
var countSubstrings = function (s) {
    let ans = 0;
    for (let index = 0; index < s.length; index++) {
        ans += tryExtend(s, index, index);
        ans += tryExtend(s, index, index + 1);
    }
    return ans;
};

function tryExtend(s, st, end) {
    let cnt = 0
    while (st >= 0 && end < s.length && s[st--] == s[end++]) {
        cnt++;
    }
    return cnt
};

/**
 * Time: O(n*n)
 * Space: O(1)
 */
(function(){
    console.log("Case 1: ", countSubstrings("abc"));
    console.log("Case 2: ", countSubstrings("aaa"));
})();