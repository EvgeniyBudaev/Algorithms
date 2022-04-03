/* Даны два массива чисел длины n. Составьте из них один массив длины 2n,
в котором числа из входных массивов чередуются (первый — второй — первый —
второй — ...). При этом относительный порядок следования чисел из одного массива
 должен быть сохранён. */

const list1 = [1, 2, 3];
const list2 = [4, 5, 6];

const alternationValuesFromArrays = (arr1, arr2) => {
	const result = [];

	for (let i = 0; i < arr1.length; i++) {
		result.push(arr1[i]);
		result.push(arr2[i]);
	}

	return result;
};

console.log(alternationValuesFromArrays(list1, list2));
