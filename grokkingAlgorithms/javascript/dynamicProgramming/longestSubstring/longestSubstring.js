// Самая длинная общая подстрока

/*
Описание: Допустим, вы открыли сайт dictionary.com. Пользователь вводит слово, а сайт возвращает определение.
Но если пользователь ввел несуществующее слово, нужно предположить, какое слово имелось ввиду.
Пользователь ищет определение fish, но он случайно ввел hish. Такого слова в словаре нет, но зато есть список похожих
слов: fish, vista.
*/

/**
 * Create a matrix
 * @param {number} rows Number of rows
 * @param {number} columns ANumber of columns
 * @returns {Array} matrix
 */
const createMatrix = (rows = 0, columns = 0) => {
    const matrix = [];

    for (let i = 0; i < rows; i++) {
        matrix[i] = Array(columns).fill(0);
    }

    return matrix;
};

/**
 * Find the longest substring
 * @param {string} firstWord First word
 * @param {string} secondWord Second word
 * @returns {string} The longest substring
 */
const longestSubstring = (firstWord = "", secondWord = "") => {
    const matrix = JSON.parse(
        JSON.stringify(createMatrix(firstWord.length, secondWord.length))
    );
    let sizeSequence = 0;
    let indexSequence = 0;

    for (let i = 0; i < firstWord.length; i++) {
        for (let j = 0; j < secondWord.length; j++) {
            if (firstWord[i] === secondWord[j]) {
                matrix[i][j] = (i && j) > 0 ? matrix[i - 1][j - 1] + 1 : 1;

                if (matrix[i][j] >= sizeSequence) {
                    sizeSequence = matrix[i][j];
                    indexSequence = i + 1;
                }
            } else {
                matrix[i][j] = 0;
            }
        }
    }

    return firstWord.slice(indexSequence - sizeSequence, indexSequence);
};

longestSubstring("vista", "hish"); // "is"
longestSubstring("fish", "hish"); // "ish"