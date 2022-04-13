/**
 * @param {number} n
 * @return {number[][]}
 */
var generateMatrix = function (n) {
    let ans = new Array(n).fill().map(() => new Array(n))

    let row = 0, col = n - 1, cnt = 1
    while (row <= col) {

        // Go Right
        for (let index = row; index <= col; index++) {
            ans[row][index] = cnt++
        }

        // Go Down
        for (let index = row + 1; index <= col; index++) {
            ans[index][col] = cnt++
        }

        // Go Left
        for (let index = col - 1; index >= row; index--) {
            ans[col][index] = cnt++
        }

        // Go top
        for (let index = col - 1; index > row; index--) {
            ans[index][row] = cnt++
        }

        row++
        col--
    }

    return ans
};

(function(){
    console.log("Case 1: ", generateMatrix(3))
    console.log("Case 1: ", generateMatrix(1))
})()
