/**
 * @param {string} s
 * @param {number} k
 * @return {string}
 */
 var removeDuplicates = function(s, k) {
    let data = [], tracker = [];
    for(const char of s){
        
        let counter = 1;
        if(data.length && data[data.length - 1] == char){
            counter += tracker[tracker.length-1]
        }
        
        data.push(char);
        tracker.push(counter);
        
        if(counter == k){
            while(counter--){
                data.pop();
                tracker.pop();
            }
        }
    }
    
    return data.join("");
};

/**
 * Time: O(n) + O(n) => O(n);
 * Space: O(n)
 */

(function(){
    console.log("Case 1: ", removeDuplicates("abcd", 2)); // => "abcd"
    console.log("Case 2: ", removeDuplicates("deeedbbcccbdaa", 3)); // => "aa"
    console.log("Case 3: ", removeDuplicates("pbbcggttciiippooaais", 2)); // => "ps"
})();