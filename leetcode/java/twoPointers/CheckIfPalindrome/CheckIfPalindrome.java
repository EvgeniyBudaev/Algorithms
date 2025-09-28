package twoPointers.CheckIfPalindrome;

/*
Учитывая строку s, верните true, если это палиндром, и false в противном случае.
Строка является палиндромом, если она читается как вперед, так и назад одинаково.
*/

public class CheckIfPalindrome {
    public static void main(String[] args) {
        System.out.println(checkIfPalindrome("racecar")); // true
        System.out.println(checkIfPalindrome("aleba")); // false
    }

    // checkIfPalindrome проверяет, является ли строка палиндромом.
    // time: O(n), space: O(1)
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
