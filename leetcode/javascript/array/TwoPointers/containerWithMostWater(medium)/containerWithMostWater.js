/* https://leetcode.com/problems/container-with-most-water/description/
Вам дан целочисленный массив высотой длины n. Нарисовано n вертикальных линий, конечными точками которых являются
 (i, 0) и (i, height[i]).
Найдите две линии, которые вместе с осью X образуют контейнер, в котором содержится больше всего воды.
Возвращайте максимальное количество воды, которое может хранить контейнер.
Обратите внимание, что вы не можете наклонять контейнер

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Пояснение: Вышеупомянутые вертикальные линии представлены массивом [1,8,6,2,5,4,8,3,7].
В этом случае максимальная площадь воды (синяя секция), которую может содержать контейнер, равна 49.

Input: height = [1,1]
Output: 1

Complexity
Time complexity: O(n)
Space complexity: O(1)
*/

/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function(height) {
  let maxArea = 0, left = 0, right = height.length - 1;

  while (left < right) {
    const minHeight = Math.min(height[left], height[right]);
    const width = right - left;
    maxArea = Math.max(maxArea, width * minHeight);
    if (height[left] < height[right]) {
      left++;
    } else {
      right--;
    }
  }

  return maxArea;
};
console.log(maxArea([1,8,6,2,5,4,8,3,7])); // 49