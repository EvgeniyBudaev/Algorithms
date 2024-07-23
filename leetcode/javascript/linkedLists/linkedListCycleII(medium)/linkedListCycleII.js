/* https://leetcode.com/problems/linked-list-cycle-ii/description/

Учитывая заголовок связанного списка, верните узел, с которого начинается цикл. Если цикла нет, верните ноль.
В связанном списке существует цикл, если в списке есть некоторый узел, к которому можно снова добраться, непрерывно
следуя по следующему указателю. Внутри pos используется для обозначения индекса узла, к которому подключен следующий
указатель хвоста (с индексом 0). Это -1, если цикла нет. Обратите внимание, что pos не передается в качестве параметра.
Не изменяйте связанный список.

Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.
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
 * @param {ListNode} head
 * @return {ListNode}
 */
var detectCycle = function(head) {
    if (head === null || head.next === null) return null;
    let slow = head;
    let fast = head;
    while (fast !== null && fast.next !== null) {
        slow = slow.next;
        fast = fast.next.next;
        if (slow === fast) {
            slow = head;
            while (slow !== fast){
                slow = slow.next;
                fast = fast.next;
            }
            return slow;
        }
    }
    return null;
};

const head = createLinkedList([3,2,0,-4]);
console.log(detectCycle(head));