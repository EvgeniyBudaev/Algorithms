/*
Учитывая текстовую строку и слова (список строк), верните все пары индексов [i, j] так, чтобы подстрока
text[i]...text[j] находилась в списке слов.
*/

// TrieNode represents each node of the Trie.
class TrieNode {
    children;
    isEndOfWord;
    // Constructor initializes the children to have 26 positions for each alphabet letter.
    constructor() {
        this.children = Array(26).fill(null); // Initialize all positions to null
        this.isEndOfWord = false;
    }
    // Method to insert a word into the trie.
    insert(word) {
        let node = this; // Start from the root node
        for (const c of word) {
            const index = c.charCodeAt(0) - 'a'.charCodeAt(0); // Convert character to index (0-25)
            // If there's no node for this character, create it.
            if (!node.children[index]) {
                node.children[index] = new TrieNode();
            }
            // Move to the child node
            node = node.children[index];
        }
        // Once all characters are inserted, mark this node as the end of a word.
        node.isEndOfWord = true;
    }
}

// Функция для возврата пар индексов для всех слов в «тексте», присутствующих в массиве «слова».
function indexPairs(text, words) {
    // Создайте Trie для заданного набора слов для эффективного поиска.
    const trie = new TrieNode();
    words.forEach(word => {
        trie.insert(word);
    });

    const n = text.length;
    const result = []; // Для хранения пар индексов.

    // Перебираем каждый символ текста.
    for (let i = 0; i < n; ++i) {
        let node = trie;
        // Начните проверку слов из каждого индекса «i».
        for (let j = i; j < n; j++) {
            const index = text.charCodeAt(j) - 'a'.charCodeAt(0); // Преобразование символа в индекс.
            // Если символ не является началом какого-либо слова, разбейте его.
            if (!node.children[index]) break;
            // В противном случае перейдите к дочернему узлу.
            node = node.children[index];
            // Если узел отмечает конец слова, запишите пару индексов.
            if (node.isEndOfWord) {
                result.push([i, j]);
            }
        }
    }
    // Вернуть список пар индексов.
    return result;
}

const text = "thestoryofleetcodeandme";
const words = ["story","fleet","leetcode"];
console.log(indexPairs(text, words)); // [[3,7],[9,13],[10,17]]