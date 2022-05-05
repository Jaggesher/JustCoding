/**
 * @param {number[]} gas
 * @param {number[]} cost
 * @return {number}
 */
var canCompleteCircuit = function (gas, cost) {
    let totalTank = 0, currentTank = 0, start = 0
    for (let i = 0; i < gas.length; i++) {
        const diff = gas[i] - cost[i]
        totalTank += diff;
        currentTank += diff;
        if (currentTank < 0) {
            start = i + 1;
            currentTank = 0
        }
    }
    return totalTank >= 0 ? start : -1
};

(function () {
    console.log("Case 1: ", canCompleteCircuit([1, 2, 3, 4, 5], [3, 4, 5, 1, 2]));
    console.log("Case 2: ", canCompleteCircuit([2, 3, 4], [3, 4, 3]));
})();