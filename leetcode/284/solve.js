/**
 * // This is the Iterator's API interface.
 * // You should not implement it, or speculate about its implementation.
 * function Iterator() {
 *    @ return {number}
 *    this.next = function() { // return the next number of the iterator
 *       ...
 *    }; 
 *
 *    @return {boolean}
 *    this.hasNext = function() { // return true if it still has numbers
 *       ...
 *    };
 * };
 */

/**
 * @param {Iterator} iterator
 */
var PeekingIterator = function (iterator) {
    this.peakItem = null;
    this.iterator = iterator;
};

/**
 * @return {number}
 */
PeekingIterator.prototype.peek = function () {
    if (!this.peakItem) this.peakItem = this.iterator.next();
    return this.peakItem;
};

/**
 * @return {number}
 */
PeekingIterator.prototype.next = function () {
    const tm = this.peakItem ?? this.iterator.next();
    this.peakItem = null
    return tm;
};

/**
 * @return {boolean}
 */
PeekingIterator.prototype.hasNext = function () {
    return this.iterator.hasNext() || Boolean(this.peakItem);
};

/** 
 * Your PeekingIterator object will be instantiated and called as such:
 * var obj = new PeekingIterator(arr)
 * var param_1 = obj.peek()
 * var param_2 = obj.next()
 * var param_3 = obj.hasNext()
 */