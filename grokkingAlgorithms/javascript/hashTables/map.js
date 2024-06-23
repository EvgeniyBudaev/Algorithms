// Map – это коллекция ключ/значение. Но основное отличие в том, что Map позволяет использовать ключи любого типа.

const map = new Map();
map.set("1", "str1");    // строка в качестве ключа
map.set(1, "num1");      // цифра как ключ
map.set(true, "bool1");  // булево значение как ключ
console.log(map.get("1"));


// Перебор Map
const recipeMap = new Map([
    ["огурец", 500],
    ["помидор", 350],
    ["лук",    50]
]);

// перебор по ключам (овощи)
for (let vegetable of recipeMap.keys()) {
    console.log(vegetable); // огурец, помидор, лук
}

// перебор по значениям (числа)
for (let amount of recipeMap.values()) {
    console.log(amount); // 500, 350, 50
}

// перебор по элементам в формате [ключ, значение]
for (let entry of recipeMap) { // то же самое, что и recipeMap.entries()
    console.log(entry); // огурец,500 (и так далее)
}


// Object.entries: Map из Object
let obj = {
    name: "John",
    age: 30
};
// Object.entries возвращает массив пар ключ-значение: [ ["name","John"], ["age", 30] ]
let peopleMap = new Map(Object.entries(obj));
console.log(peopleMap.get('name') ); // John


// Object.fromEntries: Object из Map
let fruitsMap = new Map();
map.set('banana', 1);
map.set('orange', 2);
map.set('meat', 4);
let fruits = Object.fromEntries(map.entries());
// готово!
// obj = { banana: 1, orange: 2, meat: 4 }
console.log(fruits.orange); // 2

/*
Map – коллекция пар ключ-значение.

Методы и свойства:

new Map([iterable]) – создаёт коллекцию, можно указать перебираемый объект (обычно массив) из пар [ключ,значение] для инициализации.
map.set(key, value) – записывает по ключу key значение value.
map.get(key) – возвращает значение по ключу или undefined, если ключ key отсутствует.
map.has(key) – возвращает true, если ключ key присутствует в коллекции, иначе false.
map.delete(key) – удаляет элемент по ключу key.
map.clear() – очищает коллекцию от всех элементов.
map.size – возвращает текущее количество элементов.

Отличия от обычного объекта Object:

Что угодно может быть ключом, в том числе и объекты.
Есть дополнительные методы, свойство size.
 */