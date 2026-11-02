import java.util.Arrays;

public class Training {
    public static void main(String[] args) {
        int[] arr = {1, 0, 2, 3, 0, 4, 5, 0};
        duplicateZeros(arr);
        System.out.println(Arrays.toString(arr));
    }

    private static void duplicateZeros(int[] arr) {
        for (int i = 0; i < arr.length; i++) {
            if (arr[i] == 0 && arr.length != 0) {
                System.arraycopy(arr, i, arr, i + 1, arr.length - i - 1);
                i++;
            }
        }
    }
}
