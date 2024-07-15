/* https://leetcode.com/problems/reverse-nodes-in-k-group/description/

Учитывая заголовок связанного списка, поменяйте местами узлы списка k за раз и верните измененный список.
k — целое положительное число, меньшее или равное длине связанного списка. Если количество узлов не кратно k, то
пропущенные узлы в конечном итоге должны остаться такими, какие они есть.
Вы не можете изменять значения в узлах списка, можно изменять только сами узлы.

Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]
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
var reverseKGroup = function(head, k) {
    const sentinel = tail = new ListNode(0, head);
    while (true) {
        let [ start, last ]= moveNode(tail, k);
        if (!last) break;
        reverse([ start, tail.next, start ])
        const next = tail.next;
        tail.next = last;
        tail = next;
    }
    return sentinel.next;
};

const moveNode = (curr, k) => {
    const canMove = () => k && curr;
    while (canMove()) {
        curr = curr.next;
        k--;
    }
    return [ (curr?.next || null), curr ];
}

const reverse = ([ prev, curr, start ]) => {
    const isSame = () => curr === start;
    while (!isSame()) {
        const next = curr.next;
        curr.next = prev;
        prev = curr;
        curr = next;
    }
}

const head = createLinkedList([1,2,3,4,5]);
console.log(reverseKGroup(head, 2)); // [2,1,4,3,5]
