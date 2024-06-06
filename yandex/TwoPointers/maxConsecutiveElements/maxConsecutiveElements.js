// maxConsecutiveElements - функция для нахождения максимального количества элементов в строке
function maxConsecutiveElements(inputStr) {
    let result = 0;
    let curLetterIdx = 0;
    let length = inputStr.length;

    while (curLetterIdx < length) {
        let nextLetterIdx = curLetterIdx;
        while (nextLetterIdx < length &&
        inputStr[nextLetterIdx] === inputStr[curLetterIdx]) {
            nextLetterIdx++;
        }
        result = Math.max(result, nextLetterIdx - curLetterIdx);
        curLetterIdx = nextLetterIdx;
    }

    return result;
}

console.log(maxConsecutiveElements("abb")); // 2
