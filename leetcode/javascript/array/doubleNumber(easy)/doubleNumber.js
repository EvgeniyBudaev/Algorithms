function doubleNumber(array) {
  return array.map(item => item * 2);
}

const arr = [17, 23, 31, 44, 59];
console.log(doubleNumber(arr)); // возвращает [34,46,62,88,118];
