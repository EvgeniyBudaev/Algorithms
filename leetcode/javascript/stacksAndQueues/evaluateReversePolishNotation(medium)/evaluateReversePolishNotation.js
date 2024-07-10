/* https://leetcode.com/problems/evaluate-reverse-polish-notation/description/

Вам дан массив строковых токенов, который представляет арифметическое выражение в обратной польской нотации.
Оцените выражение. Возвращает целое число, представляющее значение выражения.

Обратите внимание, что:

Допустимые операторы: «+», «-», «*» и «/».
Каждый операнд может быть целым числом или другим выражением.
Деление между двумя целыми числами всегда сокращается до нуля.
Никакого деления на ноль не будет.
Входные данные представляют собой допустимое арифметическое выражение в обратной польской записи.
Ответ и все промежуточные вычисления можно представить в виде 32-битного целого числа.

Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9

Input: tokens = ["4","13","5","/","+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6
*/

/**
 * Time O(N) | Space(N)
 * @param {string[]} tokens
 * @return {number}
 */
var evalRPN = function(tokens) {
    let stack = [];
    for (const char of tokens) {/* Time O(N) */
        const isOperation = char in OPERATORS;
        if (isOperation) {
            const value = performOperation(char, stack);
            stack.push(value);      /* Space O(N) */
            continue;
        }
        stack.push(Number(char));   /* Space O(N) */
    }
    return stack.pop();
};

const OPERATORS = {
    '+': (a, b) => a + b,
    '-': (a, b) => a - b,
    '*': (a, b) => a * b,
    '/': (a, b) => Math.trunc(a / b)
};

const performOperation = (char, stack) => {
    const [ rightNum, leftNum ] = [ stack.pop(), stack.pop() ];
    const operation = OPERATORS[char];
    return operation(leftNum, rightNum);
}

console.log(evalRPN(["2","1","+","3","*"])); // 9
