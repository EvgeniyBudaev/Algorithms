/*
Есть массив, длина которого заранее неизвестна.
Допишите функцию, которая вернёт последний элемент переданного массива.
*  */

var words = ['Ночь', 'Улица', 'Фонарь', 'Аптека', 'Бессмысленный', 'Тусклый', 'Свет'];

getLastElementOfArray(words); // возвращает Свет;

function getLastElementOfArray(array) {
  // Напишите код здесь
  const result = array[array.length -1];
  console.log(result);
}
