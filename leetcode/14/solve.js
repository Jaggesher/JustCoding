/**
 * Time: O(n*m)
 * Space: O(1)
 * @param {string[]} strs
 * @return {string}
 */
 var longestCommonPrefix = function(strs) {
    let ans = "", ln = 202
    for(let vl of strs) ln = Math.min(ln, vl.length);
    
    for(let index= 0; index < ln; index++){
        for(let i = 1; i < strs.length; i++){
            if(strs[i-1][index] != strs[i][index]){
                return ans
            }
        }
        ans += strs[0][index]
    }
    
    return ans
};

(function(){
    console.log("Case 1: ", longestCommonPrefix(["flower","flow","flight"]))
    console.log("Case 2: ", longestCommonPrefix(["dog","racecar","car"]))
})();