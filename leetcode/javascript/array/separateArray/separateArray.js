/* */

const numbers = [1, 2, 3, 4, 5, 6];

const result = separateArray(numbers);
// должен вернуть { even: [2, 4, 6], odd: [1, 3, 5] }

function separateArray(array) {
  // Ваш код здесь
  const even = [];
  const odd = [];
  array.forEach(item => {
    if (item % 2 === 0) {
      even.push(item);
    } else {
      odd.push(item);
    }
  })
  return { even: even, odd: odd };
}

console.log("Результат: ", result);
