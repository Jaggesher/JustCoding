function ListNode(val, next) {
  this.val = val === undefined ? 0 : val;
  this.next = next === undefined ? null : next;
}
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function (l1, l2) {
    let head = new ListNode(0, null), extra = 0
    let temp = head
    while(l1 != null || l2 != null || extra > 0){
        if(l1 != null){
            extra += l1.val
            l1 = l1.next
        }
        
        if(l2 != null){
            extra += l2.val
            l2 = l2.next
        }
        
        temp.next = new ListNode(extra % 10, null)
        temp = temp.next
        extra = Math.floor(extra/10)
    }
    return head.next
};
