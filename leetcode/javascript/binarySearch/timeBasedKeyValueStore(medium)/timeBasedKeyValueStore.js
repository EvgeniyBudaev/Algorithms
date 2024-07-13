/* https://leetcode.com/problems/time-based-key-value-store/description/

Разработайте структуру данных «ключ-значение», основанную на времени, которая может хранить несколько значений для
одного и того же ключа с разными метками времени и получать значение ключа в определенную метку времени.

Реализуйте класс TimeMap:

TimeMap() Инициализирует объект структуры данных.
void set(String key, String value, int timestamp) Сохраняет ключевой ключ со значением значения в заданную временную
метку.
String get(String key, int timestamp) Возвращает значение, такое, которое set было вызвано ранее,
с timestamp_prev <= timestamp. Если таких значений несколько, возвращается значение, связанное с наибольшим значением
timestamp_prev. Если значений нет, возвращается «».

Input
["TimeMap", "set", "get", "get", "set", "get", "get"]
[[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
Output
[null, null, "bar", "bar", null, "bar2", "bar2"]

Explanation
TimeMap timeMap = new TimeMap();
timeMap.set("foo", "bar", 1);  // store the key "foo" and value "bar" along with timestamp = 1.
timeMap.get("foo", 1);         // return "bar"
timeMap.get("foo", 3);         // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
timeMap.set("foo", "bar2", 4); // store the key "foo" and value "bar2" along with timestamp = 4.
timeMap.get("foo", 4);         // return "bar2"
timeMap.get("foo", 5);         // return "bar2"
*/

var TimeMap = function() {
    this.map = new Map();
};

/**
 * @param {string} key
 * @param {string} value
 * @param {number} timestamp
 * @return {void}
 */
TimeMap.prototype.set = function(key, value, timestamp) {
    const map = this.map;
    if (!map.has(key)) map.set(key, []);
    map.get(key).push([value, timestamp]);
};

/**
 * @param {string} key
 * @param {number} timestamp
 * @return {string}
 */
TimeMap.prototype.get = function(key, timestamp) {
    const arr = this.map.get(key) || [];

    let [low, high] = [0, arr.length - 1];
    let res = "";
    while (low <= high) {
        const mid = Math.floor((low + high) / 2);
        const [v, t] = arr[mid];
        if (timestamp === t) return v;
        if (timestamp >= t) {
            low = mid + 1;
            res = v;
        } else high = mid - 1;

    }
    return res;
};

/**
 * Your TimeMap object will be instantiated and called as such:
 * var obj = new TimeMap()
 * obj.set(key,value,timestamp)
 * var param_2 = obj.get(key,timestamp)
 */