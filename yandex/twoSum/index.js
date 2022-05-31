/* TwoSum. Дан массив целых чисел numbers и целое число x.
Нужно найти в массиве два элемента, сумма которых равняется x  */

const initialData = [2, 1, 3, 5, 5];

/* Наивный способ. Медленно работает. Скорость N2 / 4 */
const twoSum = (numberList, x) => { // 1 и 3
	const result = [];

	for (let i = 0; i < numberList.length; i++) {
		for (let j = i + 1; j < numberList.length; j++) {
			if (numberList[i] + numberList[j] === x) {
				result.push(numberList[i]);
				result.push(numberList[j]);
			}
		}
	}

	if (result.length === 0) {
		return 'Нет такого числа';
	}

	return result;
};

console.log(twoSum(initialData, 4));
