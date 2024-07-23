/* https://leetcode.com/problems/swap-nodes-in-pairs/description/

Учитывая связанный список, поменяйте местами каждые два соседних узла и верните его голову. Вы должны решить проблему,
не изменяя значения в узлах списка (т. е. изменять можно только сами узлы).

Input: head = [1,2,3,4]
Output: [2,1,4,3]

Input: head = []
Output: []

Input: head = [1]
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
 * @param {ListNode} head
 * @return {ListNode}
 */
var swapPairs = function(head) {
    if (!head || !head.next) return head;

    let NewHead = null;
    let prev = null;
    let curr = head;

    while (curr && curr.next) {
        let nextPair = curr.next.next;
        let temp = curr.next;
        curr.next = nextPair;
        temp.next = curr;

        if (prev) prev.next = temp;
        else NewHead = temp;

        prev = curr;
        curr = nextPair;
    }

    return NewHead;
};

const head = createLinkedList([1,2,3,4]);
console.log(swapPairs(head)); // [2,1,4,3]