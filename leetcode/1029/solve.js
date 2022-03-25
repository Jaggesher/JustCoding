/**
 * @param {number[][]} costs
 * @return {number}
 */
var twoCitySchedCost = function (costs) {
    costs.sort((a, b) => Math.abs(b[0] - b[1]) - Math.abs(a[0] - a[1]))
    let a = 0, b = 0, n = costs.length / 2, ans = 0
    costs.forEach((cost) => {
        if (((cost[0] <= cost[1]) && a < n) || b >= n) {
            a++
            ans += cost[0]
        } else {
            b++
            ans += cost[1]
        }
    })
    return ans
};

(function () {
    console.log("Case 1:", twoCitySchedCost([[10, 20], [30, 200], [400, 50], [30, 20]]))
    console.log("Case 2:", twoCitySchedCost([[259, 770], [448, 54], [926, 667], [184, 139], [840, 118], [577, 469]]))
    console.log("Case 3:", twoCitySchedCost([[515, 563], [451, 713], [537, 709], [343, 819], [855, 779], [457, 60], [650, 359], [631, 42]]))
})()