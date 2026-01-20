package binarySearch.TimeBasedKeyValueStore;

/* 981. Time Based Key-Value Store
https://leetcode.com/problems/time-based-key-value-store/description/

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

public class TimeBasedKeyValueStore {
    public static void main(String[] args) {
        TimeMap timeMap = new TimeMap();
        timeMap.set("foo", "bar", 1);
        System.out.println(timeMap.get("foo", 1)); // "bar"
        System.out.println(timeMap.get("foo", 3)); // "bar" (latest before 3)
        timeMap.set("foo", "bar2", 4);
        System.out.println(timeMap.get("foo", 4)); // "bar2"
        System.out.println(timeMap.get("foo", 5)); // "bar2"
        System.out.println(timeMap.get("foo", 0)); // "" 
    }
}
