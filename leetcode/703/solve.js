/**
 * @param {number} k
 * @param {number[]} nums
 */
var KthLargest = function (k, nums) {
    this.k = k
    // https://github.com/datastructures-js/priority-queue
    this.pq = new MinPriorityQueue()
    for (let vl of nums) {
        this.pq.enqueue(vl)
    }
    while (this.pq.size() > this.k) {
        this.pq.dequeue()
    }
};

/** 
 * @param {number} val
 * @return {number}
 */
KthLargest.prototype.add = function (val) {
    this.pq.enqueue(val)
    if (this.pq.size() > this.k) {
        this.pq.dequeue()
    }
    return this.pq.front().element
};

/** 
 * Your KthLargest object will be instantiated and called as such:
 * var obj = new KthLargest(k, nums)
 * var param_1 = obj.add(val)
 */