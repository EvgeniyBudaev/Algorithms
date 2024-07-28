/* https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/
solution https://leetcode.com/problems/longest-word-in-dictionary/solutions/5484506/js-rust-trie-approach-w-explanation/

Учитывая массив строковых слов, представляющих словарь английского языка, верните самое длинное слово в словах, которое
может быть построено по одному символу из других слов в словах.
Если существует более одного возможного ответа, верните самое длинное слово с наименьшим лексикографическим порядком.
Если ответа нет, верните пустую строку.
Обратите внимание, что слово должно строиться слева направо, при этом каждый дополнительный символ добавляется в конец
предыдущего слова.

Input: words = ["w","wo","wor","worl","world"]
Output: "world"
Пояснение: Слово «мир» может состоять из букв «w», «wo», «wor» и «worl» по одному символу.

Input: words = ["a","banana","app","appl","ap","apply","apple"]
Output: "apple"
Explanation: Both "apply" and "apple" can be built from other words in the dictionary. However, "apple" is
 lexicographically smaller than "apply".
*/

/**
 * @param {string[]} words
 * @return {string}
 */
var longestWord = function(words) {
    const trie = new Trie();

    for (const word of words) {
        trie.insert(word);
    }

    return trie.bfs();
};

class TrieNode {
    constructor() {
        this.children = new Map();
        this.word = null;
    }
}

class Trie {
    constructor() {
        this.root = new TrieNode();
    }

    insert(word) {
        let node = this.root;

        for (const char of word) {
            let child = node.children.get(char);

            if (!child) {
                child = new TrieNode();
                node.children.set(char, child);
            }

            node = child;
        }

        node.word = word;
    }

    bfs() {
        const queue = [this.root];
        let ans = '';

        while (queue.length > 0) {
            const currentNode = queue.shift();

            for (const child of currentNode.children.values()) {
                const word = child.word;

                if (word) {
                    if (
                        word.length > ans.length ||
                        (word.length === ans.length && word < ans)
                    ) {
                        ans = word;
                    }

                    queue.push(child);
                }
            }
        }

        return ans;
    }
}

const words = ["w","wo","wor","worl","world"];
console.log(longestWord(words)); // "world"