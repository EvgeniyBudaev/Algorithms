/* https://leetcode.com/problems/find-k-pairs-with-smallest-sums/description/

Вам даны два массива целых чисел nums1 и nums2, отсортированные в порядке неубывания, и целое число k.
Определите пару (u, v), состоящую из одного элемента первого массива и одного элемента второго массива.
Верните k пар (u1, v1), (u2, v2), ..., (uk, vk) с наименьшими суммами.

Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
Output: [[1,2],[1,4],[1,6]]
Объяснение: Из последовательности возвращаются первые 3 пары: [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7 ,6],[11,4],[11,6]

Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
Output: [[1,1],[1,1]]
Объяснение: Из последовательности возвращаются первые 2 пары: [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1 ,3],[1,3],[2,3]
*/

/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @param {number} k
 * @return {number[][]}
 */
var kSmallestPairs = function(nums1, nums2, k) {
    const heap = new MinHeap();
    const result = [];

    if (nums1.length === 0 || nums2.length === 0) return [];

    for (let i = 0; i < Math.min(k, nums1.length); i++) {
        heap.push([nums1[i] + nums2[0], i, 0]);
    }

    while (k > 0 && heap.size() > 0) {
        const [val, i, j] = heap.pop();
        result.push([nums1[i], nums2[j]]);
        k--;

        if (j + 1 < nums2.length) {
            heap.push([nums1[i] + nums2[j + 1], i, j + 1]);
        }
    }

    return result;
};

class MinHeap {
    constructor() {
        this.heap = [];
    }

    push(val) {
        this.heap.push(val);
        this._bubbleUp();
    }

    pop() {
        if (this.size() === 1) return this.heap.pop();
        const smallest = this.heap[0];
        this.heap[0] = this.heap.pop();
        this._sinkDown(0);
        return smallest;
    }

    size() {
        return this.heap.length;
    }

    _bubbleUp() {
        let index = this.heap.length - 1;
        while (index > 0) {
            const parentIndex = Math.floor((index - 1) / 2);
            if (this.heap[parentIndex][0] > this.heap[index][0]) {
                [this.heap[parentIndex], this.heap[index]] = [this.heap[index], this.heap[parentIndex]];
                index = parentIndex;
            } else {
                break;
            }
        }
    }

    _sinkDown(index) {
        const left = 2 * index + 1;
        const right = 2 * index + 2;
        let smallest = index;

        if (left < this.size() && this.heap[smallest][0] > this.heap[left][0]) {
            smallest = left;
        }

        if (right < this.size() && this.heap[smallest][0] > this.heap[right][0]) {
            smallest = right;
        }

        if (smallest !== index) {
            [this.heap[smallest], this.heap[index]] = [this.heap[index], this.heap[smallest]];
            this._sinkDown(smallest);
        }
    }
}

const nums1 = [1,7,11];
const nums2 = [2,4,6];
const k = 3;

console.log(kSmallestPairs(nums1, nums2, k));