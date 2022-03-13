/**
 * @param {string} s
 * @return {boolean}
 */
 var isValid = function (s) {
    let stack  = new Array()
    const opening = ['(', '{', '[']
    for(let vl of s){
        if (opening.includes(vl)){
            stack.push(vl == '('? ')': vl == '{' ? '}' : ']')
        }else if(stack.length == 0 || stack.pop() != vl){
            return false
        }
    }
    return stack.length == 0 
};

(function () {
    console.log("Case 1: ", isValid("()"))
    console.log("Case 2: ", isValid("()[]{}"))
    console.log("Case 3: ", isValid("(]"))
})()