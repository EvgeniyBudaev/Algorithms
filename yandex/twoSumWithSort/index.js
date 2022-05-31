/* TwoSum. Дан массив целых чисел numbers и целое число x.
Нужно найти в массиве два элемента, сумма которых равняется x  */

const initialData = [2, 1, 3, 5, 5]; // [1, 3]

/* Вариант №1. Оптимизированный алгоритм. Скорость N-1 */
const twoSum_with_sort = (numberList, x) => { // 1 и 3
	const result = [];
	const sortedList = numberList.sort((a, b) => a - b); // [1, 2, 3, 5, 5]
	let left = 0;
	let right = sortedList.length - 1;

	while (left < right) {
		const sum = sortedList[left] + sortedList[right];
		if (sum === x) {
			result.push([sortedList[left], sortedList[right]]);
			break;
		} else if (sum < x) {
			left++;
		} else {
			right--;
		}
	}

	if (result.length > 0) {
		if (result.length === 1) {
			return result[0];
		} else {
			return result[Math.floor(Math.random() * result.length)];
		}
	} else {
		return [];
	}
};

console.log(twoSum_with_sort(initialData, 4));
