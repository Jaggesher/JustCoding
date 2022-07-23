/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} x
 * @return {ListNode}
 */
var partition = function (head, x) {
    let firstStart = new ListNode(), first = firstStart;
    let secStart = new ListNode(), sec = secStart;
    while (head) { // O(n)
        if (head.val < x) {
            first.next = head;
            first = first.next;
        } else {
            sec.next = head;
            sec = sec.next
        }
        head = head.next
    }

    first.next = secStart.next;
    sec.next = null;
    return firstStart.next;
};

/**
 * Time: O(n)
 * Space: O(1)
 */
(function(){
    console.log(`No test case`);
})();