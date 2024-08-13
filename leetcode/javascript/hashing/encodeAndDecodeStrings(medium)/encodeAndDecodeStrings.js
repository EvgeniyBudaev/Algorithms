/* 659 · Encode and Decode Strings

Разработайте алгоритм для кодирования списка строк в одну строку. Закодированная строка затем декодируется обратно в
исходный список строк.
Пожалуйста, реализуйте кодирование и декодирование

Input: ["neet","code","love","you"]
Output:["neet","code","love","you"]

Input: ["we","say",":","yes"]
Output: ["we","say",":","yes"]
*/

/**
 * Chunk Transfer Encoding
 * Time O(N) | Space O(1)
 * https://leetcode.com/problems/encode-and-decode-strings/
 * @param {string[]} strs
 * @return {string}
 */
// Кодирует список строк в одну строку.
// Длина каждой строки сохраняется как префикс фиксированного размера перед фактическим содержимым строки.
var encode = (strs, sb = []) => {
    for (const str of strs) {/* Time O(N) */
        const code = getCode(str);/* Time O(1) */ // для "Hello" -> 00000101
        const encoding = `${code}${str}`;
        sb.push(encoding);
    }
    return sb.join(''); /* Time O(N) | Ignore Auxillary Space O(N) */
}

const getCode = (str) => str
    .length
    // Преобразование длины в двоичную систему: .toString(2) - метод преобразует число (в данном случае длину строки) в
    // строку, представляющую это число в двоичной системе счисления
    .toString(2)
    // Добавление ведущих нулей: .padStart(8, '0') - метод добавляет символы ('0' в данном случае) в начало строки до
    // тех пор, пока ее общая длина не достигнет указанного значения (8). Это делается для обеспечения фиксированной
    // ширины результата, что особенно полезно при работе с битовыми данными или форматировании данных.
    .padStart(8, '0');

/**
 * Chunk Transfer Encoding
 * Time O(N) | Space O(1)
 * https://leetcode.com/problems/encode-and-decode-strings/
 * @param {string} str
 * @return {string[]}
 */
// Декодирует одну строку в список строк, считывая префикс длины фиксированного размера
// и затем считываем соответствующее количество символов.
function decode(str, output = []) {
    for (let left = 0, right = (left + 8), length = 0;
         left < str.length;
         left = (right + length), right = (left + 8)
    ) {
        const countString = str.slice(left, right); // 00000101
        length = parseInt(countString, 2); // 5
        const decoding = str.slice(right, (right + length));
        output.push(decoding);
    }
    return output;
}

console.log("encode: ", encode(["Hello", "World"])); // "00000101Hello00000101World"
console.log("decode: ", decode("00000101Hello00000101World")); // ["Hello", "World"]