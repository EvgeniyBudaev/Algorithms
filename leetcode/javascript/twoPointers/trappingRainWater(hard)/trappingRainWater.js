/* 42. Trapping Rain Water
https://leetcode.com/problems/trapping-rain-water/description/

Учитывая n неотрицательных целых чисел, представляющих карту высот, где ширина каждой полосы равна 1, вычислите,
сколько воды она может удержать после дождя.

Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Пояснение: Приведенная выше карта высот (черное сечение) представлена массивом [0,1,0,2,1,0,1,3,2,1,2,1].
В данном случае задерживается 6 единиц дождевой воды (синяя секция).
*/

/**
 * @param {number[]} height
 * @return {number}
 */
var trap = function (height) {
    // Если массив пустой, воды быть не может
    if (!height || height.length === 0) return 0;

    let left = 0;                      // левая граница
    let leftMaxValue = height[left];   // максимальная высота слева
    let right = height.length - 1;     // правая граница
    let rightMaxValue = height[right]; // максимальная высота справа
    let sum = 0;                       // общее количество удержанной воды

    while (left < right) {
        // Если максимальная высота слева меньше или равна максимальной высоте справа, 
        // то уровень воды в текущей левой ячейке ограничивается именно leftMaxValue.
        if (leftMaxValue <= rightMaxValue) {
            sum += leftMaxValue - height[left];
            left++;
            leftMaxValue = Math.max(leftMaxValue, height[left]);
        } else {
            // Иначе уровень воды ограничивается rightMaxValue
            sum += rightMaxValue - height[right];
            right--;
            rightMaxValue = Math.max(rightMaxValue, height[right]);
        }
    }

    return sum;
};