/*
Есть массив, длина которого заранее неизвестна.
Допишите функцию, которая вернёт последний элемент переданного массива.
*  */

function getLastElementOfArray(array) {
  return array[array.length - 1];
}

var words = ['Ночь', 'Улица', 'Фонарь', 'Аптека', 'Бессмысленный', 'Тусклый', 'Свет'];
console.log(getLastElementOfArray(words)); // возвращает Свет;
