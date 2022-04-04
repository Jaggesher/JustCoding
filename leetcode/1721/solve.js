
function ListNode(val, next) {
    this.val = (val === undefined ? 0 : val)
    this.next = (next === undefined ? null : next)
}

/**
 * @param {ListNode} head
 * @param {number} k
 * @return {ListNode}
 */
var swapNodes = function (head, k) {
    let first = head, last = head
    for (let index = 1, tm = head; tm != null; index++, tm = tm.next) {
        if (index == k) {
            first = tm
            last = head
        }
        if (index > k) {
            last = last.next
        }
    }

    let tm = first.val
    first.val = last.val
    last.val = tm

    return head
};

(function () {
    console.log("No test case")
})();