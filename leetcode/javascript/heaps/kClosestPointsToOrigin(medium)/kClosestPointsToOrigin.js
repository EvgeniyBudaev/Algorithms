// https://leetcode.com/problems/k-closest-points-to-origin/solutions/762781/javascript-sort-minheap-and-maxheap-solutions/

/*
Учитывая массив точек, где Points[i] = [xi, yi] представляет точку на плоскости XY и целое число k,
верните k ближайших точек к началу координат (0, 0).
Расстояние между двумя точками на плоскости XY — это евклидово расстояние (т. е. √(x1 — x2)2 + (y1 — y2)2).
Вы можете вернуть ответ в любом порядке. Ответ гарантированно будет уникальным (за исключением порядка,
в котором он находится).
*/

/*
Input: points = [[1,3],[-2,2]], k = 1
Output: [[-2,2]]
Объяснение:
Расстояние между (1, 3) и началом координат равно sqrt(10).
Расстояние между (-2, 2) и началом координат равно sqrt(8).
Поскольку sqrt(8) < sqrt(10), (-2, 2) ближе к началу координат.
Нам нужны только ближайшие k = 1 точки от начала координат, поэтому ответ — просто [[-2,2]].
 */

/**
 * @param {number[][]} points
 * @param {number} k
 * @return {number[][]}
 */
/*
min heap approach:
мы можем создать minHeap всего набора данных за время O(n), если начнем с n/2 и составим кучу каждого родителя
(см. метод Floyd https://en.wikipedia.org/wiki/Binary_heap#Building_a_heap)
затем мы удаляем k раз из кучи -> k * log(n) (необходимо складывать в кучу при каждом удалении)
runtime: O(N + k log (N))
space: O(1) since we are doing it in place
*/
var kClosest = function(points, k) {
    // мы можем построить кучу на месте
    let p = Math.floor((points.length - 2) / 2) // last parent
    for(let i = p; i >= 0; i--) {
        heapifyDown(points, i, distance)
    }

    // теперь нам нужно удалить наименьшее (points[0]) k раз
    let solution = []
    for(let i=0; i<k; i++) {
        solution.push(remove(points, distance))
    }

    return solution

    // прочитать 0, заменить 0 последней позицией, heapifyDown
    function remove(heap, weightFunction) {
        let val = heap[0]
        heap[0] = heap.pop()
        heapifyDown(heap, 0, weightFunction)
        return val
    }

    // сравнить с детьми, поменять местами с самыми маленькими, повторить
    function heapifyDown(heap, idx, weightFunction) {
        let left = 2*idx + 1
        let right = 2*idx + 2
        let smallest = left

        if(left >= heap.length) return

        if(right < heap.length && weightFunction(heap[left]) > weightFunction(heap[right])) {
            smallest = right
        }

        if (weightFunction(heap[idx]) > weightFunction(heap[smallest])) {
            [heap[idx], heap[smallest]] = [heap[smallest], heap[idx]]
            heapifyDown(heap, smallest, weightFunction)
        }
    }

    function distance(point) {
        return point[0] * point[0] + point[1] * point[1]
    }
}


/*
max heap approach:
иметь максимальную кучу размера k, поэтому мы должны сделать N вставок, которые принимают log(k)
в этом случае нам нужно будет реализовать heapify up (вставить) и heapify down (удалить)

runtime: O(N log(k))
space: O(k)
*/
var kClosest = function(points, k) {
    let heap = []

    // теперь нам нужно попытаться сложить все точки в кучу
    for(let i=0; i<points.length; i++) {
        if(heap.length >= k && distance(points[i]) > distance(heap[0])) { // it's bigger than the max, we can just skip it
            continue
        }
        add(heap, points[i], distance)
        if(heap.length > k) {
            remove(heap, distance)
        }
    }

    return heap

    // добавить в конце, собрать в кучу
    function add(heap, node, weightFunction) {
        heap.push(node)
        heapifyUp(heap, heap.length - 1, weightFunction)
    }

    // сравните с родительским и при необходимости поменяйте местами, повторите
    function heapifyUp(heap, idx, weightFunction) {
        if(idx === 0) return
        let parent = Math.floor((idx-1) / 2)
        if(weightFunction(heap[idx]) > weightFunction(heap[parent])) {
            [heap[idx], heap[parent]] = [heap[parent], heap[idx]]
            heapifyUp(heap, parent, weightFunction)
        }
    }

    // прочитать 0, заменить 0 последней позицией, heapifyDown
    function remove(heap, weightFunction) {
        let val = heap[0]
        heap[0] = heap.pop()
        heapifyDown(heap, 0, weightFunction)
        return val
    }

    // сравнить с детьми, поменять местами с самым большим, повторить
    function heapifyDown(heap, idx, weightFunction) {
        let left = 2*idx + 1
        let right = 2*idx + 2
        let biggest = left

        if(left >= heap.length) return

        if(right < heap.length && weightFunction(heap[left]) < weightFunction(heap[right])) {
            biggest = right
        }

        if (weightFunction(heap[idx]) < weightFunction(heap[biggest])) {
            [heap[idx], heap[biggest]] = [heap[biggest], heap[idx]]
            heapifyDown(heap, biggest, weightFunction)
        }
    }

    function distance(point) {
        return point[0] * point[0] + point[1] * point[1]
    }
}

const points = [[1,3],[-2,2]];
console.log(kClosest(points, 1)); // [[-2,2]]
