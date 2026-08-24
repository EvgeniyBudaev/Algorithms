package twoPointers.ValidMountainArray;

/* 941. Valid Mountain Array
https://leetcode.com/problems/valid-mountain-array/

Учитывая массив целых чисел arr, верните true если и только если он является допустимым массивом гор.
Напомним, что arr является горным массивом тогда и только тогда, когда:

arr.length >= 3
Существует некоторое i с 0 < i < arr.length - 1 таким , что:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

Input: arr = [2,1]
Output: false

Input: arr = [3,5,5]
Output: false

Input: arr = [0,3,2,1]
Output: true
*/

public class ValidMountainArray {
    public static void main(String[] args) {
        int[] arr = {0, 3, 2, 1};
        System.out.println(validMountainArray(arr)); // true
    }

    // validMountainArray проверяет, является ли массив горным массивом.
    // time: O(n), space: O(1)
    public static boolean validMountainArray(int[] arr) {
        if (arr == null || arr.length < 3) {
            return false;
        }

        int left = 0, right = arr.length - 1;

        // Двигаем указатели навстречу друг другу
        while (left < right) {
            if (arr[left] < arr[left + 1]) {
                left++;       // Идем вверх по левому склону
            } else if (arr[right] < arr[right - 1]) {
                right--;      // Идем вверх по правому склону
            } else {
                // Если ни один из указателей не может двигаться 
                // (склон плоский, яма или указатели пересеклись не на пике)
                return false;
            }
        }

        // Указатели встретились на пике. 
        // Пик не должен находиться на самом краю массива (иначе это просто склон без спуска/подъема).
        return left != 0 && right != arr.length - 1;
    }
}
