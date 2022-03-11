function ListNode(val, next) {
    this.val = (val === undefined ? 0 : val)
    this.next = (next === undefined ? null : next)
}

/**
* @param {ListNode} head
* @param {number} k
* @return {ListNode}
*/
var rotateRight = function (head, k) {
    if (k == 0 || head == null) {
        return head
    }

    let temp = head, count = 1
    while (temp.next != null) {
        count++
        temp = temp.next
    }
    temp.next = head;
    temp = head;
    count = count - (k % count)
    while (count > 1) {
        temp = temp.next
        count--
    }
    head = temp.next
    temp.next = null
    return head
};

(function(){
    console.log("No test case")
})()