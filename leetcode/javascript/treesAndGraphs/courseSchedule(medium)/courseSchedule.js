/* https://leetcode.com/problems/course-schedule/

Всего вам необходимо пройти numCourses курсов, помеченных от 0 до numCourses - 1. Вам предоставляется массив
предварительных требований, где prequires[i] = [ai, bi] указывает, что вы должны сначала пройти курс bi, если хотите
пройти конечно ай.
Например, пара [0, 1] указывает, что для прохождения курса 0 необходимо сначала пройти курс 1.
Верните true, если вы можете закончить все курсы. В противном случае верните false.

Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Объяснение: Всего нужно пройти 2 курса.
Чтобы пройти курс 1, вам необходимо закончить курс 0. Так что это возможно.

Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Объяснение: Всего нужно пройти 2 курса.
Чтобы пройти курс 1, вы должны закончить курс 0, а чтобы пройти курс 0, вы также должны закончить курс 1.
Так что это невозможно.
*/

/**
 * @param {number} numCourses
 * @param {number[][]} prerequisites
 * @return {boolean}
 */
var canFinish = function(numCourses, prerequisites) {
    const preReqGraph = prerequisites.reduce((graph, [course, preReq]) => {
        graph[course].push(preReq)
        return graph
    }, new Array(numCourses).fill(0).map(() => []));

    const canFinishCourse = new Array(numCourses).fill(false);
    const visited = new Array(numCourses).fill(false);

    const dfs = (course) => {
        if (visited[course]) return canFinishCourse[course];

        visited[course] = true;
        for (const preReq of preReqGraph[course]) {
            const result = dfs(preReq);
            canFinishCourse[preReq] = result;
            if (!result) return false;
        }
        return true;
    }

    for (let i = 0; i < numCourses; i++) {
        if (canFinishCourse[i]) continue;
        canFinishCourse[i] = dfs(i);
        if (!canFinishCourse[i]) return false;
    }

    return true;
};

console.log(canFinish(2, [[1,0]])); // true

