/* https://leetcode.com/problems/copy-list-with-random-pointer/description/

Связный список длины n дан таким образом, что каждый узел содержит дополнительный случайный указатель, который может
указывать на любой узел в списке или на ноль.

Создайте глубокую копию списка. Глубокая копия должна состоять ровно из n совершенно новых узлов, где значение каждого
нового узла равно значению соответствующего исходного узла. И следующий, и случайный указатель новых узлов должны
указывать на новые узлы в скопированном списке, чтобы указатели в исходном списке и скопированном списке представляли
одно и то же состояние списка. Ни один из указателей в новом списке не должен указывать на узлы исходного списка.

Например, если в исходном списке есть два узла X и Y, где X.random --> Y, то для соответствующих двух узлов x и y в
скопированном списке x.random --> y.

Верните заголовок скопированного связанного списка.

Связанный список представляется на входе/выходе как список из n узлов.
Каждый узел представлен парой [val, random_index], где:

val: целое число, представляющее Node.val
Случайный_индекс: индекс узла (диапазон от 0 до n-1), на который указывает случайный указатель, или ноль, если он не
указывает ни на один узел.
Вашему коду будет присвоен только заголовок исходного связанного списка.

Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]
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
 * // Definition for a _Node.
 * function _Node(val, next, random) {
 *    this.val = val;
 *    this.next = next;
 *    this.random = random;
 * };
 */

/**
 * Time O(N) | Space O(1)
 * @param {_Node} head
 * @return {_Node}
 */
var copyRandomList = function(head) {
    if (!head) return null;
    cloneNode(head);         /* Time O(N) */
    connectRandomNode(head); /* Time O(N) */
    return connectNode(head);/* Time O(N) */
};

const cloneNode = (curr) => {
    while (curr) {          /* Time O(N) */
        const node = new Node(curr.val);

        node.next = curr.next;
        curr.next = node;
        curr = node.next;
    }
    return curr;
}

const connectRandomNode = (curr) => {
    while (curr) {         /* Time O(N) */
        curr.next.random = curr.random?.next || null;
        curr = curr.next.next;
    }
}

const connectNode = (head) => {
    let [ prev, curr, next ] = [ head, head.next, head.next ];

    while (prev) {        /* Time O(N) */
        prev.next = prev.next.next;
        curr.next = curr?.next?.next || null;
        prev = prev.next;
        curr = curr.next;
    }
    return next
}

const head = createLinkedList([[7,null],[13,0],[11,4],[10,2],[1,0]]);
console.log(copyRandomList(head)); // [[7,null],[13,0],[11,4],[10,2],[1,0]]