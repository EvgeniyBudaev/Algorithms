// Объект Set – это особый вид коллекции: «множество» значений (без ключей), где каждое значение может появляться только
// один раз.

let set = new Set();

let john = { name: "John" };
let pete = { name: "Pete" };
let mary = { name: "Mary" };

// считаем гостей, некоторые приходят несколько раз
set.add(john);
set.add(pete);
set.add(mary);
set.add(john);
set.add(mary);

// set хранит только 3 уникальных значения
console.log(set.size); // 3

for (let user of set) {
    console.log(user.name); // John (потом Pete и Mary)
}


// Перебор объекта Set
let fruitsSet = new Set(["апельсин", "яблоко", "банан"]);

for (let value of fruitsSet) console.log(value);

// то же самое с forEach:
fruitsSet.forEach((value, valueAgain, set) => {
    console.log(value);
});

/*
Set имеет те же встроенные методы, что и Map:

set.values() – возвращает перебираемый объект для значений,
set.keys() – то же самое, что и set.values(), присутствует для обратной совместимости с Map,
set.entries() – возвращает перебираемый объект для пар вида [значение, значение], присутствует для обратной
 совместимости с Map.
 */

/*
Set – коллекция уникальных значений, так называемое «множество».

Методы и свойства:

new Set(iterable) – создаёт Set, можно указать перебираемый объект со значениями для инициализации.
set.add(value) – добавляет значение (если оно уже есть, то ничего не делает), возвращает тот же объект set.
set.delete(value) – удаляет значение, возвращает true если value было в множестве на момент вызова, иначе false.
set.has(value) – возвращает true, если значение присутствует в множестве, иначе false.
set.clear() – удаляет все имеющиеся значения.
set.size – возвращает количество элементов в множестве.
 */