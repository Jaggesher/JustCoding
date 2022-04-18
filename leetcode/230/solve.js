function TreeNode(val, left, right) {
    this.val = (val === undefined ? 0 : val)
    this.left = (left === undefined ? null : left)
    this.right = (right === undefined ? null : right)
}

/**
 * @param {TreeNode} root
 * @param {number} k
 * @return {number}
 */
var kthSmallest = function (root, k) {
    let ans = 0
    let traverse = (node) => {
        if (!node || !k) return;
        traverse(node.left)
        k--
        if (k == 0) {
            ans = node.val
            return;
        }
        traverse(node.right)
    }
    traverse(root)
    return ans
};

(function(){
    console.log("No Test Case")
})()