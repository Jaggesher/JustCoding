/**
 * @param {number[][]} grid
 * @param {number} k
 * @return {number[][]}
 */
 var shiftGrid = function(grid, k) {
    let previous = null, m = grid.length, n = grid[0].length
    while(k--){
        for(let i = 0; i < m; i++){
            previous = grid[i][n-1]
            for(let j = 0; j < n; j++){
                [grid[i][j], previous] = [previous, grid[i][j]]
            }
        }
        
        previous = grid[m-1][0]
        for(let i = 0; i < m; i++){
            [grid[i][0], previous]= [previous, grid[i][0]]
        }
    }
    return grid
};

(function(){
    console.log("Case 1: ", shiftGrid([[1,2,3],[4,5,6],[7,8,9]], 1))
    console.log("Case 2: ", shiftGrid([[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]], 4))
    console.log("Case 3: ", shiftGrid([[1,2,3],[4,5,6],[7,8,9]], 9))
})();