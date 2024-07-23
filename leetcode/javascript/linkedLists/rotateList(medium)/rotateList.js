/* https://leetcode.com/problems/rotate-list/description/

Учитывая заголовок связанного списка, поверните список вправо на k мест.

Input: head = [1,2,3,4,5], k = 2
Output: [4,5,1,2,3]
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
 * @param {ListNode} head
 * @param {number} k
 * @return {ListNode}
 */
var rotateRight = function(head, k) {
    let stack = [];
    let node = head;
    let i = 0;
    while(node) {
        stack.push(node);
        node = node.next;
        i++;
    }

    let n = k % i;

    while(n) {
        node = stack.pop();
        node.next = head;
        head = node;
        n--;
    }

    if (stack.length) {
        node = stack.pop();
        node.next = null;
    }

    return head;
};

const head = createLinkedList([1,2,3,4,5]);
console.log(rotateRight(head, 2)); // [4,5,1,2,3]
