
function Node(val, next, random) {
    this.val = val;
    this.next = next;
    this.random = random;
};


/**
 * @param {Node} head
 * @return {Node}
 */
var copyRandomList = function (head) {
    let ans = new Node(0, null, null)
    let tracker = new Array()
    
    for(let index = 0, temp= ans, original = head; original != null; index++){
        let tm = new Node(original.val, null, null)
        tracker.push(tm)
        temp.next = tm
        temp = tm
        original.val = index
        original = original.next
    }
    for(let cp= ans.next, original= head; original != null && cp != null; cp= cp.next, original = original.next){
        if(original.random != null){
            cp.random = tracker[original.random.val]
        }
    }
    return ans.next
};

(function(){
    console.log("No test case")
})()