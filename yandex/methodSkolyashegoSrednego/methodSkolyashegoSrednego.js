/* Метод скользящего среднего */

const initialData = [4, 3, 8, 1, 5, 6, 3];

/* Оптимизированный алгоритм */
const averageMoving = (array, k) => { // [5, 4, 4.67, 4, 4.67]
	const result = [];
	let current_sum = 0;

	for (let i = 0; i < k; i++) {
		current_sum += array[i];
	}
	const current_avg = (current_sum / k).toFixed(2);
	result.push(Number(current_avg));

	for (let j = 0; j < array.length-k+1; j++) {
		current_sum = current_sum - array[j] + array[j+k];
		const current_avg = (current_sum / k).toFixed(2);
		if (current_avg === "NaN") {
			continue;
		}
		result.push(Number(current_avg));
	}
	return result;
};

console.log(averageMoving(initialData, 3));
