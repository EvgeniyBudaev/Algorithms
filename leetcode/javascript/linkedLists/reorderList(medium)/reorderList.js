/* https://leetcode.com/problems/reorder-list/description/

Вам дан заголовок односвязного списка. Список можно представить в виде:

L0 → L1 → … → Ln - 1 → Ln
Измените порядок списка, придав ему следующий вид:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
Вы не можете изменять значения в узлах списка. Изменять можно только сами узлы.

Input: head = [1,2,3,4]
Output: [1,4,2,3]

Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]
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
 * @return {void} Do not return anything, modify head in-place instead.
 */
var reorderList = function(head) {
    const mid = getMid(head);           /* Time O(N) */
    const reversedFromMid = reverse(mid);/* Time O(N) */
    reorder(head, reversedFromMid);      /* Time O(N) */
};

const getMid = (head) => {
    let [ slow, fast ] = [ head, head ];
    while (fast && fast.next) {         /* Time O(N) */
        slow = slow.next;
        fast = fast.next.next;
    }
    return slow;
};

const reverse = (head) => {
    let [ prev, curr, next ] = [ null, head, null ];
    while (curr) {                      /* Time O(N) */
        next = curr.next;
        curr.next = prev;

        prev = curr;
        curr = next;
    }
    return prev;
}

const reorder = (l1, l2) => {
    let [ first, next, second ] = [ l1, null, l2 ];
    while (second.next) {              /* Time O(N) */
        next = first.next;
        first.next = second;
        first = next;
        next = second.next;
        second.next = first;
        second = next;
    }
}

const head = createLinkedList([1,2,3,4]);
console.log(reorderList(head)); // [1,4,2,3]