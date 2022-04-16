function TreeNode(val, left, right) {
    this.val = (val === undefined ? 0 : val)
    this.left = (left === undefined ? null : left)
    this.right = (right === undefined ? null : right)
}

/**
 * @param {TreeNode} root
 * @return {TreeNode}
 */
var convertBST = function (root) {
    let totalSum = 0
    let traverse = (node) => {
        if (!node) return
        traverse(node.right)
        totalSum += node.val
        node.val = totalSum
        traverse(node.left)
    }
    traverse(root)
    return root
};

(function(){
    console.log("No test case")
})()