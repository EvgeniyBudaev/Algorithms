
public class Training {
    public static void main(String[] args) {
        int[] nums = {1, 2, 4, 6, 8, 9, 14, 15};
        System.out.println(checkForTarget(nums, 13)); // true
    }

    private static boolean checkForTarget(int[] nums, int target) {
        int left = 0, right = nums.length - 1; // левая и правая границы массива

        while (left < right) { // пока левая граница меньше правой
            int sum = nums[left] + nums[right]; // сумма пары
            if (sum == target) {
                return true;
            } else if (sum > target) { // если сумма больше целевого значения, то уменьшаем правую границу
                right--;
            } else { // если сумма меньше целевого значения, то увеличиваем левую границу
                left++;
            }
        }

        return false; // если не найдено
    }
}
