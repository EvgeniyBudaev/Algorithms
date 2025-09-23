package twoPointers.IsSubsequence;

/*
Учитывая две строки s и t, верните true, если s является подпоследовательностью t, или false в противном случае.

Подпоследовательность строки — это последовательность символов, которую можно получить путем удаления некоторых
(или ни одного) символов из исходной строки, сохраняя при этом относительный порядок оставшихся символов.
Например, «ace» является подпоследовательностью «abcde», а «aec» — нет.
*/

public class IsSubsequence {
    public static void main(String[] args) {
        System.out.println(isSubsequence("ace", "abcde")); // true
        System.out.println(isSubsequence("aec", "abcde")); // false
    }

    // isSubsequence проверяет, является ли строка s подпоследовательностью строки t.
    // time: O(n), space: O(1)
    private static boolean isSubsequence(String s, String t) {
        int left = 0, right = 0;
        
        while (left < s.length() && right < t.length()) { // пока не достигнут конец одной из строк
            if (s.charAt(left) == t.charAt(right)) { // если символы совпадают
                left++;
            }
            right++;
        }
        
        return left == s.length(); // если все символы из s были найдены в t, то возвращаем true
    }
}
