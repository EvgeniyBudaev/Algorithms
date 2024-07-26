/* https://leetcode.com/problems/course-schedule-ii/description/

Всего вам необходимо пройти numCourses курсов, помеченных от 0 до numCourses - 1. Вам предоставляется массив
предварительных требований, где prequires[i] = [ai, bi] указывает, что вы должны сначала пройти курс bi, если хотите пройти конечно ай.

Например, пара [0, 1] указывает, что для прохождения курса 0 необходимо сначала пройти курс 1.
Верните порядок курсов, которые вам необходимо пройти, чтобы закончить все курсы. Если правильных ответов много, верните
любой из них. Если пройти все курсы невозможно, верните пустой массив.

Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Объяснение: Всего нужно пройти 2 курса. Чтобы пройти курс 1, вы должны закончить курс 0. Поэтому правильный порядок
курсов — [0,1].

Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Пояснение: Всего необходимо пройти 4 курса. Чтобы пройти курс 3, вам необходимо пройти оба курса 1 и 2. Оба курса 1 и 2
следует пройти после окончания курса 0. Таким образом, один правильный порядок курса — [0,1,2,3].
Другой правильный порядок — [0,2,1,3]

Input: numCourses = 1, prerequisites = []
Output: [0]
*/

/**
 * @param {number} numCourses
 * @param {number[][]} prerequisites
 * @return {number[]}
 */
var findOrder = function(numCourses, prerequisites) {
    let indegrees = new Array(numCourses).fill(0);
    let adj = new Array(numCourses).fill(0).map(() => new Array());

    for (let [a, b] of prerequisites) {
        adj[b].push(a);
        indegrees[a]++;
    }

    let queue = [];

    for (let i = 0; i < numCourses; i++) {
        if (indegrees[i] === 0) {
            queue.push(i);
        }
    }

    let reply = [];

    while (queue.length) {
        let currCourseNum = queue.shift();
        reply.push(currCourseNum);
        for (let nextCourseNum of adj[currCourseNum]) {
            indegrees[nextCourseNum]--;
            if (indegrees[nextCourseNum] === 0) {
                queue.push(nextCourseNum);
            }
        }
    }

    return reply.length === numCourses ? reply : [];
};

console.log(findOrder(2, [[1,0]])); // [0,1]