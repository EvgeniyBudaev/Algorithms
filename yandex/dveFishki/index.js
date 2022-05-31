/* Нужно вывести два числа —– очки на двух фишках, в сумме дающие k.
Если таких пар несколько, то можно вывести любую из них.
Если таких пар не существует, то вывести [].  */

const initialData = [-1, -9, -7, 3, -6]; // [-1, 3]
// const initialData = [-93, 67, 84, -30, -96, 22, 40, -11, -39, 11]; // None

/* Наивный способ. Медленно работает. Скорость N2 / 4 */
function twoSum(array, targetSum) {
	const result = [];

	for (let i = 0; i < array.length; i++) {
		for (let j = i + 1; j < array.length; j++) {
			if (array[i] + array[j] === targetSum) {
				result.push([array[i], array[j]]);
			}
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
}

console.log(twoSum(initialData, 2));
// console.log(twoSum(initialData, -186));
