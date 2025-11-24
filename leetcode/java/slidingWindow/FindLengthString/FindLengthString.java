package slidingWindow.FindLengthString;

/*
Вам дана двоичная строка s (строка, содержащая только «0» и «1»). Вы можете выбрать до одного «0» и изменить его на «1».
Какова длина самой длинной подстроки, содержащей только «1»?

Например, если задано s = «1101100111», ответ будет равен 5. Если вы выполните переворот по индексу 2,
строка станет 1111100111.

Input: nums = "1101100111"
Output: 5
*/

public class FindLengthString {
    public static void main(String[] args) {
        System.out.println(findLengthString("1101100111")); // 5
    }

    // findLengthString вычисляет длину самой длинной подстроки, состоящей только из '1',
    // при условии, что не более одного '0' можно заменить на '1' (или не заменять ничего).
    // time: O(n), space: O(1)
    private static int findLengthString(String s) {
        int left = 0, zeroCount = 0, result = 0;

        for (int right = 0; right < s.length(); right++) {
            // Если слева стоит ноль, увеличиваем счетчик нулей
            if (s.charAt(right) == '0') {
                zeroCount++;
            }

            // Если количество нулей больше 1, сдвигаем left
            while (zeroCount > 1) {
                // Если слева стоит ноль, уменьшаем счетчик нулей
                if (s.charAt(left) == '0') {
                    zeroCount--;
                }
                left++; // Сдвигаем левый указатель
            }

            // Обновляем максимальную длину подстроки
            result = Math.max(result, right - left + 1);
        }

        return result; // Возвращаем максимальную длину
    }
}
