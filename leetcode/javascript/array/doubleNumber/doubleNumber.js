const arr = [17, 23, 31, 44, 59];

doubleNumber(arr); // возвращает [34,46,62,88,118];

function doubleNumber(array) {
  // Ваш код
  const res = array.map(item => item * 2);
  console.log(res)
};
