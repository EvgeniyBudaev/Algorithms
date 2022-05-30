/*
Допишите функцию в файле index.js.
При клике на кнопку «Использовать купон на 15%» вызывается функция
applyDiscount, она должна применить скидку 15% к каждому товару в корзине и
изменить итоговую стоимость заказа. Помните, что скидка должна примениться один
раз. Для решения этой задачи вам поможет функция из предыдущего задания —
getDomNodesBySelector.
*/

let count = 0;

function getDomNodesBySelector(selector) {
  return Array.from(document.querySelectorAll(selector));
}

function applyDiscount() {
  // Ваш код
  if (count >= 1) return;
  count += 1;
  const discount = 0.85;
  const totalPrice = document.querySelector(".total-price-value");
  totalPrice.innerHTML = totalPrice.textContent * discount;
  const priceList = getDomNodesBySelector('.price-value');
  priceList.forEach(node => {
    node.innerText = node.textContent * discount;
  });
}

document.querySelector('.total__button').addEventListener('click', applyDiscount);
