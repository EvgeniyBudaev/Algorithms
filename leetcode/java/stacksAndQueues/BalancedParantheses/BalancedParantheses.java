package stacksAndQueues.BalancedParantheses;

/* Balanced Parantheses

Дана строка A, состоящая только из '(' и ')'.
Вам нужно выяснить, сбалансированы ли скобки в A или нет, если сбалансированы, то вернуть 1, иначе вернуть 0.

Input: A = "(()())"
Output: 1

Input: A = "(()"
Output: 0
*/

public class BalancedParantheses {
    public static void main(String[] args) {
        String str = "(()())";
        System.out.println(solve(str)); // 1
    }

    // solve проверяет, сбалансированы ли скобки в строке.
    // time: O(n), space: O(1)
    private static int solve(String str) {
        int balance = 0; // Счетчик открывающих скобок

        for (char ch : str.toCharArray()) {
            if (ch == '(') {
                balance++;
            } else {
                balance--;
            }

            // Если баланс стал отрицательным - сразу возвращаем 0
            if (balance < 0) {
                return 0;
            }
        }

        // Возвращаем 1, если баланс = 0, иначе 0
        return balance == 0 ? 1 : 0; // Если баланс не равен нулю, значит скобки не сбалансированы
    }
}
