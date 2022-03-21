/**
 * @param {string} s
 * @return {number[]}
 */
var partitionLabels = function (s) {
    let lastSeen = new Array(27)
    let ans = new Array()

    for (let i = 0; i < s.length; i++) {
        lastSeen[s.charCodeAt(i) - 97] = i
    }

    let st = 0, seen = lastSeen[s.charCodeAt(0) - 97]

    for (let i = 0; i < s.length; i++) {
        const charCode = s.charCodeAt(i) - 97
        if (seen == i) {
            ans.push(i - st + 1)
            st = i + 1
            seen = st < s.length ? lastSeen[s.charCodeAt(st) - 97] : -1
        } else if (seen < lastSeen[charCode]) {
            seen = lastSeen[charCode]
        }
    }
    return ans
};

(function () {
    console.log("Case 1: ", partitionLabels("ababcbacadefegdehijhklij"))
    console.log("Case 2: ", partitionLabels("eccbbbbdec"))
})()