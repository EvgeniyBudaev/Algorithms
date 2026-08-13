public class Training {
    public static void main(String[] args) {
        int[] arr = {10, 2, 5, 3};
        System.out.println(checkIfExist(arr));
    }

    private static boolean checkIfExist(int[] arr) {
        int left = 0, right = 1;

        while (left < arr.length - 1) {
            if (arr[left] == arr[right] * 2 || arr[right] == arr[left] * 2) {
                return true;
            } else if (right == arr.length - 1) {
                left++;
                right = left + 1;
            } else {
                right++;
            }
        }

        return false;
    }
}
