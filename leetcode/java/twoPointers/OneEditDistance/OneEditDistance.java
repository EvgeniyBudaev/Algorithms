package twoPointers.OneEditDistance;

/* 161. One Edit Distance
https://github.com/EvgeniyBudaev/doocs_leetcode/tree/main/solution/0100-0199/0161.One%20Edit%20Distance

Учитывая две строки s и t, определите, находятся ли они на расстоянии одного редактирования друг от друга.

s Вставьте в него ровно один символ , чтобы получить t
s Удалите из него ровно один символ, чтобы получить t
Замена ровно одного символа другим символом в s дает t

Input: s = "ab", t = "acb"
Output: true
Объяснение: в строку s можно вставить 'c',  чтобы получить t.

Input: s = "cab", t = "ad"
Output: false
 */
public class OneEditDistance {

    public static void main(String[] args) {
        System.out.println(isOneEditDistance("ab", "acb")); // true
        System.out.println(isOneEditDistance("cab", "ad")); // false
    }

    // isOneEditDistance определяет, находятся ли строки на расстоянии одного редактирования друг от друга.
    // time: O(n), space: O(1)
    private static boolean isOneEditDistance(String s, String t) {
        int m = s.length(), n = t.length();

        // Меняем местами строки, чтобы s всегда была длиннее строки t.
        if (m < n) {
            return isOneEditDistance(t, s);
        }

        // Если разница в длинах больше 1, то строки не могут быть на расстоянии одного редактирования друг от друга.
        if (m - n > 1) {
            return false;
        }

        // s: "acb", t: "ab"
        for (int i = 0; i < n; i++) {
            // Если символы строки s и t не совпадают, то проверяем, что строка s имеет длину n+1.
            if (s.charAt(i) != t.charAt(i)) {
                // Если строки имеют одинаковую длину, то проверяем, что строки равны, начиная с i+1 символа.
                if (m == n) {
                    return s.substring(i + 1).equals(t.substring(i + 1));
                }
                // Если строки имеют разную длину, то проверяем, что строка s имеет длину n+1.
                // s[i+1:]: "b" t[i:]: "b
                return s.substring(i + 1).equals(t.substring(i));
            }
        }

        // Если все символы строки s и t совпадают, то проверяем, что строка s имеет длину n+1.
        return m == n + 1;
    }
}
