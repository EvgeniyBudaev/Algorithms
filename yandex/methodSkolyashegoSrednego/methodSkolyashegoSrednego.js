/* Метод скользящего среднего */

const initialData = [4, 3, 8, 1, 5, 6, 3];

const averageMoving = (data, k) => {
	const result = [];
	for (let begin_index = 0; begin_index < data.length-k+1; begin_index++) {
		const end_index = begin_index + k;
		let current_sum = 0;
		for (let j = begin_index; j < end_index; j++) {
			current_sum += data[j];
		}
		const current_avg = (current_sum / k).toFixed(2);
		if (current_avg === "NaN") {
			continue;
		}
		result.push(Number(current_avg));
	}
	return result;
};

console.log(averageMoving(initialData, 3));
