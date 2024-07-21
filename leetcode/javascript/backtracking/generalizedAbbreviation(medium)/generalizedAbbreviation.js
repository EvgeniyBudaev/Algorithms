/* https://algo.monster/liteproblems/320

Учитывая строковое слово, верните список всех возможных обобщенных сокращений слова. Верните ответ в любом порядке.

Input: word = "word"
Output: ["4","3d","2r1","2rd","1o2","1o1d","1or1","1ord","w3","w2d","w1r1","w1rd","wo2","wo1d","wor1","word"]

Input: word = "a"
Output: ["1","a"]
*/

// Глобальная переменная для хранения окончательного списка сокращений
let abbreviations  = [];

//Функция, инициирующая генерацию сокращений для слова.
function generateAbbreviations(word) {
    abbreviations = [];

    // TВременный массив для отслеживания текущей прогрессии сокращений.
    let currentAbbreviation = [];

    // Запустите поиск в глубину (DFS) с начала слова.
    dfs(word, currentAbbreviation);

    // Возвращает список всех возможных сокращений после завершения DFS.
    return abbreviations;
}

// Вспомогательная функция для выполнения поиска в глубину для создания всех сокращений.
function dfs(remainingSubstring, currentAbbreviation) {
    //
    // Если оставшаяся подстрока пуста, соедините все собранные части и прибавьте к результату.
    if (remainingSubstring === "") {
        abbreviations.push(currentAbbreviation.join(""));
        return;
    }

    // Попробуйте сократить числа всех возможных длин (вплоть до длины оставшейся подстроки).
    for (let i = 1; i <= remainingSubstring.length; i++) {
        // Добавьте сокращение номера для текущей длины.
        currentAbbreviation.push(i.toString());

        // Если символы остались, добавьте следующий символ и продолжите DFS.
        if (i < remainingSubstring.length) {
            currentAbbreviation.push(remainingSubstring.charAt(i));
            dfs(remainingSubstring.substring(i + 1), currentAbbreviation);

            // Возврат — удалите последнего персонажа, чтобы исследовать новые пути.
            currentAbbreviation.pop();
        } else {
            // Если символов не осталось, продолжите DFS с текущим сокращением.
            dfs("", currentAbbreviation);
        }

        // Backtracking - удалите сокращение номера, чтобы попробовать другую длину сокращения.
        currentAbbreviation.pop();
    }

    // Попробуйте оставить первый символ как есть и продолжить DFS для остальной части строки.
    currentAbbreviation.push(remainingSubstring.charAt(0));
    dfs(remainingSubstring.substring(1), currentAbbreviation);

    // Backtracking - удалите персонажа, чтобы дерево путей оставалось чистым для новых путей.
    currentAbbreviation.pop();
}

console.log(generateAbbreviations("word")); // ["4","3d","2r1","2rd","1o2","1o1d","1or1","1ord","w3","w2d","w1r1","w1rd","wo2","wo1d","wor1","word"]