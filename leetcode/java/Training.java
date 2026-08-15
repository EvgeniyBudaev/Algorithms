
public class Training {
    public static void main(String[] args) {
        int[] height = {1, 8, 6, 2, 5, 4, 8, 3, 7};
        System.out.println(maxArea(height)); // 49
    }

    public static int maxArea(int[] height) {
        int maxAreaContainer = 0;
        int left = 0, right = height.length - 1;

        while (left < right) {
            int minHeight = Math.min(height[left], height[right]);
            int width = right - left;
            int currentArea = minHeight * width;
            maxAreaContainer = Math.max(maxAreaContainer, currentArea);

            if (height[left] < height[right]) {
                left++;
            } else {
                right--;
            }
        }

        return maxAreaContainer;
    }
}
