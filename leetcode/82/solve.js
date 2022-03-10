function ListNode(val, next) {
  this.val = val === undefined ? 0 : val;
  this.next = next === undefined ? null : next;
}

/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var deleteDuplicates = function (head) {
  let dummy = new ListNode(0, head);
  let temp = dummy;
  while (temp.next != null) {
    let tempNext = temp.next;
    if (tempNext.next != null && tempNext.val == tempNext.next.val) {
      while (tempNext.next != null && tempNext.val == tempNext.next.val) {
        tempNext = tempNext.next;
      }
      temp.next = tempNext.next;
    } else {
      temp = temp.next;
    }
  }
  return dummy.next;
};
