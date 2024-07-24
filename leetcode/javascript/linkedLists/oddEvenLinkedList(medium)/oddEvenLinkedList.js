/* https://leetcode.com/problems/odd-even-linked-list/description/

Учитывая заголовок односвязного списка, сгруппируйте вместе все узлы с нечетными индексами, за которыми следуют узлы с
четными индексами, и верните переупорядоченный список.
Первый узел считается нечетным, второй — четным и так далее.
Обратите внимание, что относительный порядок внутри четных и нечетных групп должен оставаться таким, каким он был во
входных данных.
Вы должны решить задачу со сложностью дополнительного пространства O(1) и сложностью времени O(n).

Input: head = [1,2,3,4,5]
Output: [1,3,5,2,4]
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
var oddEvenList = function(head) {

};

const head = createLinkedList([1,2,3,4,5]);
console.log(oddEvenList(head)); // [1,3,5,2,4]