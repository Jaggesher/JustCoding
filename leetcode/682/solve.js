/**
 * @param {string[]} ops
 * @return {number}
 */
var calPoints = function (ops) {
    let records = [], ans = 0
    for(let val of ops){
        switch(val){
            case "+":
                records.push(records[records.length-1]+records[records.length-2])
                break;
            case "D":
                records.push(records[records.length -1] * 2)
                break;
            case "C":
                records.pop()
                break;
            default:
                records.push(parseInt(val))
                break
        }
    }
    
    for(let val of records){
        ans += val
    }

    return ans
};

(function(){
    console.log("Case 1: ", calPoints(["5","2","C","D","+"]))
    console.log("Case 2: ", calPoints(["5","-2","4","C","D","9","+","+"]))
    console.log("Case 3: ", calPoints(["1"]))
})()