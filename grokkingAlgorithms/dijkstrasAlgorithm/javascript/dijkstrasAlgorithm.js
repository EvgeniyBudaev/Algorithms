// the graph
const graph = {};
graph.start = {};
graph.start.a = 6;
graph.start.b = 2;

graph.a = {};
graph.a.fin = 1;

graph.b = {};
graph.b.a = 3;
graph.b.fin = 5;

graph.fin = {}; // У конечного узла нет соседей

// The costs table
const costs = {};
costs.a = 6;
costs.b = 2;
costs.fin = Infinity;

// the parents table
const parents = {};
parents.a = "start";
parents.b = "start";
parents.fin = null;

let processed = []; // Массив уже обработанных узлов

/**
 * Find the lowest node
 * @param {Object} itCosts Hash table
 * @returns {(string|null)} The lowest node
 */
const findLowestCostNode = itCosts => {
    let lowestCost = Infinity;
    let lowestCostNode = null;

    Object.keys(itCosts).forEach(node => {
        const cost = itCosts[node];
        // Если это самая низкая стоимость на данный момент и она еще не обработана...
        if (cost < lowestCost && processed.indexOf(node) === -1) {
            // ... установите его как новый узел с наименьшей стоимостью.
            lowestCost = cost;
            lowestCostNode = node;
        }
    });
    return lowestCostNode;
};

let node = findLowestCostNode(costs);

while (node !== null) {
    const cost = costs[node];
    // Перебрать всех соседей этого узла
    const neighbors = graph[node];
    Object.keys(neighbors).forEach(n => {
        const newCost = cost + neighbors[n];
        // Если дешевле добраться до этого соседа, пройдя через этот узел
        if (costs[n] > newCost) {
            // ... обновить стоимость для этого узла
            costs[n] = newCost;
            // Этот узел становится новым родителем для этого соседа
            parents[n] = node;
        }
    });

    // Отметить узел как обработанный
    processed = processed.concat(node);

    // Найдите следующий узел для обработки и выполните цикл
    node = findLowestCostNode(costs);
}

// Код для восстановления пути
function reconstructPath(parents, start, end) {
    const path = [end];
    while (path[0] !== start) {
        path.unshift(parents[path[0]]);
    }
    return path;
}
console.log("parents", parents);
const pathHistory = reconstructPath(parents, 'start', 'fin');

console.log("Cost from the start to each node:");
console.log(costs); // { a: 5, b: 2, fin: 6 }
console.log(pathHistory); // ['start', 'b', 'a', 'fin']
