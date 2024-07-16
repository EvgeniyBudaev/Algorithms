/* https://leetcode.com/problems/remove-linked-list-elements/description/

Учитывая заголовок связанного списка и целое число val, удалите все узлы связанного списка, у которых есть
Node.val == val, и верните новый заголовок.

Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]

Input: head = [], val = 1
Output: []

Input: head = [7,7,7,7], val = 7
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
 * @param {ListNode} head
 * @param {number} val
 * @return {ListNode}
 */
var removeElements = function(head, val) {
    if (!head) return head;

    while (head) {
        if (head.val === val) head = head.next;
        else break;
    }

    let curr = head;
    while (curr && curr.next) {
        if (curr.next.val === val) curr.next = curr.next.next;
        else curr = curr.next;
    }

    return head;
};

const head = createLinkedList([1,2,6,3,4,5,6], 6);
console.log(removeElements(head, 6)); // [1,2,3,4,5]