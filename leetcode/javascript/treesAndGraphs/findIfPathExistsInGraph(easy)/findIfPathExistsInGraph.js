/*
Существует двунаправленный граф с n вершинами, где каждая вершина помечена от 0 до n – 1 (включительно).
Ребра в графе представлены как ребра двумерного целочисленного массива, где каждое ребро [i] = [ui, vi] обозначает
двунаправленное ребро между вершиной ui и вершиной vi. Каждая пара вершин соединена не более чем одним ребром,
и ни одна вершина не имеет ребра сама по себе.
Вы хотите определить, существует ли допустимый путь от источника вершины до места назначения вершины.
Учитывая ребра и целые числа n, источник и пункт назначения, возвращайте true, если существует действительный путь от
источника к месту назначения, или false в противном случае.
*/

/*
Input: n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
Output: true
Explanation: There are two paths from vertex 0 to vertex 2:
- 0 → 1 → 2
- 0 → 2
 */

/*
Input: n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], source = 0, destination = 5
Output: false
Explanation: There is no path from vertex 0 to vertex 5.
 */

/**
 * @param {number} n
 * @param {number[][]} edges
 * @param {number} source
 * @param {number} destination
 * @return {boolean}
 */
var validPath = function(n, edges, source, destination) {
    let g = new Array(n);
    for (let i = 0; i < n; i++) g[i] = [];
    for (let i of edges) {
        g[i[0]].push(i[1]);
        g[i[1]].push(i[0]);
    }
    let vis = new Array(n).fill(0);
    rec(source,g,vis);
    return vis[destination];
};

var rec = (node, g, vis) => {
    vis[node] = 1;
    for (let i of g[node]) {
        if (!vis[i]) rec(i, g, vis);
    }
};

const example1 = [[0,1],[1,2],[2,0]];
const tree1 = buildTree(example1);
console.log(validPath(3, tree1, 0, 2)); // true