package main

import "fmt"

/* 146. LRU Cache
https://leetcode.com/problems/lru-cache/description/

Разработайте структуру данных, которая следует ограничениям кэша Least Recently Used (LRU).

Реализуйте класс LRUCache:

LRUCache(int capacity) Инициализирует кэш LRU с положительным размером емкости.
int get(int key) Возвращает значение ключа, если ключ существует, в противном случае возвращает -1.
void put(int key, int value) Обновляет значение ключа, если ключ существует. В противном случае добавляет пару
ключ-значение в кэш. Если количество ключей превышает емкость этой операции, удаляет ключ, который использовался
меньше всего.
Функции get и put должны выполняться со средней временной сложностью O(1).

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

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1)) // 1
	lru.Put(3, 3)
	fmt.Println(lru.Get(2)) // -1
	lru.Put(4, 4)
	fmt.Println(lru.Get(1)) // -1
	fmt.Println(lru.Get(3)) // 3
	fmt.Println(lru.Get(4)) // 4
}

// Node - Узел двусвязного списка.
type Node struct {
	key   int   // Ключ узла
	value int   // Значение ключа
	prev  *Node // Ссылка на предыдущий узел
	next  *Node // Ссылка на следующий узел
}

// LRUCache - структура данных LRU.
type LRUCache struct {
	capacity int           // Емкость кэша
	cache    map[int]*Node // Карта для хранения узлов кэша
	head     *Node         // фиктивный головной узел (самый недавно использованный)
	tail     *Node         // фиктивный хвостовой узел (самый давно использованный)
}

// Constructor - конструктор структуры данных.
func Constructor(capacity int) LRUCache {
	lru := LRUCache{ // Создаем экземпляр структуры данных
		capacity: capacity,                // Задаем емкость кэша
		cache:    make(map[int]*Node),     // Создаем пустую карту для кэша
		head:     &Node{key: 0, value: 0}, // Создаем фиктивный головной узел
		tail:     &Node{key: 0, value: 0}, // Создаем фиктивный хвостовой узел
	}
	lru.head.next = lru.tail // Переназначаем ссылки соседних узлов
	lru.tail.prev = lru.head // Переназначаем ссылки соседних узлов
	return lru               // Возвращаем созданный экземпляр структуры данных
}

// remove - функция для удаления узла из списка.
// time: O(1), space: O(1)
func (this *LRUCache) remove(node *Node) {
	prev := node.prev // Получаем ссылки на соседние узлы
	next := node.next // Получаем ссылки на соседние узлы
	prev.next = next  // Переназначаем ссылки соседних узлов
	next.prev = prev  // Переназначаем ссылки соседних узлов
}

// add -  функция добавления узла в начало списка (после head).
// time: O(1), space: O(1)
func (this *LRUCache) add(node *Node) {
	node.prev = this.head      // Присваиваем ссылки на соседние узлы
	node.next = this.head.next // Присваиваем ссылки на соседние узлы
	this.head.next.prev = node // Присваиваем ссылки на соседние узлы
	this.head.next = node      // Присваиваем ссылки на соседние узлы
}

// moveToHead -  функция перемещения узла в начало списка.
// time: O(1), space: O(1)
func (this *LRUCache) moveToHead(node *Node) {
	this.remove(node) // Удаляем из текущего места
	this.add(node)    // Добавляем в начало
}

// Get - функция получения значения по ключу.
// time: O(1), space: O(1)
func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node) // Перемещаем в начало
		return node.value     // Возвращаем значение
	}
	return -1
}

// Put - функция добавления ключа и значения в кэш.
// time: O(1), space: O(1)
func (this *LRUCache) Put(key int, value int) {
	// Если ключ уже существует, обновляем значение и перемещаем в начало
	if node, ok := this.cache[key]; ok {
		node.value = value    // Обновляем значение
		this.moveToHead(node) // Перемещаем в начало
		return
	}

	// Если достигли capacity, удаляем последний элемент
	if len(this.cache) == this.capacity {
		last := this.tail.prev       // Получаем последний элемент
		this.remove(last)            // Удаляем из списка
		delete(this.cache, last.key) // Удаляем из кэша
	}

	// Создаем новый узел и добавляем в начало
	newNode := &Node{key: key, value: value} // Создаем новый узел
	this.cache[key] = newNode                // Добавляем в кэш
	this.add(newNode)                        // Добавляем в начало списка
}
