/* https://leetcode.com/problems/sort-list/description/
solution https://leetcode.com/problems/sort-list/solutions/5113112/extreme-in-depth-stepwise/

Учитывая заголовок связанного списка, верните список после его сортировки по возрастанию.

Input: head = [4,2,1,3]
Output: [1,2,3,4]
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
var sortList = function(head) {
    if (!head || !head.next) return head;

    const merge = (left, right) => {
        const dummy = new ListNode(-1);
        let curr = dummy;

        while (left && right) {
            if (left.val <= right.val) {
                curr.next = left;
                left = left.next;
            } else {
                curr.next = right;
                right = right.next;
            }
            curr = curr.next;
        }

        curr.next = left || right;

        return dummy.next;
    };

    const findMiddle = (node) => {
        let slow = node;
        let fast = node.next;

        while (fast && fast.next) {
            slow = slow.next;
            fast = fast.next.next;
        }

        return slow;
    };

    const mid = findMiddle(head);
    const rightHead = mid.next;
    mid.next = null;

    const sortedLeft = sortList(head);
    const sortedRight = sortList(rightHead);

    return merge(sortedLeft, sortedRight);
};
const head = createLinkedList([4,2,1,3]);
console.log(sortList(head)); // [1,2,3,4]