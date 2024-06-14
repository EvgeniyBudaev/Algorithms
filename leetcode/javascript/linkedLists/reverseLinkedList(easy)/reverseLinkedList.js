/*
Учитывая заголовок односвязного списка, переверните список и верните перевернутый список.
 */

/*
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
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
 * @return {ListNode}
 */
var reverseList = function(head) {
    let prev = null;

    while (head) {
        const nextTemp = head.next;
        head.next = prev;
        prev = head;
        head = nextTemp;
    }

    return prev;
};

const head = createLinkedList([1,2,3,4,5]);
console.log(reverseList(head)); // [5,4,3,2,1]