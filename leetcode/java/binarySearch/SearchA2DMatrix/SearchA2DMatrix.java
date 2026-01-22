package binarySearch.SearchA2DMatrix;

/* 74. Search a 2D Matrix
https://leetcode.com/problems/search-a-2d-matrix/description/

Вам дана целочисленная матрица размера m x n со следующими двумя свойствами:
Каждая строка сортируется в порядке неубывания.
Первое целое число каждой строки больше последнего целого числа предыдущей строки.
Учитывая целочисленную цель, верните true, если цель находится в матрице, или false в противном случае.
Вы должны написать решение с временной сложностью O(log(m * n)).

Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true

Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false
*/

public class SearchA2DMatrix {
    public static void main(String[] args) {
        int[][] matrix = {{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}};
        System.out.println(searchMatrix(matrix, 3)); // true
    }

    // searchMatrix возвращает true, если цель находится в матрице, иначе false.
    private static boolean searchMatrix(int[][] matrix, int target) {
        if (matrix.length == 0 || matrix[0].length == 0) {
            return false;
        }

        for (int[] row : matrix) {
            // Используем бинарный поиск для каждой строки
            int left = 0;
            int right = row.length - 1;
            
            while (left <= right) {
                int mid = (left + right) / 2;
                // Если текущий элемент равен целевому значению, возвращаем true
                if (row[mid] == target) {
                    return true;
                // Если текущий элемент меньше целевого значения, продолжаем поиск в правой половине строки
                } else if (row[mid] < target) {
                    left = mid + 1;
                } else {
                    right = mid - 1;
                }
            }
        }

        return false;
    }
}
