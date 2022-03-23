/**
 * @param {number} startValue
 * @param {number} target
 * @return {number}
 */
var brokenCalc = function (startValue, target) {
    let ans = 0
    while (target > startValue) {
        ans++
        if (target & 1) {
            target++
        } else {
            target = target >> 1
        }
    }

    return ans + startValue - target
};

(function () {
    console.log("Case 1: ", brokenCalc(2, 3))
    console.log("Case 2: ", brokenCalc(5, 8))
    console.log("Case 3: ", brokenCalc(3, 10))
})()