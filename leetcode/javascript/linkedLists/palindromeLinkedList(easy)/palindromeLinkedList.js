/* https://leetcode.com/problems/palindrome-linked-list/description/

Учитывая заголовок односвязного списка, верните true, если это палиндром или ложь в противном случае.

Input: head = [1,2,2,1]
Output: true

Input: head = [1,2]
Output: false
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
 * @return {boolean}
 */
var isPalindrome = function(head) {
    const listVals = [];
    while (head) {
        listVals.push(head.val);
        head = head.next;
    }
    let left = 0, right = listVals.length - 1;
    while (left <= right) {
        if (listVals[left] !== listVals[right]) return false;
        left++;
        right--;
    }
    return true;
};

function isPalindromeLinkedList(head) {
    // Найдите середину связанного списка
    let slow = head;
    let fast = head;
    while (fast && fast.next) {
        slow = slow.next;
        fast = fast.next.next;
    }
    // Перевернуть вторую половину связанного списка
    let prev = null;
    let curr = slow;
    while (curr) {
        let next = curr.next;
        curr.next = prev;
        prev = curr;
        curr = next;
    }
    // Сравните первую половину связанного списка с перевернутой второй половиной
    let p1 = head;
    let p2 = prev;
    while (p2) {
        if (p1.val !== p2.val) return false;
        p1 = p1.next;
        p2 = p2.next;
    }
    return true;
}

const head = createLinkedList([1,2,2,1]);
console.log(isPalindrome(head)); // true
