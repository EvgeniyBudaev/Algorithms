/*
        Ваш коллега написал функцию. Её задача — находить такое число,
        при котором перемножение этого числа и всех предшествующих ему значений
        до единицы не будут превышать переданное в функцию значение.
        В математике операция по умножению числа на все предшествующие ему
        значения до единицы называется факториалом числа. Например, если
        переданное значение равно 25, то ближайший факториал будет
        равен 4. (1*2*3*4 = 24), а если переданное значение равно 137,
        то результатом функции будет 5 (1*2*3*4*5 = 120). Но вот незадача, при
        вводе значений в поле вызывается эта функция и её вызов приводит к
        ошибке: ReferenceError: i is not defined. Поправьте написанный код
        функции так, чтобы ошибка больше не проявилась.
*/

const form = document.forms.formWithInput;
const printResult = document.querySelector('.content__result');
const inputNumber = document.querySelector('.form__input');
const congratulation = document.querySelector('#congratulation');

function findNearestFactorial(value) {
  if (value === '' || value === undefined) {
    return '*';
  }
  if (value <= 0) {
    console.log(value);
    return 1;
  }
  let currentValue = 1;
  let i;
  for (i = 1; currentValue * i <= value; i++) {
    currentValue *= i;
  }
  return i - 1;
}

form.addEventListener('keyup', evt => {
  evt.preventDefault();
  printResult.textContent = findNearestFactorial(inputNumber.value);
  congratulation.textContent = inputNumber.value && 'Вау, это успех!';
});

form.addEventListener('submit', evt => {
  evt.preventDefault();
});
