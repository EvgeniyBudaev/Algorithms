/*
У вас есть палочки с целыми положительными длинами.
Вы можете соединить любые две палочки длиной X и Y в одну, заплатив стоимость X + Y.
Вы выполняете это действие до тех пор, пока не останется одна палочка.
Вернуть таким образом минимальную стоимость соединения всех данных палочек в одну.
*/

/*
Input: sticks = [2,4,3]
Output: 14
 */

/**
 * @param {number[]} sticks
 * @return {number}
 */
var connectSticks = function(sticks) {
    const pq = new Heap(sticks);
    let ans = 0;
    while (pq.size() > 1) {
        const x = pq.pop();
        const y = pq.pop();
        ans += x + y;
        pq.push(x + y);
    }
    return ans;
};

class Heap {
    data = [null]; // Инициализируйте фиктивным элементом с индексом 0.
    lt = (i, j) => this.compare(this.data[i], this.data[j]) < 0;

    constructor(data = []) {
        this.data = [null,...data];
        this.compare = (lhs, rhs) => (lhs < rhs? -1 : lhs > rhs? 1 : 0);
        for (let i = this.size(); i > 0; i--) this.heapify(i);
    }

    size() {
        return this.data.length - 1;
    }

    push(v) {
        this.data.push(v);
        let i = this.size();
        while (i >> 1!== 0 && this.lt(i, Math.floor(i / 2))) this.swap(i, Math.floor(i / 2));
    }

    pop() {
        this.swap(1, this.size());
        const top = this.data.pop();
        this.heapify(1);
        return top;
    }

    top() {
        return this.data[1];
    }

    heapify(i) {
        while (true) {
            let min = i;
            const [l, r, n] = [i * 2, i * 2 + 1, this.data.length];
            if (l < n && this.lt(l, min)) min = l;
            if (r < n && this.lt(r, min)) min = r;
            if (min!== i) {
                this.swap(i, min);
                i = min;
            } else break;
        }
    }

    clear() {
        this.data = [null];
    }

    swap(i, j) {
        const d = this.data;
        [d[i], d[j]] = [d[j], d[i]];
    }
}

console.log(connectSticks([2,4,3])); // 14