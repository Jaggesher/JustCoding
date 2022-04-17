/**
 * @param {character[][]} board
 * @return {boolean}
 */
var isValidSudoku = function (board) {
    let tracker = new Set()

    for (let i = 0; i < 9; i++) {
        const grid = (3 * Math.floor(i / 3))
        for (let j = 0; j < 9; j++) {
            if (board[i][j] == '.') continue;
            const rowId = `row-${i}-${board[i][j]}`;
            const colId = `col-${j}-${board[i][j]}`;
            const gridId = `grid-${grid + Math.floor(j / 3)}-${board[i][j]}`;

            if (tracker.has(rowId) || tracker.has(colId) || tracker.has(gridId)) {
                return false
            }
            tracker.add(rowId)
            tracker.add(colId)
            tracker.add(gridId)
        }
    }

    return true
};

(function () {
    const board1 = [["5", "3", ".", ".", "7", ".", ".", ".", "."]
        , ["6", ".", ".", "1", "9", "5", ".", ".", "."]
        , [".", "9", "8", ".", ".", ".", ".", "6", "."]
        , ["8", ".", ".", ".", "6", ".", ".", ".", "3"]
        , ["4", ".", ".", "8", ".", "3", ".", ".", "1"]
        , ["7", ".", ".", ".", "2", ".", ".", ".", "6"]
        , [".", "6", ".", ".", ".", ".", "2", "8", "."]
        , [".", ".", ".", "4", "1", "9", ".", ".", "5"]
        , [".", ".", ".", ".", "8", ".", ".", "7", "9"]]
    console.log("Case 1: ", isValidSudoku(board1))

    const board2 = [["8", "3", ".", ".", "7", ".", ".", ".", "."]
        , ["6", ".", ".", "1", "9", "5", ".", ".", "."]
        , [".", "9", "8", ".", ".", ".", ".", "6", "."]
        , ["8", ".", ".", ".", "6", ".", ".", ".", "3"]
        , ["4", ".", ".", "8", ".", "3", ".", ".", "1"]
        , ["7", ".", ".", ".", "2", ".", ".", ".", "6"]
        , [".", "6", ".", ".", ".", ".", "2", "8", "."]
        , [".", ".", ".", "4", "1", "9", ".", ".", "5"]
        , [".", ".", ".", ".", "8", ".", ".", "7", "9"]]
    console.log("Case 2: ", isValidSudoku(board2))
})()