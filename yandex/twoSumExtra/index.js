/* TwoSum. Дан массив целых чисел numbers и целое число x.
Нужно найти в массиве два элемента, сумма которых равняется x  */

// const initialData = [2, 1, 3, 5, 5]; // [1, 3]
const initialData = [-3, 1, 1, 2, 6, 6, 8, 10]; // []

/* Вариант №2. Оптимизированный алгоритм. Скорость быстрее, чем у варианта №1 */
const twoSum_extra_ds = (numbers, x) => {
	// Создаём вспомогательную структуру данных с быстрым поиском элемента.
	const previous = new Set();

	for (let i = 0; i < numbers.length; i++) {
		const current = numbers[i];
		const need = x - current;

		if (previous.has(need)) {
			return [current, need];
		} else {
			previous.add(current);
		}
	}

	return [];
};

// console.log(twoSum_extra_ds(initialData, 4));
console.log(twoSum_extra_ds(initialData, 100));
