import java.util.Arrays;

public class Training {
    public static void main(String[] args) {
        int[] numbers = new int[]{2, 7, 11, 15};
        System.out.println(Arrays.toString(twoSum(numbers, 9))); // [1,2]
    }

    private static int[] twoSum(int[] numbers, int target) {
        int left = 0, right = numbers.length - 1; // Инициализируем указатели на начало и конец массива

        while (left < right) { // Продолжаем поиск, пока левый указатель не станет больше правого
            int sum = numbers[left] + numbers[right]; // Вычисляем сумму текущих двух чисел
            if (sum == target) { // Если сумма равна целевому значению, возвращаем индексы чисел
                return new int[]{left + 1, right + 1};
            } else if (sum < target) { // Если сумма меньше целевого значения, сдвигаем левый указатель вправо
                left++;
            } else { // Если сумма больше целевого значения, сдвигаем правый указатель влево
                right--;
            }
        }

        return new int[]{-1, -1}; // Если не найдено ни одного решения, вернуть [-1,-1]
    }
}
