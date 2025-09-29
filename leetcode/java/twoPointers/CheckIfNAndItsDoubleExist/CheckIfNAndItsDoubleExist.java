package twoPointers.CheckIfNAndItsDoubleExist;

/* 1346. Check If N and Its Double Exist
https://leetcode.com/problems/check-if-n-and-its-double-exist/description/

Учитывая массив целых чисел, проверьте, существуют ли два индекса i и j такие, что:
i != j
0 <= i, j < arr.length
arr[i] == 2 * arr[j]

Input: arr = [10,2,5,3]
Output: true
Explanation: For i = 0 and j = 2, arr[i] == 10 == 2 * 5 == 2 * arr[j]

Input: arr = [3,1,7,11]
Output: false
Пояснение: Не существует i и j, удовлетворяющих этим условиям.
*/

public class CheckIfNAndItsDoubleExist {
    public static void main(String[] args) {
        int[] arr = {10, 2, 5, 3};
        System.out.println(checkIfExist(arr)); // true
    }

    // checkIfExist проверяет, существует ли в массиве arr два различных индекса i и j,
    // таких что один элемент равен удвоенному другому (т.е. arr[i] == 2 * arr[j] или arr[j] == 2 * arr[i]).
    // time: O(n^2), space: O(1)
    private static boolean checkIfExist(int[] arr) {
        int left = 0, right = 1; // инициализируем два указателя на начало и конец массива

        while (left < arr.length - 1) { // запускаем цикл для перемещения левого указателя
            if (arr[left] == arr[right] * 2 || arr[right] == arr[left] * 2) { // проверяем условие на равенство
                return true;
                // если достигнут конец массива, перемещаем левый указатель вправо и обновляем правый
            } else if (right == arr.length - 1) {
                left++;
                right = left + 1;
            } else {
                right++; // перемещаем правый указатель вправо
            }
        }

        return false; // если не найдено ни одной пары, возвращаем false
    }
}
