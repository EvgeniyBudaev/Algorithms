/*
Существует неориентированное дерево с n узлами, помеченными от 0 до n - 1 и n - 1 ребром.
Вам дан двумерный целочисленный массив ребер длины n - 1, где ребра [i] = [ai, bi] указывают на то, что между узлами ai
и bi в дереве есть ребро. Вам также предоставляется ограниченный целочисленный массив, который представляет
ограниченные узлы.
Возвращает максимальное количество узлов, которых вы можете достичь из узла 0, не посещая узел с ограниченным доступом.
Обратите внимание, что узел 0 не будет ограниченным узлом.
*/

/*
Input: n = 7, edges = [[0,1],[1,2],[3,1],[4,0],[0,5],[5,6]], restricted = [4,5]
Output: 4
Explanation: The diagram above shows the tree.
We have that [0,1,2,3] are the only nodes that can be reached from node 0 without visiting a restricted node.
 */

/*
Input: n = 7, edges = [[0,1],[0,2],[0,5],[0,4],[3,2],[6,5]], restricted = [4,2,1]
Output: 3
Explanation: The diagram above shows the tree.
We have that [0,5,6] are the only nodes that can be reached from node 0 without visiting a restricted node.
 */

/**
 * @param {number} n
 * @param {number[][]} edges
 * @param {number[]} restricted
 * @return {number}
 */
var reachableNodes = function(n, edges, restricted) {
    // алгоритм поиска в ширину (BFS - Breadth-First Search)

    // Создается массив root размером n+1, где каждый элемент является пустым массивом. Этот массив будет использоваться
    // для представления графа в виде списка смежности. Узлы индексируются с 0 до n.
    const root = new Array(n + 1).fill().map(() => []);

    // Затем, для каждой пары вершин (v1, v2) из входного массива edges, добавляются ребра между ними в обе стороны
    // (т.е., если есть ребро от v1 к v2, то оно также существует от v2 к v1)
    for (const [v1, v2] of edges) {
        root[v1].push(v2);
        root[v2].push(v1);
    }

    // Создается множество visited для отслеживания уже посещенных узлов
    const vizited = new Set();
    // Массив res используется для хранения узлов, которые будут посещены в процессе поиска
    const res = [];

    // Входной массив restricted содержит узлы, которые не могут быть посещены. Эти узлы добавляются в множество visited
    for (const item of restricted) {
        vizited.add(item);
    }

    // Поиск в ширину
    // Инициализируется очередь queue с начальным узлом (обычно это 0 или любой другой указанный корневой узел)
    const queue = [0];
    while (queue.length) { // Пока очередь не пуста, выполняются следующие действия
        // Извлекается первый узел из очереди (node), добавляется в результат (res) и помечается как посещенный visited
        const node = queue.shift();
        res.push(node);
        vizited.add(node);

        for (const vertex of root[node]) {
            // Для каждого соседа текущего узла, который еще не был посещен (visited.has(vertex)),
            // он добавляется в очередь
            if (!vizited.has(vertex)) {
                queue.push(vertex);
            }
        }
    }

    // После завершения поиска в ширину, функция возвращает количество узлов, доступных из корневого узла, что
    // соответствует длине массива res
    return res.length;
};

// var reachableNodes = function(n, edges, restricted) {
//     const adjList = buildAdjList(n, edges);
//     const visited = {};
//     let max = 0;
//
//     for(let item of restricted) { visited[item] = true; }
//
//     const dfs = (node, adjList, visited) => {
//         visited[node] = true;
//         max++;
//
//         for(let neighbor of adjList[node]) {
//             if(!visited[neighbor]) {
//                 dfs(neighbor, adjList, visited);
//             }
//         }
//     }
//
//     dfs(0, adjList, visited);
//
//     return max;
// };
//
// const buildAdjList = (n, edges) => {
//     const adjList = Array.from({length:n}, () => []);
//
//     for(let edge of edges) {
//         let [src, dest] = edge;
//         adjList[src].push(dest);
//         adjList[dest].push(src);
//     }
//
//     return adjList;
// }

const edges = [[0,1],[1,2],[3,1],[4,0],[0,5],[5,6]];
const restricted = [4,5];
console.log(reachableNodes(7, edges, restricted)); // 4