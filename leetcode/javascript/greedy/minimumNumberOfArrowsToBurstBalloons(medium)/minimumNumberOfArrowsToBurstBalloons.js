/* https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/description/

К плоской стене, представляющей плоскость XY, приклеено несколько сферических воздушных шаров. Выноски представлены в
виде двумерного целочисленного массива точек, где points[i] = [xstart, xend] обозначает выноску, горизонтальный диаметр
которой простирается между xstart и xend. Вы не знаете точных координат Y воздушных шаров.

Стрелки можно стрелять прямо вертикально (в положительном направлении Y) из разных точек вдоль оси X. Воздушный шар с
xstart и xend разрывается стрелой, выпущенной в x, если xstart <= x <= xend. Нет ограничений на количество выпущенных
стрел. Выпущенная стрела продолжает двигаться вверх бесконечно, разрывая все воздушные шары на своем пути.

Учитывая точки массива, верните минимальное количество стрел, которое необходимо выпустить, чтобы лопнуть все
воздушные шары.

Input: points = [[10,16],[2,8],[1,6],[7,12]]
Output: 2
Пояснение: Воздушные шары можно лопнуть двумя стрелами:
- Выстрелите стрелой в точку x = 6, лопнув шарики [2,8] и [1,6].
- Выстрелите стрелой в точку x = 11, лопнув шарики [10,16] и [7,12].
*/

/**
 * @param {number[][]} points
 * @return {number}
 */
var findMinArrowShots = function(points) {
    if (points.length === 0) return -1;
    points.sort((a, b) => a[1] - b[1]);

    let result = [points[0]];
    let arrows = 1; // 1 стрела обязательно понадобится для взрыва шариков

    for (let i = 1; i < points.length; i++) {
        let prevLastValue = result[result.length - 1];
        let currFirstValue = points[i];

        if (prevLastValue[1] < currFirstValue[0]) {
            arrows++;
            result.push(currFirstValue);
        }
    }

    return arrows;
};

const points = [[10,16],[2,8],[1,6],[7,12]];
console.log(findMinArrowShots(points)); // 2