/**
 * @param {string} s
 * @return {boolean}
 */
var validPalindrome = function (s) {
    let st = 0
    let end = s.length - 1
    while (st < end) {
        if (s.charAt(st) != s.charAt(end)) {
            return checkPalindrome(s, st + 1, end) || checkPalindrome(s, st, end - 1)
        }
        st++
        end--
    }
    return true
};

let checkPalindrome = function (s, st, end) {
    while (st < end) {
        if (s.charAt(st) != s.charAt(end)) {
            return false
        }
        st++
        end--
    }
    return true
};

(function () {
    console.log("Case 1: ", validPalindrome("aba"))
    console.log("Case 2: ", validPalindrome("abca"))
    console.log("Case 3: ", validPalindrome("abc"))
})()
