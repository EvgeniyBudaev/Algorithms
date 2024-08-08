/* https://leetcode.com/problems/trapping-rain-water/description/

Учитывая n неотрицательных целых чисел, представляющих карту высот, где ширина каждой полосы равна 1, вычислите,
сколько воды она может удержать после дождя.

Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Пояснение: Приведенная выше карта высот (черное сечение) представлена массивом [0,1,0,2,1,0,1,3,2,1,2,1].
В данном случае задерживается 6 единиц дождевой воды (синяя секция).
*/

/**
 * Time O(N) | Space O(1)
 * @param {number[]} height
 * @return {number}
 */
var trap = function(height) {
    let left = 0, leftMaxValue = height[0];
    let right = height.length - 1, rightMaxValue = height[right];
    let sum = 0;

    while (left < right) {
        if (leftMaxValue <= rightMaxValue) {
            sum += leftMaxValue - height[left];
            left++;
            leftMaxValue = Math.max(leftMaxValue, height[left]);
        } else {
            sum += rightMaxValue - height[right];
            right--;
            rightMaxValue = Math.max(rightMaxValue, height[right]);
        }
    }

    return sum;
};

console.log(trap([0,1,0,2,1,0,1,3,2,1,2,1])); // 6