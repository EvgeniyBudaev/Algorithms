/* https://leetcode.com/problems/implement-trie-prefix-tree/description/

Дерево (произносится как «попробуй») или префиксное дерево — это древовидная структура данных, используемая для
эффективного хранения и извлечения ключей в наборе строк. Существуют различные приложения этой структуры данных, такие
как автозаполнение и проверка орфографии.

Реализуйте класс Trie:

Trie() Инициализирует объект дерева.
void Insert(String word) Вставляет строковое слово в дерево.
boolean search(String word) Возвращает true, если строковое слово находится в дереве (т. е. было вставлено ранее), и
false в противном случае. Логическое значение startWith(String prefix) Возвращает true, если ранее было вставлено
строковое слово с префиксом-префиксом, и false в противном случае.

Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]

Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True
*/

var Trie = function() {
    this.try = {};
};

/**
 * @param {string} word
 * @return {void}
 */
Trie.prototype.insert = function(word) {
    let curr = this.try;
    for(let c of word) {
        if(!curr[c]) curr[c] = {};
        curr = curr[c];
    }
    curr.isWord = true;
};

/**
 * @param {string} word
 * @return {boolean}
 */
Trie.prototype.search = function(word) {
    let curr = this.try;
    for(let c of word) {
        curr = curr[c];
        if(curr == null) return false;
    }
    return !!curr.isWord;
};

/**
 * @param {string} prefix
 * @return {boolean}
 */
Trie.prototype.startsWith = function(prefix) {
    let curr = this.try;
    for(let c of prefix) {
        curr = curr[c];
        if(curr == null) return false;
    }
    return true;
};

/**
 * Your Trie object will be instantiated and called as such:
 * var obj = new Trie()
 * obj.insert(word)
 * var param_2 = obj.search(word)
 * var param_3 = obj.startsWith(prefix)
 */