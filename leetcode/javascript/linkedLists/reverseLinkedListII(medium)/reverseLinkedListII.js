/*
Учитывая заголовок односвязного списка и два целых числа слева и справа, где слева <= справа, поменяйте местами узлы
списка с позиции слева на позицию справа и верните перевернутый список.
 */

/*
Input: head = [1,2,3,4,5], left = 2, right = 4
Output: [1,4,3,2,5]
 */

/*
Input: head = [5], left = 1, right = 1
Output: [5]
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
 * @param {number} left
 * @param {number} right
 * @return {ListNode}
 */
var reverseBetween = function(head, left, right) {
    if (!head || left === right) return head;

    const dummy = new ListNode(0);
    dummy.next = head;
    let prev = dummy;

    for (let i = 0; i < left - 1; ++i) {
        prev = prev.next;
    }

    let current = prev.next;
    let nextNode = null;

    for (let i = 0; i < right - left + 1; ++i) {
        const temp = current.next;
        current.next = nextNode;
        nextNode = current;
        current = temp;
    }

    prev.next.next = current;
    prev.next = nextNode;

    return dummy.next;
};

const head = createLinkedList([1,2,3,4,5]);
console.log(reverseBetween(head, 2, 4)); // [1,4,3,2,5]