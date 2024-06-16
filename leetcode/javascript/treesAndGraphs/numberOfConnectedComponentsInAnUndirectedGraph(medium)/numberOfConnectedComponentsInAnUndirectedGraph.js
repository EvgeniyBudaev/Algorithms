/*
Учитывая n узлов, помеченных от 0 до n - 1, и список неориентированных ребер (каждое ребро представляет собой пару
узлов), напишите функцию для определения количества связных компонентов в неориентированном графе.
*/

/*
Input: n = 5 and edges = [[0, 1], [1, 2], [3, 4]]

     0          3
     |          |
     1 --- 2    4

Output: 2
 */

/*
Input: n = 5 and edges = [[0, 1], [1, 2], [2, 3], [3, 4]]

     0           4
     |           |
     1 --- 2 --- 3

Output:  1
 */

/**
 * @param {number} n
 * @param {number[][]} edges
 * @return {number}
 */
var countComponents = function (n, edges) {
    let count = 0;
    let graph = {};

    for (let i = 0; i < n; i++) {
        graph[i] = [];
    }

    for (let [k, v] of edges) {
        graph[k].push(v);
        graph[v].push(k);
    }

    let visited = new Set();

    function  dfs(node) {
        if (visited.has(node)) return  0;
        visited.add(node);
        for (let n of graph[node]) {
            dfs(n);
        }
        return  1;
    }

    for (let key in graph) {
        key = parseInt(key);
        count  +=  dfs(key);
    }

    return  count;
}

console.log(countComponents(5, [[0, 1], [1, 2], [3, 4]])); // 2