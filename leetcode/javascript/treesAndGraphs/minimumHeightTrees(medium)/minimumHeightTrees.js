/* https://leetcode.com/problems/minimum-height-trees/description/

Дерево — это неориентированный граф, в котором любые две вершины соединены ровно одним путем. Другими словами, любой
связный граф без простых циклов является деревом.

Учитывая дерево из n узлов, помеченных от 0 до n - 1, и массив из n - 1 ребер, где ребра [i] = [ai, bi] указывают, что
между двумя узлами ai и bi в дереве существует ненаправленное ребро. , вы можете выбрать любой узел дерева в качестве
корня. Когда вы выбираете узел x в качестве корня, результирующее дерево имеет высоту h. Среди всех возможных корневых
деревьев деревья с минимальной высотой (т.е. min(h)) называются деревьями минимальной высоты (MHT).

Возвращает список корневых меток всех MHT. Вы можете вернуть ответ в любом порядке.
Высота дерева с корнями — это количество ребер на самом длинном нисходящем пути между корнем и листом.

Input: n = 4, edges = [[1,0],[1,2],[1,3]]
Output: [1]
Объяснение: Как показано, высота дерева равна 1, когда корнем является узел с меткой 1, который является единственным MHT.
*/

/**
 * @param {number} n
 * @param {number[][]} edges
 * @return {number[]}
 */
var findMinHeightTrees = function(n, edges) {
    if (!edges.length) return [0];
    const depth = Array(n).fill(0);
    const graph = Array.from({ length: n }, () => []);

    for (const [parent, child] of edges) {
        graph[parent].push(child);
        graph[child].push(parent);
        depth[child]++;
        depth[parent]++;
    }

    const queue = [];
    let front = 0;
    for (let i = 0; i < depth.length; i++) {
        if (depth[i] === 1) queue.push(i);
    }
    while (n > 2) {
        const popEle = queue.length - front;
        n -= popEle;
        for (let i = 0; i < popEle; i++) {
            const currentElm = queue[front++];
            for (const currentGraph of graph[currentElm]) {
                depth[currentGraph]--;
                if (depth[currentGraph] === 1) queue.push(currentGraph);
            }
        }
    }
    return queue.slice(front);
};

console.log(findMinHeightTrees(4, [[1,0],[1,2],[1,3]])); // [1]