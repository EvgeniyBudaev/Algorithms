// https://www.youtube.com/watch?v=8WCfY0Uh1iM&ab_channel=%D0%9E%D0%B1%D1%83%D1%87%D0%B5%D0%BD%D0%B8%D0%B5HTML%2CCSS%2CJavaScript%7C%D0%90%D0%BD%D0%B0%D1%82%D0%BE%D0%BB%D0%B8%D0%B9%D0%98%D0%B2%D0%B0%D1%88%D0%BE%D0%B2
// Задача: нужно вывести кратчайший путь из точки А в точку Б

// Взвешенный граф (указаны расстояния между вершинами)
const graph = {
    a: { b: 2, c: 1, i: 3 },
    b: { a: 2, d: 3 },
    c: { a: 1, d: 1 },
    d: { b: 3, c: 1, e: 5 },
    e: { d: 5, i: 2 },
    i: { a: 3, e: 2 },
};

const findLowestCostNode = (distances, visited) => {
    let lowestCost = Infinity;
    let lowestCostNode = null;
    for (const key in distances) {
        if (lowestCost > distances[key] && !visited.has(key)) {
            lowestCost = distances[key];
            lowestCostNode = key;
        }
    }
    return lowestCostNode;
};

function dijkstra(graph, start, end) {
    const distances = {}; // Список расстояний
    const visited = new Set(); // Список посещенных вершин
    const path = {}; // История пути

    for (const key in graph) {
        if (key !== start) {
            distances[key] = Infinity;
        } else {
            distances[start] = 0;
        }
    }

    while (!visited.has(end)) {
        const node = findLowestCostNode(distances, visited);
        const neighbors = graph[node];
        for (const key in neighbors) {
            const newDistance = distances[node] + neighbors[key];
            if (newDistance < distances[key]) {
                distances[key] = newDistance;
                path[key] = node;
            }
        }

        visited.add(node);
    }

    const shortPath = [];
    let current = end;

    while (current !== start)  {
        shortPath.unshift(current);
        current = path[current];
    }

    shortPath.unshift(start);
    return shortPath;
}

console.log(dijkstra(graph, 'a', 'e')); // ['a', 'i', 'e']