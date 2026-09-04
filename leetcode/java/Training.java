
public class Training {
    public static void main(String[] args) {
        System.out.println(validPalindrome("aba"));   // true
        System.out.println(validPalindrome("abca"));  // true
        System.out.println(validPalindrome("abc"));   // false
    }

    private static boolean validPalindrome(String s) {
        int left = 0, right = s.length() - 1;

        while (left < right) {
            // Если символы не равны, проверяем, можно ли сделать строку палиндромом, удалив один символ
            if (s.charAt(left) != s.charAt(right)) {
                return isPalindrome(s, left + 1, right) || isPalindrome(s, left, right - 1);
            }
            left++;
            right--;
        }
        return true; // Если цикл завершился успешно, строка является палиндромом
    }

    private static boolean isPalindrome(String s, int left, int right) {
        while (left < right) {
            // Если символы не равны, строка не является палиндромом
            if (s.charAt(left) != s.charAt(right)) {
                return false;
            }
            left++;
            right--;
        }
        return true; // Если цикл завершился успешно, строка является палиндромом
    }
}
