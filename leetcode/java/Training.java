import java.util.Arrays;

import twoPointers.SquaresOfASortedArray.SquaresOfASortedArray;

public class Training {
    public static void main(String[] args) {
        int[] nums = new int[]{-4, -1, 0, 3, 10};
        // Создаем экземпляр класса для вызова нестатического метода
        SquaresOfASortedArray solution = new SquaresOfASortedArray();
        System.out.println(Arrays.toString(solution.sortedSquares(nums))); // [0,1,9,16,100]
    }

    private int[] sortedSquares(int[] nums) {
        int n = nums.length; // Длина массива
        int[] result = new int[n];
        int left = 0; // Индекс первого элемента массива nums
        int right = n - 1; // Индекс последнего элемента массива nums
        int index = n - 1; // Индекс последнего элемента массива result

        while (left <= right) {
            int leftSquare = nums[left] * nums[left];
            int rightSquare = nums[right] * nums[right];

            if (leftSquare > rightSquare) {
                result[index] = leftSquare;
                left++;
            } else {
                result[index] = rightSquare;
                right--;
            }
            index--;
        }

        return result;
    }
}
