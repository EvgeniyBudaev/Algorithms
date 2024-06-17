/*
Вам дан список бомб. Дальность действия бомбы определяется как область, где можно ощутить ее эффект.
Эта область имеет форму круга, в центре которого находится бомба.

Бомбы представлены двумерным целочисленным массивом с нулевым индексом, где бомбы[i] = [xi, yi, ri]. xi и yi обозначают
координату X и координату Y местоположения i-й бомбы, тогда как ri обозначает радиус ее дальности.

Вы можете взорвать одну бомбу. Когда бомба взорвана, она взорвет все бомбы, находящиеся в радиусе ее действия.
Эти бомбы в дальнейшем взорвут бомбы, находящиеся в зоне их действия.

Учитывая список бомб, верните максимальное количество бомб, которые можно взорвать, если вам разрешено взорвать только
одну бомбу.
*/

/*
Input: bombs = [[2,1,3],[6,1,4]]
Output: 2
Explanation:
The above figure shows the positions and ranges of the 2 bombs.
If we detonate the left bomb, the right bomb will not be affected.
But if we detonate the right bomb, both bombs will be detonated.
So the maximum bombs that can be detonated is max(1, 2) = 2.
 */

/*
Input: bombs = [[1,1,5],[10,10,5]]
Output: 1
Explanation:
Detonating either bomb will not detonate the other bomb, so the maximum number of bombs that can be detonated is 1.
 */

/**
 * @param {number[][]} bombs
 * @return {number}
 */
var maximumDetonation = function(bombs) {
    const n = {};

    for (let i = 0 ; i < bombs.length ; i++) {
        const x1 = bombs[i][0];
        const y1 = bombs[i][1];
        const r = bombs[i][2];

        for (let j = 0; j < bombs.length ; j++) {
            const x2 = bombs[j][0];
            const y2 = bombs[j][1];

            if (Math.pow((x1-x2), 2) + Math.pow((y1-y2), 2) <= Math.pow(r, 2)) {
                n[i] = (n[i] || []);
                n[i].push(j);
            }
        }
    }

    let max = 0;

    for (let i = 0 ; i < bombs.length ; i++) {
        const det = new Set([i]);
        const q = [...n[i]];

        while (q.length) {
            const curr = q.shift();

            if (det.has(curr)) {
                continue;
            }

            det.add(curr);
            q.push(...n[curr]);
        }

        if (det.size === bombs.length) {
            return bombs.length;
        }

        max = Math.max(max, det.size);
    }

    return max;
};

const bombs = [[2,1,3],[6,1,4]];
console.log(maximumDetonation(bombs)); // 2