/* https://leetcode.com/problems/find-the-smallest-divisor-given-a-threshold/description/
Учитывая массив целых чисел nums и целочисленный threshold, мы выберем положительный целочисленный divisor, разделим на
него весь массив и просуммируем результат деления. Найдите наименьший divisor, чтобы упомянутый выше результат был
меньше или равен threshold.
Каждый результат деления округляется до ближайшего целого числа, большего или равного этому элементу.
(Например: 7/3 = 3 и 10/2 = 5).
Тестовые случаи генерируются для того, чтобы был ответ.

Input: nums = [1,2,5,9], threshold = 6
Output: 5
Пояснение: Мы можем получить сумму 17 (1+2+5+9), если делитель равен 1.
Если делитель равен 4, мы можем получить сумму 7 (1+1+2+3), а если делитель равен 5, сумма будет 5 (1+1+1+2).

Input: nums = [44,22,33,11,1], threshold = 5
Output: 44
 */

/**
 * @param {number[]} nums
 * @param {number} threshold
 * @return {number}
 */
var smallestDivisor = function(nums, threshold) {
    let left = 1;
    // Использование деструктуризации массива с Math.max — это изящный,
    // простой способ найти максимальное значение в массиве чисел
    let right = Math.max(...nums);

    // sumDivisions - функция для суммирования делителей
    let sumDivisions = (div) => {
        let total = 0;

        for (const num of nums) {
            total += Math.ceil(num / div);
        }

        return total;
    };

    while (left <= right) {
        let mid = Math.floor((left + right) / 2);

        if (sumDivisions(mid) <= threshold) {
            // Половина проблемного пространства за счет снижения «макс.»
            right = mid - 1;
        } else {
            // Половина проблемного пространства за счет увеличения «мин»
            left = mid + 1;
        }
    }

    // слева сейчас находится на минимальном пороге
    return left;
};

console.log(smallestDivisor([1,2,5,9], 6)); // 5