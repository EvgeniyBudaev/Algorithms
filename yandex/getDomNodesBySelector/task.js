/*
Напишите функцию, которая будет возвращать коллекцию DOM-узлов всех цен на
основании переданного селектора в виде Array (не NodeList).
 */

function getDomNodesBySelector(selector) {
  // Ваш код
  return [...document.querySelectorAll(selector)];
}

console.log(getDomNodesBySelector('.price-value'));