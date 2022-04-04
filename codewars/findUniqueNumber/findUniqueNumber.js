/* Функция findUniq находит уникальное значение в массиве */

function findUniq(arr) {
  return arr.filter(item=> arr.indexOf(item) === arr.lastIndexOf(item))[0];
}

console.log(findUniq([ 1, 1, 1, 2, 1, 1 ])); // 2
