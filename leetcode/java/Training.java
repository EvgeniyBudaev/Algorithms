
public class Training {
    public static void main(String[] args) {
        System.out.println(isPalindrome("A man, a plan, a canal: Panama")); // true
    }

    private static boolean isAlphaNumeric(char c) {
        return Character.isLetter(c) || Character.isDigit(c);
    }

    private static boolean isPalindrome(String s) {
        int left = 0, right = s.length() - 1;

        while (left < right) {
            // пропускаем символы, которые не являются буквенно-цифровыми
            boolean skipLeft = !isAlphaNumeric(s.charAt(left));
            if (skipLeft) {
                left++;
                continue;
            }

            // пропускаем символы, которые не являются буквенно-цифровыми
            boolean skipRight = !isAlphaNumeric(s.charAt(right));
            if (skipRight) {
                right--;
                continue;
            }

            // сравниваем символы, которые являются буквенно-цифровыми
            boolean endsEqual = Character.toLowerCase(s.charAt(left)) == Character.toLowerCase(s.charAt(right));
            if (!endsEqual) {
                return false;
            }

            left++;
            right--;
        }

        return true; // строка является палиндромом
    }
}
