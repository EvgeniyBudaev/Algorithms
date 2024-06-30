/*
Есть массив, длина которого заранее неизвестна.
Допишите функцию, которая вернёт последний элемент переданного массива.
*  */

var words = ['Ночь', 'Улица', 'Фонарь', 'Аптека', 'Бессмысленный', 'Тусклый', 'Свет'];

function getLastElementOfArray(array) {
  return array[array.length - 1];
}

console.log(getLastElementOfArray(words)); // возвращает Свет;
