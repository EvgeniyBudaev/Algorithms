/*
Вам дан целочисленный массив с индексом 0, где Piles[i] представляет количество камней в i-й куче, и целое число k.
Вам следует применить следующую операцию ровно k раз:
Выберите любую груду[i] и удалите из нее камни пола(piles[i]/2).
Обратите внимание, что вы можете применить операцию к одной и той же стопке более одного раза.
Возвращает минимально возможное общее количество камней, оставшихся после выполнения k операций.
Floor(x) — наибольшее целое число, которое меньше или равно x (т. е. округляет x в меньшую сторону).
*/

/*
Input: piles = [5,4,9], k = 2
Output: 12
Пояснение: Шаги возможного сценария:
- Примените операцию к стопке 2. Полученные стопки — [5,4,5].
- Примените операцию к стопке 0. Полученные стопки — [3,4,5].
Общее количество камней в [3,4,5] — 12.
 */

/**
 * @param {number[]} piles
 * @param {number} k
 * @return {number}
 */
var minStoneSum = function(piles, k) {
    const max = new MaxHeap(piles);

    while (k > 0) {
        const remove = max.remove();
        const divide = Math.floor(remove / 2);
        max.insert(remove - divide);
        k--;
    }

    return piles.reduce((a, b) => a + b);
};

class MaxHeap {
    constructor(array) {
        this.heap = this.buildHeap(array);
    }

    buildHeap(array) {
        const firstParentIdx = Math.floor((array.length - 2) / 2);
        for (let currentIdx = firstParentIdx; currentIdx >= 0; currentIdx--) {
            this.siftDown(currentIdx, array.length - 1, array);
        }
        return array;
    }

    siftDown(currentIdx, endIdx, heap) {
        let childOneIdx = currentIdx * 2 + 1;

        while (childOneIdx <= endIdx) {
            const childTwoIdx = currentIdx * 2 + 2 <= endIdx ? currentIdx * 2 + 2 : -1;

            let idxToSwap;

            if (childTwoIdx !== -1 && heap[childTwoIdx] > heap[childOneIdx]) {
                idxToSwap = childTwoIdx;
            } else {
                idxToSwap = childOneIdx;
            }

            if (heap[idxToSwap] > heap[currentIdx]) {
                this.swap(currentIdx, idxToSwap, heap);

                currentIdx = idxToSwap;
                childOneIdx = currentIdx * 2 + 1;
            } else {
                return;
            }
        }
    }

    siftUp(currentIdx, heap) {
        let parentIdx = Math.floor((currentIdx - 1) / 2);

        while (currentIdx > 0 && heap[currentIdx] > heap[parentIdx]) {
            this.swap(currentIdx, parentIdx, heap);
            currentIdx = parentIdx;
            parentIdx = Math.floor((currentIdx - 1) / 2);
        }
    }

    peek() {
        return this.heap[0];
    }

    remove() {
        this.swap(0, this.heap.length - 1, this.heap);
        const valueToRemove = this.heap.pop();

        this.siftDown(0, this.heap.length - 1, this.heap);
        return valueToRemove;
    }

    insert(value) {
        this.heap.push(value);
        this.siftUp(this.heap.length - 1, this.heap);
    }

    swap(i, j, heap) {
        const temp = heap[j];
        heap[j] = heap[i];
        heap[i] = temp;
    }
}

console.log(minStoneSum([5,4,9], 2)); // 12