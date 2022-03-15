/**
 * @param {string} s
 * @return {string}
 */
 var minRemoveToMakeValid = function(s) {
     let data = s.split('')
    let stack = new Array()
    for(let i = 0; i< data.length; i++){
        if(data[i] == '('){
            stack.push(i)
        }else if(data[i] == ')'){
            stack.length > 0 ? stack.pop() : data[i] = '$'
        }
    }
    
    for(let index of stack){
        data[index] = '$'
    }

    return data.filter((x) => x != '$').join('')
};

(function(){
    console.log("Case 1: ", minRemoveToMakeValid("lee(t(c)o)de)"))
	console.log("Case 2: ", minRemoveToMakeValid("a)b(c)d"))
	console.log("Case 3: ", minRemoveToMakeValid("))(("))
})()