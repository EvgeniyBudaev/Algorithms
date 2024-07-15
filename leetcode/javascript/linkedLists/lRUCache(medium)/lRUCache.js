/* https://leetcode.com/problems/lru-cache/description/

Спроектируйте структуру данных, которая соответствует ограничениям кэша наименее недавно используемого (LRU).

Реализуйте класс LRUCache:

LRUCache(int емкость) Инициализировать кэш LRU с емкостью положительного размера.
int get(int key) Возвращает значение ключа, если ключ существует, в противном случае возвращает -1.
void put(int key, int value) Обновить значение ключа, если ключ существует. В противном случае добавьте пару
ключ-значение в кеш. Если количество ключей превышает емкость этой операции, удалите ключ, который использовался реже
всего. Каждая из функций get и put должна выполняться со средней временной сложностью O(1).

Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4
*/

/**
 * @param {number} capacity
 */
var LRUCache = function(capacity) {
    this.queue = new Map();
    this.capacity = capacity;
};

/**
 * @param {number} key
 * @return {number}
 */
LRUCache.prototype.get = function(key) {
    if(this.queue.has(key)){
        const value = this.queue.get(key);
        this.queue.delete(key);
        this.queue.set(key, value);
        return value;
    }
    return -1;
};

/**
 * @param {number} key
 * @param {number} value
 * @return {void}
 */
LRUCache.prototype.put = function(key, value) {
    if(this.queue.has(key)) {
        this.queue.delete(key);
        this.queue.set(key, value);
    } else {
        if(this.queue.size >= this.capacity){
            const [firstKey] = this.queue.keys();
            this.queue.delete(firstKey);
        }
        this.queue.set(key, value)
    }
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * var obj = new LRUCache(capacity)
 * var param_1 = obj.get(key)
 * obj.put(key,value)
 */
