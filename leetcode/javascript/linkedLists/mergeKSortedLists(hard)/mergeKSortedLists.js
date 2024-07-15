/* https://leetcode.com/problems/merge-k-sorted-lists/description/

Вам дан массив из k списков связанных списков, каждый связанный список отсортирован в порядке возрастания.
Объедините все связанные списки в один отсортированный связанный список и верните его.

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6

Input: lists = []
Output: []

Input: lists = [[]]
Output: []
*/

class ListNode {
    constructor(val, next) {
        this.val = (val!== undefined? val : 0);
        this.next = (next!== undefined? next : null);
    }
}
// Функция для создания односвязного списка
function createLinkedList(arr) {
    let head = null;
    let current = null;
    arr.forEach(val => {
        const newNode = new ListNode(val);
        if (!head) {
            head = newNode;
            current = newNode;
        } else {
            current.next = newNode;
            current = newNode;
        }
    });
    return head;
}

/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode[]} lists
 * @return {ListNode}
 */
var mergeKLists = function(lists) {
    const minHeap = getMinHeap(lists);
    return mergeLists(minHeap);
};

const getMinHeap = (lists) => {
    const heap = new MinPriorityQueue({ priority: ({ val }) => val });
    for (const node of lists) {
        if (!node) continue;
        heap.enqueue(node);
    }
    return heap;
}


const mergeLists = (minHeap) => {
    let sentinel = tail = new ListNode();
    while (!minHeap.isEmpty()) {
        const node = minHeap.dequeue().element;
        tail.next = node;
        tail = tail.next;
        if (!node.next) continue;
        minHeap.enqueue(node.next);
    }
    return sentinel.next;
}

const lists = createLinkedList([[1,4,5],[1,3,4],[2,6]]);
console.log(mergeKLists(lists));