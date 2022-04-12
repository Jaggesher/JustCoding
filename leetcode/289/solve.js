/**
 * @param {number[][]} board
 * @return {void} Do not return anything, modify board in-place instead.
 */
var gameOfLife = function (board) {
    const dirs = [[-1, -1], [1, 1], [0, 1], [1, 0], [0, -1], [-1, 0], [1, -1], [-1, 1]]
    let m = board.length, n = board[m - 1].length
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            let neighbors = 0
            for (let dir of dirs) {
                let u = i + dir[0], v = j + dir[1]
                if (u < 0 || v < 0 || u >= m || v >= n) {
                    continue;
                }
                neighbors += (board[u][v] % 2)
            }

            if (board[i][j] == 1 && (neighbors < 2 || neighbors > 3)) {
                board[i][j] = 3
            } else if (board[i][j] == 0 && neighbors == 3) {
                board[i][j] = 2
            }
        }
    }

    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (board[i][j] == 2) {
                board[i][j] = 1
            } else if (board[i][j] == 3) {
                board[i][j] = 0
            }
        }
    }
};

(function(){
    let case1 = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
    gameOfLife(case1)
    console.log(case1)

    let case2 = [[1,1],[1,0]]
    gameOfLife(case2)
    console.log(case2)
})();