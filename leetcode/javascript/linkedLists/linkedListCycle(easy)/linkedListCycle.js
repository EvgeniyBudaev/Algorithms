/* https://leetcode.com/problems/linked-list-cycle/description/

Учитывая заголовок, заголовок связанного списка, определите, есть ли в связанном списке цикл.
В связанном списке существует цикл, если в списке есть некоторый узел, к которому можно снова добраться, непрерывно
следуя по следующему указателю. Внутри pos используется для обозначения индекса узла, к которому подключен следующий
указатель хвоста. Обратите внимание, что pos не передается в качестве параметра.
Верните true, если в связанном списке есть цикл. В противном случае верните false.

Input: head = [3,2,0,-4], pos = 1
Output: true
Пояснение: В связанном списке есть цикл, хвост которого соединяется с 1-м узлом (с индексом 0).
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
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * Time O(N) | Space O(1)
 * @param {ListNode} head
 * @return {boolean}
 */
var hasCycle = function(head) {
    let [ slow, fast ] = [ head, head];

    while (fast && fast.next) {/* Time O(N) */
        slow = slow.next;
        fast = fast.next.next;
        const hasCycle = slow === fast;
        if (hasCycle) return true;
    }

    return false;
};

const head = createLinkedList([3,2,0,-4]);
console.log(hasCycle(head, 1));