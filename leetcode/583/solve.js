/**
 * @param {string} word1
 * @param {string} word2
 * @return {number}
 */
var minDistance = function (word1, word2) {
    const n1 = word1.length, n2 = word2.length;
    return (n1 + n2) - (2 * LCS(word1, word2));
};

/**
 * Time: O(nm)
 * Space: O(nm)
 * @param {string} word1
 * @param {string} word2
 * @return {number}
 */
function LCS(word1, word2) {
    const n1 = word1.length, n2 = word2.length, dp = new Array(n1 + 1).fill().map(() => new Array(n2 + 1).fill(0));
    for (let i = 1; i <= n1; i++) {
        for (let j = 1; j <= n2; j++) {
            if (word1[i - 1] == word2[j - 1]) {
                dp[i][j] = dp[i - 1][j - 1] + 1;
            } else {
                dp[i][j] = Math.max(dp[i - 1][j], dp[i][j - 1])
            }
        }
    }
    return dp[n1][n2];
};

(function () {
    console.log("Case 1: ", minDistance("sea", "eat"))       // 2
    console.log("Case 2: ", minDistance("leetcode", "etco")) //4
})()