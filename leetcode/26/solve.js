/**
 * @param {number[]} nums
 * @return {number}
 */
 var removeDuplicates = function(nums) {
    let pos = 0
    for(let i = 1; i< nums.length; i++){
        if(nums[pos] != nums[i]){
            nums[++pos] = nums[i]
        }
    }
    return pos+1
};

(function(){
    console.log("Case 1: ", removeDuplicates([1, 1, 2]))
    console.log("Case 2: ", removeDuplicates([0, 0, 1, 1, 1, 2, 2, 3, 3, 4]))
})()