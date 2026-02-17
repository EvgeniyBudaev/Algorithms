package greedy.Maximum69Number;

/* 1323. Maximum 69 Number

Вам дано целое положительное число, состоящее только из цифр 6 и 9.
Верните максимальное число, которое вы можете получить, изменив не более одной цифры
(6 становится 9, а 9 становится 6).

Input: num = 9669
Output: 9969
Explanation:
Changing the first digit results in 6669.
Changing the second digit results in 9969.
Changing the third digit results in 9699.
Changing the fourth digit results in 9666.
The maximum number is 9969.

Input: num = 9996
Output: 9999
Explanation: Changing the last digit 6 to 9 results in the maximum number.

Input: num = 9999
Output: 9999
Explanation: It is better not to apply any change.
*/

public class Maximum69Number {
    public static void main(String[] args) {
        System.out.println(maximum69Number(9996)); // 9999
    }

    // maximum69Number возвращает максимальное число, полученное из числа с заменой одной цифры 6 на 9
    // time: O(n), space: O(n), где n - длина строки
    private static int maximum69Number(int num) {
        String numStr = Integer.toString(num); // Преобразуем число в строку
        
        // Ищем первую '6' и заменяем её на '9'
        StringBuilder sb = new StringBuilder(numStr);
        for (int i = 0; i < sb.length(); i++) {
            if (sb.charAt(i) == '6') {
                sb.setCharAt(i, '9');
                break; // Заменяем только первую 6
            }
        }
        
        return Integer.parseInt(sb.toString()); // Преобразуем обратно в число
    }

    // более лаконично
    // private static int maximum69Number(int num) {
    //     return Integer.parseInt(Integer.toString(num).replaceFirst("6", "9"));
    // }
}
