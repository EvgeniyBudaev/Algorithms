const graph = {};
graph.you = ["alice", "bob", "claire"];
graph.bob = ["anuj", "peggy"];
graph.alice = ["peggy"];
graph.claire = ["thom", "jonny"];
graph.anuj = [];
graph.peggy = [];
graph.thom = [];
graph.jonny = [];

/**
 * Determine whether a person is a seller
 * @param {string} name Friend's name
 * @returns {boolean} Result of checking
 */
const personIsSeller = (name) => name[name.length - 1] === "m";

/**
 * Find a mango seller
 * @param {string} name Friend's name
 * @returns {boolean} Search results
 */
const search = (name) => {
    let searchQueue = [...graph[name]]; // Создание новой очереди и добавление элементов в очередь поиска
    const searched = []; // С помощью этого массива вы можете отслеживать, каких людей вы искали раньше

    while (searchQueue.length) { // Пока очередь не пуста
        const person = searchQueue.shift();
        if (searched.indexOf(person) === -1) { // Если элемент в очереди не был проверен раньше
            if (personIsSeller(person)) {
                console.log(`${person} is a mango seller!`);
                return true;
            }
            searchQueue = searchQueue.concat(graph[person]);
            searched.push(person);
        }
    }

    return false;
}

search("you"); // thom is a mango seller!
