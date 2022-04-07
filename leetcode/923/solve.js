/**
 * @param {number[]} arr
 * @param {number} target
 * @return {number}
 */
var threeSumMulti = function (arr, target) {
    const MOD = 1_000_000_007
    let ans = 0
    let tracker = new Array(101).fill(0)
    for (let vl of arr) {
        tracker[vl]++
    }
    
    for (let i = 0; i < arr.length; i++) {
        tracker[arr[i]]--
        let temp = new Array(101).fill(0)
        for (let j = i + 1; j < arr.length; j++) {
            temp[arr[j]]++
            let tm = target - (arr[i] + arr[j])
            if (tm >= 0 && tm < 101) {
                ans += tracker[tm] - temp[tm]
                ans = ans > MOD ? ans % MOD : ans
            }
        }
    }

    return ans
};

(function () {
    console.log("Case 1: ", threeSumMulti([1, 1, 2, 2, 3, 3, 4, 4, 5, 5], 8))
    console.log("Case 2: ", threeSumMulti([1, 1, 2, 2, 2, 2], 5))
})();