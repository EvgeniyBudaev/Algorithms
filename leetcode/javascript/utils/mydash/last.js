// _.last([1, 2, 3, 4, 5]); // => 5

const arr = [1, 2, 3];

const last = (array) => {
  const length = array == null ? 0 : array.length;
  return length ? array[length - 1] : undefined;
};

console.log(last(arr)); // => 1
