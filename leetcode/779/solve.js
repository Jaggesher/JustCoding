/**
 * @param {number} poured
 * @param {number} query_row
 * @param {number} query_glass
 * @return {number}
 */
var champagneTower = function (poured, query_row, query_glass) {
    let currentRow = [poured]
    for(let i = 0; i < query_row; i++ ){
        let nextRow = new Array(i+2).fill(0)
        for(let j = 0; j <= i; j++){
            let tm = (currentRow[j]-1) / 2
            if(tm > 0){
                nextRow[j] += tm
                nextRow[j+1] += tm
            }
        }
        currentRow = nextRow
    }
    return Math.min(currentRow[query_glass], 1)
};

(function () {
  console.log("Case 1:", champagneTower(1, 1, 1));
  console.log("Case 2:", champagneTower(2, 1, 1));
  console.log("Case 3:", champagneTower(100000009, 33, 17));
})();
