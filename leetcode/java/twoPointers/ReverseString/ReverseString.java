package twoPointers.ReverseString;

/* 344. Reverse String
https://leetcode.com/problems/reverse-string/description/

Напишите функцию, которая переворачивает строку. Входная строка задается как массив символов s.

Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
*/

public class ReverseString {
    public static void main(String[] args) {
        char[] s1 = {'h', 'e', 'l', 'l', 'o'};
        reverseString(s1);
        System.out.println(new String(s1)); // ["o","l","l","e","h"]
    }

    // reverseString переворачивает строку.
    // time: O(n), space: O(1)
    private static void reverseString(char[] s) {
        int left = 0, right = s.length - 1;

        while (left < right) {
            char temp = s[left];
            s[left] = s[right];
            s[right] = temp;
            left++;
            right--;
        }
    }
}
