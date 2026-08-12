public class Training {
    public static void main(String[] args) {
        int[] nums = {1, 2, 4, 6, 8, 9, 14, 15};
        System.out.println(checkForTarget(nums, 13));
    }

    private static boolean checkForTarget(int[] nums, int target) {
        int left = 0, right = nums.length - 1;

        while (left < right) {
            int sum = nums[left] + nums[right];
            if (sum == target) {
                return true;
            } else if (sum > target) {
                right--;
            } else {
                left++;
            }
        }

        return false;
    }
}
