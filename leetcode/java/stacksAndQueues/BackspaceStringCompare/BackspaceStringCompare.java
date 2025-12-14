package stacksAndQueues.BackspaceStringCompare;

/* 844. Backspace String Compare
https://leetcode.com/problems/backspace-string-compare/description/

Учитывая две строки s и t, верните true, если они равны, когда обе вводятся в пустые текстовые редакторы. '#' означает
символ возврата.
Обратите внимание, что после возврата пустого текста текст останется пустым.

Input: s = "ab#c", t = "ad#c"
Output: true
Пояснение: И s, и t становятся "ac".

Input: s = "ab##", t = "c#d#"
Output: true
Пояснение: И s, и t становятся "".

Input: s = "a#c", t = "b"
Output: false
Пояснение: s становится «c», а t становится «b».
*/

public class BackspaceStringCompare {
    public static void main(String[] args) {
        System.out.println(backspaceCompare("ab#c", "ad#c")); // true
    }

    // backspaceCompare проверяет, равны ли две строки, когда обе вводятся в пустые текстовые редакторы, после символа возврата.
    // time: O(n + m), space: O(n + m), где n и m - длины строк s и t соответственно.
    private static boolean backspaceCompare(String s, String t) {
        return build(s).equals(build(t));
    }

    private static String build(String str) {
        StringBuilder sb = new StringBuilder();

        for (char c : str.toCharArray()) {
            if (c != '#') {
                sb.append(c);
            } else if (sb.length() > 0) {
                sb.deleteCharAt(sb.length() - 1);
            }
        }
        
        return sb.toString(); // ac ac for "ab#c", "ad#c"
    }
}
