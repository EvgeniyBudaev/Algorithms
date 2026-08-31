import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Objects;

public class Training {
    public static void main(String[] args) {
        int[] nums1 = {1, 2, 2, 1};
        int[] nums2 = {2, 2};
        System.out.println(Arrays.toString(intersect(nums1, nums2))); // [2, 2]
    }

    private static int[] intersect(int[] nums1, int[] nums2) {
        Arrays.sort(nums1); // [1, 1, 2, 2]
        Arrays.sort(nums2); // [2, 2]

        int left = 0, right = 0;
        List<Integer> list = new ArrayList<>();

        // time: O(n), space: O(n)
        while (left < nums1.length && right < nums2.length) {
            if (nums1[left] < nums2[right]) {
                left++;
            } else if (nums1[left] > nums2[right]) {
                right++;
            } else {
                list.add(nums1[left]);
                left++;
                right++;
            }
        }

        // Конвертируем List в массив
        int[] result = list.stream()
                  .filter(Objects::nonNull)
                  .mapToInt(Integer::intValue)
                  .toArray();

        return result;
    }
}
