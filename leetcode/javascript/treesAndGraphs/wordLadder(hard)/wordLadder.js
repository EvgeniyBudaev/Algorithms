/*
Последовательность преобразования слова BeginWord в слово EndWord с использованием словаря wordList — это
последовательность слов BeginWord -> s1 -> s2 -> ... -> sk такая, что:

Каждая соседняя пара слов отличается одной буквой.
Каждый si для 1 <= i <= k находится в wordList. Обратите внимание, что BeginWord не обязательно должен находиться в
списке слов.
sk == endWord
Учитывая два слова, BeginWord и EndWord, и словарь WordList, возвращает количество слов в кратчайшей последовательности преобразования из BeginWord в EndWord или 0, если такой последовательности не существует.
*/

/*
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.
 */

/**
 * @param {string} beginWord
 * @param {string} endWord
 * @param {string[]} wordList
 * @return {number}
 */
var ladderLength = function(beginWord, endWord, wordList) {
    // Если BeginWord совпадает с EndWord, преобразование тривиально.
    if (beginWord === endWord) return 1;

    // Преобразование wordList в набор для поиска O (1)
    const wordSet = new Set(wordList);
    if (!wordSet.has(endWord)) return 0;

    let beginSet = new Set();
    let endSet = new Set();
    beginSet.add(beginWord);
    endSet.add(endWord);
    let steps = 1;

    while (beginSet.size !== 0 && endSet.size !== 0) {
        // Всегда расширяйте меньший набор, чтобы минимизировать количество операций
        if (beginSet.size > endSet.size) {
            let temp = beginSet;
            beginSet = endSet;
            endSet = temp;
        }

        let newSet = new Set();
        // Перебирать каждое слово в текущем наборе
        for (let word of beginSet) {
            // Перебираем каждый символ слова
            for (let i = 0; i < word.length; i++) {
                // Попробуйте заменить символ каждой буквой от «а» до «я»
                for (let j = 0; j < 26; j++) {
                    const newWord = word.slice(0, i) + String.fromCharCode(j + 97) + word.slice(i + 1);
                    // Если новое слово находится в конечном наборе, мы нашли кратчайший путь
                    if (endSet.has(newWord)) return steps + 1;
                    // Если новое слово есть в списке слов и его еще не посещали
                    if (wordSet.has(newWord)) {
                        newSet.add(newWord);
                        wordSet.delete(newWord); // Отметить как посещенное
                    }
                }
            }
        }

        // Перейти на следующий уровень
        beginSet = newSet;
        steps++;
    }
    return 0;
};

const beginWord = "hit";
const endWord = "cog";
const wordList = ["hot","dot","dog","lot","log","cog"];
console.log(ladderLength(beginWord, endWord, wordList)); // 5