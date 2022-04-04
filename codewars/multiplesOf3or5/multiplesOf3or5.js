/* Если мы перечислим все натуральные числа до 10, кратные 3 или 5,
 мы получим 3, 5, 6 и 9. Сумма этих кратных равна 23. */

function solution(number){
	let sum = 0;
	for (let i = 1; i < number; i++) {
		if (!(i % 3) || !(i % 5)) {
			sum += i;
		}
	}
	return sum;
}

// function solution(number){
// 	return number < 1 ? 0 : [...new Array(number).keys()]
// 		.filter(n => n % 3 === 0 || n % 5 === 0)
// 		.reduce((acc, item) => acc + item);
// }

console.log(solution(10)); // 23
