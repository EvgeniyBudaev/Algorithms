const data = [
  {name: "Саша", age: 19},
  {name: "Катя", age: 21},
  {name: "Миша", age: 17},
  {name: "Федя", age: 23},
  {name: "Клава", age: 22}
];


function calcAvgAge(array) {
  // Напишите код здесь
  const average = array.reduce((total, next) => total + next.age, 0) / array.length;
  console.log("Средний возраст: ", average);
}

calcAvgAge(data);
