// https://leetcode.com/problems/kth-largest-element-in-a-stream/solutions/5135279/javascript-solution-minheap-priority-queue-time-o-nlogk-space-o-k/

/*
Разработайте класс для поиска k-го по величине элемента в потоке. Обратите внимание, что это k-й по величине элемент в
отсортированном порядке, а не k-й отдельный элемент.

Реализуйте класс KthLargest:

KthLargest(int k, int[] nums) Инициализирует объект целым числом k и потоком целых чисел nums.
int add(int val) Добавляет целое число val к потоку и возвращает элемент, представляющий k-й по величине
элемент в потоке.
*/

/*
Input
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
Output
[null, 4, 5, 5, 8, 8]

Explanation
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8
 */

/**
 * @param {number} k
 * @param {number[]} nums
 */
var KthLargest = function(k, nums) {
    this.heap = new MinHeap(k);
    nums.forEach((val) => this.heap.add(val));
}

/**
 * @param {number} val
 * @return {number}
 */
KthLargest.prototype.add = function(val){
    this.heap.add(val);
    return this.heap.peek();
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * var obj = new KthLargest(k, nums)
 * var param_1 = obj.add(val)
 */
class MinHeap {
    constructor(size) {
        this.heap = [0];
        this.size = size + 1;
    }

    add(val) {
        if(this.heap.length < this.size){
            this.heap.push(val)
            let child = this.heap.length - 1;
            let parent = Math.floor(child /2);

            while(parent >= 1 && this.heap[child] < this.heap[parent]){
                this.swap(parent, child);
                child = parent;
                parent = Math.floor(parent/2);
            }
        }
        else if(this.heap[1] < val){
            this.heap[1] = val;
            this.heapify(1);
        }
    }

    heapify(index) {
        const heapLength = this.heap.length;
        const leftChild = 2 * index ;
        const rightChild = 2 * index + 1;
        let minValue = index;

        if(leftChild < heapLength || rightChild < heapLength) {
            if(this.heap[minValue] > this.heap[leftChild]) {
                minValue = leftChild
            }
            if(this.heap[minValue] > this.heap[rightChild]) {
                minValue = rightChild
            }
        }

        if(minValue !== index) {
            this.swap(minValue, index);
            this.heapify(minValue)
        }
    }

    peek() {
        return this.heap[1]
    }

    swap(index1, index2){
        [this.heap[index1], this.heap[index2]] = [this.heap[index2], this.heap[index1]];
    }
}