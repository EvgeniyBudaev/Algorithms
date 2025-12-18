package stacksAndQueues.EvaluateReversePolishNotation;

import java.util.ArrayDeque;
import java.util.Deque;

/* 150. Evaluate Reverse Polish Notation
https://leetcode.com/problems/evaluate-reverse-polish-notation/description/

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

public class EvaluateReversePolishNotation {
    public static void main(String[] args) {
        String[] tokens = {"2", "1", "+", "3", "*"};
        System.out.println(evalRPN(tokens)); // 9
    }

    // evalRPN оценивает выражения в обратной польской нотации
    // time: O(n), space: O(n), где n - количество токенов в выражении
    private static int evalRPN(String[] tokens) {
        Deque<Integer> stack = new ArrayDeque<>(); // Стек для хранения операндов

        for (String token : tokens) {
            switch (token) {
                case "+", "-", "*", "/" -> {
                    int right = stack.pop(); // Получаем и удаляем два последних числа из стека 
                    int left = stack.pop();
                    int result;
                    switch (token) {
                        case "+" -> result = left + right;
                        case "-" -> result = left - right;
                        case "*" -> result = left * right;
                        case "/" -> result = left / right;
                        default -> throw new IllegalStateException("Unexpected operator: " + token);
                    }
                    stack.push(result); // Выполняем операцию и добавляем результат обратно в стек
                }
                default -> {
                    // Это число — парсим и добавляем в стек
                    stack.push(Integer.valueOf(token));
                }
            }
        }

        return stack.pop(); // Остаток в стеке - результат вычисления выражения
    }
}
