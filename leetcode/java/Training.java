
public class Training {
    public static void main(String[] args) {
        int[] arr = {0, 3, 2, 1};
        System.out.println(validMountainArray(arr)); // true
    }

    public static boolean validMountainArray(int[] arr) {
        if (arr == null || arr.length < 3) {
            return false;
        }

        int left = 0, right = arr.length - 1;

        // Двигаем указатели навстречу друг другу
        while (left < right) {
            if (arr[left] < arr[left + 1]) {
                left++;       // Идем вверх по левому склону
            } else if (arr[right] < arr[right - 1]) {
                right--;      // Идем вверх по правому склону
            } else {
                // Если ни один из указателей не может двигаться 
                // (склон плоский, яма или указатели пересеклись не на пике)
                return false;
            }
        }

        // Указатели встретились на пике. 
        // Пик не должен находиться на самом краю массива (иначе это просто склон без спуска/подъема).
        return left != 0 && right != arr.length - 1;
    }
}
