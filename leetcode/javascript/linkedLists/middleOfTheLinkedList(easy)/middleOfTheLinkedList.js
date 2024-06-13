/*
Учитывая заголовок односвязного списка, верните средний узел связанного списка.
Если есть два средних узла, верните второй средний узел.
 */

/*
Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3. // Объяснение: Средний узел списка — это узел 3.
 */

/*
Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
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
var middleNode = function(head) {
    if (!head || !head.next) {
        return head;
    }

    let slow = head;
    let fast = head;

    while (fast && fast.next){
        fast = fast.next.next;
        slow = slow.next;
    }
    return slow;
};

const head = createLinkedList([1, 2, 3, 4, 5]);
console.log(middleNode(head)); // [3,4,5]
