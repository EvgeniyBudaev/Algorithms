public class Training {
    public static void main(String[] args) {
        System.out.println(checkIfPalindrome("racecar")); // true
        System.out.println(checkIfPalindrome("aleba")); // false
    }

    private static boolean checkIfPalindrome(String s) {
        int left = 0, right = s.length() - 1;

        while (left < right) {
            if (s.charAt(left) != s.charAt(right)) {
                return false;
            }
            left++;
            right--;
        }
        return true;
    }
}
