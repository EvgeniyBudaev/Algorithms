/* https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/

Учитывая заголовок связанного списка, удалите n-й узел из конца списка и верните его заголовок.

Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Input: head = [1], n = 1
Output: []

Input: head = [1,2], n = 1
Output: [1]
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
 * Time O(N) | Space O(1)
 * @param {ListNode} head
 * @param {number} n
 * @return {ListNode}
 */
var removeNthFromEnd = function(head, n) {
    const length = getNthFromEnd(head, n);/* Time O(N) */
    const isHead = length < 0;
    if (isHead) return head.next;
    const curr = moveNode(head, length);  /* Time O(N) */
    curr.next = curr.next.next;
    return head
};

const getNthFromEnd = (curr, n, length = 0) => {
    while (curr) {                       /* Time O(N) */
        curr = curr.next;
        length++;
    }
    return (length - n) - 1;
}

const moveNode = (curr, length) => {
    while (length) {                    /* Time O(N) */
        curr = curr.next;
        length--;
    }
    return curr;
}

const head = createLinkedList([1,2,3,4,5]);
console.log(removeNthFromEnd(head, 2)); // [1,2,3,5]