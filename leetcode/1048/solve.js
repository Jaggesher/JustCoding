/**
 * Time: O(n) + n Log(n)
 * Space: O(n)
 * @param {string[]} words
 * @return {number}
 */
var longestStrChain = function (words) {
    words.sort((a, b) => a.length - b.length) // n log n
    let ans = 0, dp = new Map();
    for (const word of words) { // O(n) * O(m) => As m is max 16 so => O(n)
        let mx = 0;
        for (let index = 0; index < word.length; index++) {
            const temp = word.substring(0, index) + word.substring(index + 1);
            mx = Math.max(mx, (dp.get(temp) ?? 0));
        }
        dp.set(word, ++mx)
        ans = Math.max(mx, ans);
    }
    return ans;
};

(function () {
    console.log("Case 1: ", longestStrChain(["a", "b", "ba", "bca", "bda", "bdca"]))
    console.log("Case 2: ", longestStrChain(["xbc", "pcxbcf", "xb", "cxbc", "pcxbc"]))
    console.log("Case 3: ", longestStrChain(["abcd", "dbqca"]))
})()