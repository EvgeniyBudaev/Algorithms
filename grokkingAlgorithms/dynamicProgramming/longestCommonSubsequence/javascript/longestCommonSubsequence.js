// Самая длинная общая подпоследовательность

/*
Описание: Допустим, вы открыли сайт dictionary.com. Пользователь вводит слово, а сайт возвращает определение.
Но если пользователь ввел несуществующее слово, нужно предположить, какое слово имелось ввиду.
Пользователь ищет определение fish, но он случайно ввел hish. Такого слова в словаре нет, но зато есть список похожих
слов: fish, vista.
*/

/**
 * Find the longest common subsequence
 * @param {string} firstWord First word
 * @param {string} secondWord Second word
 * @returns {number} The longest common subsequence
 */
const longestCommonSubsequence = (firstWord = "", secondWord = "") => {
    const matrix = JSON.parse(
        JSON.stringify(createMatrix(firstWord.length, secondWord.length))
    );
    if (matrix.length === 0 || matrix[0].length === 0) return 0;
    for (let i = 0; i < firstWord.length; i++) {
        for (let j = 0; j < secondWord.length; j++) {
            if (firstWord[i] === secondWord[j]) {
                matrix[i][j] = (i && j) > 0 ? matrix[i - 1][j - 1] + 1 : 1;
            } else {
                matrix[i][j] = Math.max(
                    i > 0 ? matrix[i - 1][j] : 0,
                    j > 0 ? matrix[i][j - 1] : 0
                );
            }
        }
    }
    return matrix[firstWord.length - 1][secondWord.length - 1];
};

longestCommonSubsequence("fish", "fosh"); // 3
longestCommonSubsequence("fort", "fosh"); // 2