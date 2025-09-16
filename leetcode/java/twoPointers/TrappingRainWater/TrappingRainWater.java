package twoPointers.TrappingRainWater;

/* 42. Trapping Rain Water
https://leetcode.com/problems/trapping-rain-water/description/

Учитывая n неотрицательных целых чисел, представляющих карту высот, где ширина каждой полосы равна 1, вычислите,
сколько воды она может удержать после дождя.

Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Пояснение: Приведенная выше карта высот (черное сечение) представлена массивом [0,1,0,2,1,0,1,3,2,1,2,1].
В данном случае задерживается 6 единиц дождевой воды (синяя секция).
*/

public class TrappingRainWater {
    public static void main(String[] args) {
        int[] height = {0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1};
        System.out.println(trap(height)); // 6
    }

    /**
     * Вычисляет количество воды, которое может быть задержано после дождя
     * time: O(n), space: O(1)
     * @param height массив высот
     * @return количество задержанной воды
     */
    public static int trap(int[] height) {
        if (height == null || height.length == 0) {
            return 0;
        }
        
        int left = 0; // левая граница
        int leftMaxValue = height[left]; // максимальное количество воды слева
        int right = height.length - 1; // правая граница
        int rightMaxValue = height[right]; // максимальное количество воды справа
        int sum = 0; // общее количество воды, которое может быть залита
        
        while (left < right) { // пока левая граница не сравняется с правой
            // если максимальная высота слева меньше максимальной высоты справа, двигаем левую границу вправо
            if (leftMaxValue <= rightMaxValue) {
                sum += leftMaxValue - height[left]; // вычисляем количество воды, которое может быть залита в текущей ячейке
                left++; // двигаем левую границу вправо
                leftMaxValue = Math.max(leftMaxValue, height[left]); // обновляем максимальную высоту слева
            } else {
                sum += rightMaxValue - height[right]; // вычисляем количество воды, которое может быть залита в текущей ячейке
                right--; // двигаем правую границу влево
                rightMaxValue = Math.max(rightMaxValue, height[right]); // обновляем максимальную высоту справа
            }
        }
        
        return sum; // возвращаем общее количество воды, которое может быть залита
    }
}
