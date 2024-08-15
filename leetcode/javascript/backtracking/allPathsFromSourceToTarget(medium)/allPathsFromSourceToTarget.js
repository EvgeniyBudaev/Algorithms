/* https://leetcode.com/problems/all-paths-from-source-to-target/description/

Дан направленный ациклический граф (DAG) из n узлов, помеченных от 0 до n-1, найдите все возможные пути от узла 0 до
узла n-1 и верните их в любом порядке.

Граф задается следующим образом: граф[i] — это список всех узлов, которые вы можете посетить из узла i
(т. е. существует направленное ребро от узла i к узлу граф[i][j]).

Input: graph = [[1,2],[3],[3],[]]
Output: [[0,1,3],[0,2,3]]
Пояснение: Есть два пути: 0 -> 1 -> 3 и 0 -> 2 -> 3.

Input: graph = [[4,3,1],[3,2,4],[3],[4],[]]
Output: [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
*/

/**
 * @param {number[][]} graph
 * @return {number[][]}
 */
var allPathsSourceTarget = function(graph) {
    const result = [];
    const dfs = function(node, path){
        if (node === graph.length - 1) result.push(path);
        for (let neighbour of graph[node]) {
            dfs(neighbour, path.concat([neighbour]));
        }
    }
    dfs(0, [0]);
    return result;
};

const graph = [[1,2],[3],[3],[]];
console.log(allPathsSourceTarget(graph)); // [[0,1,3],[0,2,3]]