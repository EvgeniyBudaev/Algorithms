/* https://github.com/EvgeniyBudaev/doocs_leetcode/blob/main/solution/0100-0199/0159.Longest%20Substring%20with%20At%20Most%20Two%20Distinct%20Characters/README_EN.md

Учитывая строку s и целое число k, верните длину самой длинной подстроки s, содержащей не более k различных символов.

Input: s = "eceba", k = 2
Output: 3
Объяснение: Это подстрока «ece» длиной 3.

Input: s = "aa", k = 1
Output: 2
Объяснение: Это подстрока «aa» длиной 2.
*/

function lengthOfLongestSubstringKDistinct(s, k) {
    const charFrequency = new Map();
    let maxLength = 0;

    for (let left = 0, right = 0; right < s.length; right++) {
        const currentChar = s[right];
        // Увеличить частоту текущего символа на карте.
        const frequency = (charFrequency.get(currentChar) || 0) + 1;
        charFrequency.set(currentChar, frequency);

        // Если имеется более двух разных символов, сузьте окно слева.
        while (charFrequency.size > k) {
            // Уменьшите частоту появления символа у «левого» указателя.
            const leftChar = s[left];
            const leftFrequency = (charFrequency.get(leftChar) || 0) - 1;
            // Если частота не равна нулю, обновите ее на карте.
            if (leftFrequency > 0) charFrequency.set(leftChar, leftFrequency);
            // Если частота равна нулю, удалите ее с карты.
            else charFrequency.delete(leftChar);
            left++;
        }

        // Вычислить текущую длину подстроки
        const currentLength = right - left + 1;
        maxLength = Math.max(maxLength, currentLength);
    }

    return maxLength;
}

console.log(lengthOfLongestSubstringKDistinct("eceba", 2)); // 3