function TreeNode(val, left, right) {
    this.val = (val === undefined ? 0 : val)
    this.left = (left === undefined ? null : left)
    this.right = (right === undefined ? null : right)
}

class BSTIterator {
    nodes = []

    constructor(root) {
        this.nodes = [root];
        this.insertAll();
    }

    next() {
        const tm = this.nodes.pop();
        if (tm.right) {
            this.nodes.push(tm.right);
            this.insertAll();
        }
        return tm.val
    }

    hasNext() {
        return this.nodes.length > 0
    }

    insertAll() {
        while (this.nodes[this.nodes.length - 1].left) {
            this.nodes.push(this.nodes[this.nodes.length - 1].left)
        }
    }
};


(function () {
    var obj = new BSTIterator(new TreeNode(1, null, null))
    console.log(obj.next())
    console.log(obj.hasNext())
})();

