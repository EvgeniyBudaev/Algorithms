function separateArray(array) {
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

const numbers = [1, 2, 3, 4, 5, 6];
console.log(separateArray(numbers)); // { even: [2, 4, 6], odd: [1, 3, 5] }
