/**
 * @param {number[]} people
 * @param {number} limit
 * @return {number}
 */
 var numRescueBoats = function (people, limit) {
    people.sort((a,b)=> a-b)
    let ans = 0, st= 0, end = people.length -1
    while(st<=end){
        if(people[st]+people[end] <= limit){
            st++
        }
        ans++
        end--
    }
    return ans
};

(function () {
    console.log("Case 1:", numRescueBoats([1, 2], 3))
    console.log("Case 2:", numRescueBoats([3, 2, 2, 1], 3))
    console.log("Case 3:", numRescueBoats([3, 5, 3, 4], 5))
})()
