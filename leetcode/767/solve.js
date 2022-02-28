/**
 * @param {string} s
 * @return {string}
 */
var reorganizeString = function (s) {
  let map = new Map();
  for (let char of s) {
    if (map.has(char)) {
      map.set(char, map.get(char) + 1);
    } else {
      map.set(char, 1);
    }
  }

  let ans = new Array(s.length);

  let lastSeen = "";

  for (let i = 0; i < ans.length; i++) {
    let mx = "";
    for (let [key, val] of map) {
      if (val > 0 && (mx == "" || map.get(mx) < val) && lastSeen != key) {
        mx = key;
      }
    }

    if (mx == "") {
      return "";
    }

    ans[i] = mx;
    lastSeen = mx;
    map.set(mx, map.get(mx) - 1);
  }

  return ans.join("");
};

(()=>{
    console.log("Case 1: ", reorganizeString("aab"))
    console.log("Case 2: ", reorganizeString("aaab"))
})()
