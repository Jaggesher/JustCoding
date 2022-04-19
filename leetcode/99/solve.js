function TreeNode(val, left, right) {
    this.val = (val === undefined ? 0 : val)
    this.left = (left === undefined ? null : left)
    this.right = (right === undefined ? null : right)
}

/**
 * @param {TreeNode} root
 * @return {void} Do not return anything, modify root in-place instead.
 */
var recoverTree = function (root) {
    let node1 = null, node2 = null, prev = null;
    const inOrderTraverse = (node) => {
        if (!node) return;
        inOrderTraverse(node.left);
        if (prev && prev.val > node.val) {
            node2 = node;
            node1 = node1 ? node1 : prev;
        }
        prev = node;
        inOrderTraverse(node.right);
    }

    inOrderTraverse(root)
    const tm = node1.val
    node1.val = node2.val
    node2.val = tm
};

(function () {
    console.log("No test case.")
})();