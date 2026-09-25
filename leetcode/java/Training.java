
public class Training {
    public static void main(String[] args) {
        int[] arr = {10, 2, 5, 3};
        System.out.println(checkIfExist(arr)); // true
    }

    private static boolean checkIfExist(int[] arr) {
        int left = 0, right = 1; // инициализируем два указателя на начало и конец массива

        while (left < arr.length - 1) { // запускаем цикл для перемещения левого указателя
            if (arr[left] == arr[right] * 2 || arr[right] == arr[left] * 2) { // проверяем условие на равенство
                return true;
                // если достигнут конец массива, перемещаем левый указатель вправо и обновляем правый
            } else if (right == arr.length - 1) {
                left++;
                right = left + 1;
            } else {
                right++; // перемещаем правый указатель вправо
            }
        }

        return false; // если не найдено ни одной пары, возвращаем false
    }
}
