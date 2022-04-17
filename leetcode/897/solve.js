

/**
 * Tree Node
 * @param {node Val} val 
 * @param {left ref} left 
 * @param {right ref} right 
 */
function TreeNode(val, left, right) {
    this.val = (val === undefined ? 0 : val)
    this.left = (left === undefined ? null : left)
    this.right = (right === undefined ? null : right)
};

/**
 * @param {TreeNode} root
 * @return {TreeNode}
 */
var increasingBST = function (root) {
    let dummy = new TreeNode();
    let current = dummy

    let traverse = (node) => {
        if (node == null) return;

        traverse(node.left)

        node.left = null
        current.right = node
        current = current.right

        traverse(node.right)
    }
    traverse(root)

    return dummy.right
};

(function(){
    console.log("No test case")
})()