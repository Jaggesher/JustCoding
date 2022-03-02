/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
 var isSubsequence = function(s, t) {
    let sIndex = 0, tIndex = 0
    while(sIndex < s.length && tIndex < t.length){
        if (s[sIndex] === t[tIndex]){
            sIndex++
        }
        tIndex++
    }
    
    return sIndex === s.length
};

(function(){
    console.log("Case 1: ", isSubsequence("abc", "ahbgdc"))
	console.log("Case 2: ", isSubsequence("axc", "ahbgdc"))
})()