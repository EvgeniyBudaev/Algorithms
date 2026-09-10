import java.util.Arrays;

public class Training {
    public static void main(String[] args) {
        int[] nums1 = {2, 0, 2, 1, 1, 0};
        sortColors(nums1);
        System.out.println(Arrays.toString(nums1)); // [0,0,1,1,2,2]

        int[] nums2 = {2, 0, 1};
        sortColors(nums2);
        System.out.println(Arrays.toString(nums2)); // [0,1,2]
    }

        private static void sortColors(int[] nums) {
        int i = 0; // Индекс текущего элемента
        int left = 0, right = nums.length - 1; // Левая и правая границы

        while (left < right && i <= right) { // пока левая граница меньше правой и индекс меньше правой границы
            // Если текущий элемент равен 0, то меняем его местами с элементом в начале массива
            if (nums[i] == 0) {
                swap(nums, left, i);
                left++;
                i++;
                // Если текущий элемент равен 2, то меняем его местами с элементом в конце массива
            } else if (nums[i] == 2) {
                swap(nums, right, i);
                right--;
                // Если текущий элемент равен 1, то оставляем его на месте и переходим к следующему элементу
            } else {
                i++;
            }
        }
    }

    private static void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }
}
