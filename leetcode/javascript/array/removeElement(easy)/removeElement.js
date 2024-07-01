// https://leetcode.com/problems/remove-element/description/

/*
Учитывая целочисленный массив nums и целочисленное значение, удалите все вхождения val в nums на месте. Порядок
элементов может быть изменен. Затем верните количество элементов в виде чисел, которые не равны val.
Учитывайте количество элементов в nums, которые не равны val be k. Чтобы вас приняли, вам необходимо сделать следующее:
Измените массив nums так, чтобы первые k элементов nums содержали элементы, не равные val. Остальные элементы nums не
важны, как и размер nums.
Вернуть К.
*/

/* Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]  */

/* Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]  */

/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
function removeElement(nums, val) {
	let count = 0; // Счетчик для отслеживания элементов, отличных от val
	for (let i = 0; i < nums.length; i++) {
		if (nums[i] !== val) {
			nums[count++] = nums[i];
		}
	}
	return count;
}

// var removeElement = function(nums, val) {
// 	if (!nums) return 0;
// 	for (let i = 0; i < nums.length; i++) {
// 		if (nums[i] === val) {
// 			nums.splice(i, 1);
// 			i--;
// 		}
// 	}
// 	return nums.length;
// };

const initialData = [3,2,2,3];
console.log(removeElement(initialData, 3)); // 2
