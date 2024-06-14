/*
Учитывая заголовок отсортированного связанного списка, удалите все дубликаты, чтобы каждый элемент появлялся только один
раз. Верните связанный список также отсортированным.
 */

/*
Input: head = [1,1,2]
Output: [1,2]
 */

/*
Input: head = [1,1,2,3,3]
Output: [1,2,3]
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
var deleteDuplicates = function(head) {
    let result = head;

    while (head && head.next) {
        if (head.val === head.next.val) {
            head.next = head.next.next;
        } else {
            head = head.next;
        }
    }

    return result;
}

const head = createLinkedList([1,1,2,3,3]);
console.log(deleteDuplicates(head)); // [1,2,3]