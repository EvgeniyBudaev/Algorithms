
public class Training {
    public static void main(String[] args) {
        System.out.println(validPalindrome("aba"));   // true
        System.out.println(validPalindrome("abca"));  // true
        System.out.println(validPalindrome("abc"));   // false
    }

    private static boolean isPalindrome(String s, int left, int right) {
        while (left < right) {
            if (s.charAt(left) != s.charAt(right)) {
                return false;
            }

            left++;
            right--;
        }

        return true;
    }

    private static boolean validPalindrome(String s) {
        int left = 0, right = s.length() - 1;

        while (left < right) {
            if (s.charAt(left) != s.charAt(right)) {
                return isPalindrome(s, left + 1, right) || isPalindrome(s, left, right - 1);
            }
            
            left++;
            right--;
        }

        return true;
    }
}
