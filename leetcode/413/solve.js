/**
 * @param {number[]} nums
 * @return {number}
 */
 var numberOfArithmeticSlices = function(nums) {
    let ans = 0, dp = 0
    for (let i=2; i< nums.length; i++){
        if(nums[i-1]-nums[i] === nums[i-2]-nums[i-1]){
            dp++
            ans += dp
        }else{
            dp=0
        }
    }
    return ans
};

(function(){
    console.log("Case 1:", numberOfArithmeticSlices([1,2,3,4]))
    console.log("Case 2:", numberOfArithmeticSlices([1]))
})()