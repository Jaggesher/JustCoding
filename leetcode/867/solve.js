/**
 * @param {number[][]} matrix
 * @return {number[][]}
 */
var transpose = function (matrix) {
    let m = matrix.length, n = matrix[0].length;
    const transpose = new Array(n).fill().map(() => new Array(m));
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            transpose[j][i] = matrix[i][j]
        }
    }
    return transpose;
};

/**
 * Time: O(n*m)
 * Space: O(n*m)
 */
(function () {
    console.log("Case 1: ", transpose([[1, 2, 3], [4, 5, 6], [7, 8, 9]]));
    console.log("Case 2: ", transpose([[1, 2, 3], [4, 5, 6]]));
})();