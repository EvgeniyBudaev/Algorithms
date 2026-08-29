
public class Training {
    public static void main(String[] args) {
        int[] nums = {0, 1, 0, 3, 12};
        moveZeroes(nums);
        System.out.println(Arrays.toString(nums)); // [1, 3, 12, 0, 0]
    }

    private static void moveZeroes(int[] nums) {
        int left = 0; // Индекс для следующего ненулевого элемента

        for (int right = 0; right < nums.length; right++) {
            if (nums[right] != 0) { // Если текущий элемент не равен 0
                // Меняем местами текущий элемент и следующий ненулевой элемент
                int temp = nums[right];
                nums[right] = nums[left];
                nums[left] = temp;
                left++; // Увеличиваем индекс для следующего ненулевого элемента
            }
        }
    }
}
