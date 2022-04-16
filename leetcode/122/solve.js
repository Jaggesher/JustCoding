/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function (prices) {
    let ln = prices.length
    let maxProfit = 0

    for (let n = 0; n < ln; n++) {
        let valley = n
        while ((n + 1) < ln && prices[n] <= prices[n + 1]) {
            n++
        }
        maxProfit += prices[n] - prices[valley]
    }

    return maxProfit
};

(function () {
    console.log("Case 1:", maxProfit([7, 1, 5, 3, 6, 4]))
    console.log("Case 2:", maxProfit([1, 2, 3, 4, 5]))
    console.log("Case 3:", maxProfit([7, 6, 4, 3, 1]))
})()