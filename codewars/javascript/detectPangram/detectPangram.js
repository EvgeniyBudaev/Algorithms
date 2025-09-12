/* Панграмма — это предложение, в котором каждая буква алфавита встречается хотя
 бы по одному разу. */

// isPangram возвращает true, если строка является панграммой, иначе false.
// time: O(n), space: O(1), n - длина строки
function isPangram(string) {
	const alphabet = 'abcdefghijklmnopqrstuvwxyz';
	return [...alphabet].every(letter => string.toLowerCase().includes(letter));
}

// function isPangram(string){
// 	const REGEX = /([a-z])(?!.*\1)/g;
// 	return (string.toLowerCase().match(REGEX) || []).length === 26;
// }

// function isPangram(string) {
// 	const alphabetList = [...'abcdefghijklmnopqrstuvwxyz'];
// 	return alphabetList.every((letter) => string.toLowerCase().includes(letter));
// }

console.log(isPangram('The quick brown fox jumps over the lazy dog.')); // true
console.log(isPangram('This is not a pangram.')); // false
console.log(isPangram('Cwm fjord bank glyphs vext quiz is a pangram')); // true
