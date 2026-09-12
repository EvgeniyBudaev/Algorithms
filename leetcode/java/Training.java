
public class Training {
    public static void main(String[] args) {
        System.out.println(isPalindrome(121)); // true
        System.out.println(isPalindrome(-121)); // false
    }

    private static boolean isPalindrome(int x) {
        // Отрицательные числа не могут быть палиндромами
        if (x < 0) {
            return false;
        }

        // Числа, оканчивающиеся на 0 (кроме 0) не могут быть палиндромами
        if (x % 10 == 0 && x != 0) {
            return false;
        }

        int reversed = 0;

        // Переворачиваем половину числа
        while (x > reversed) {
            reversed = reversed * 10 + x % 10;
            x /= 10;
        }

        // Для четного количества цифр: x == reversed
        // Для нечетного количества цифр: x == reversed / 10
        return x == reversed || x == reversed / 10;
    }
}
