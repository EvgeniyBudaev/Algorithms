/* https://baihuqian.github.io/2018-08-12-line-reflection/

Учитывая n точек на двумерной плоскости, найдите, существует ли такая линия, параллельная оси Y, которая отражает данные
точки.

Input: [[1,1],[-1,1]]
Output: true

Input: [[1,1],[-1,-1]]
Output: false
*/

function isReflected(points) {
    let minX = Infinity, maxX = -Infinity;
    const pointsSet = new Set();

    // Находим минимальное и максимальное значения X и добавляем точки в set
    for (const [x, y] of points) {
        minX = Math.min(x, minX);
        maxX = Math.max(x, maxX);
        pointsSet.add(`${x},${y}`);
    }

    const middle = minX + maxX;

    // Проверяем, есть ли отраженная точка для каждой исходной точки
    for (const [x, y] of points) {
        const reflectedX = middle - x;
        if (!pointsSet.has(`${reflectedX},${y}`)) return false;
    }

    return true;
}

console.log(isReflected([[1,1],[-1,1]])); // true
console.log(isReflected([[1,1],[-1,-1]])); // false