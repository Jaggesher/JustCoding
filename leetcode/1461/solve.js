/**
 * @param {string} s
 * @param {number} k
 * @return {boolean}
 */
var hasAllCodes = function (s, k) {
    const tracker = new Set();
    for (let range = k; range <= s.length; range++) {
        tracker.add(s.substring(range - k, range));
    }
    return tracker.size == (2 ** k);
};

/**
 * Time: O(n)
 * Space: O(m) where m is the length of all possible code. 
 */
(function () {
    console.log("Case 1: ", hasAllCodes("00110110", 2)); // True
    console.log("Case 2: ", hasAllCodes("0110", 1)); // True
    console.log("Case 3: ", hasAllCodes("0110", 2)); // False
})()