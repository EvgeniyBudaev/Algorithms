/* https://leetcode.com/problems/number-of-longest-increasing-subsequence/description/

Учитывая целочисленный массив nums, верните количество самых длинных возрастающих подпоследовательностей.
Обратите внимание, что последовательность должна быть строго возрастающей.

Input: nums = [1,3,5,4,7]
Output: 2
Пояснение: Две самые длинные возрастающие подпоследовательности — это [1, 3, 4, 7] и [1, 3, 5, 7].

Input: nums = [2,2,2,2,2]
Output: 5
Пояснение: Длина самой длинной возрастающей подпоследовательности равна 1, существует 5 возрастающих
подпоследовательностей длины 1, поэтому выведите 5.
*/

/**
 * @param {number[]} nums
 * @return {number}
 */
var findNumberOfLIS = function(nums) {
    // Мы инициализируем массив с именем "tracker" для отслеживания самой длинной возрастающей подпоследовательности,
    // заканчивающейся в каждой индексируем и инициализируем все позиции с помощью 1, поскольку каждый элемент в nums
    // можно рассматривать как подпоследовательность в сам.
    let tracker = new Array(nums.length).fill(1);


    // При отслеживании мы поймем, что иногда в nums имеется более одного набора предшествующих элементов
    // которая может создать самую длинную возрастающую подпоследовательность, заканчивающуюся индексом, который мы
    // сейчас обрабатываем. Лучший способ отслеживать это — использовать отдельный массив, называемый частотой.
    // Здесь будет каждый элемент инициализируется значением 1. Это означает, что изначально может быть только один
    // самый длинный подпоследовательность, заканчивающаяся любым индексом.
    let frequency = new Array(nums.length).fill(1);

    // Теперь начнем обработку слева направо, начиная со второго элемента с индексом 1
    // Пропускаем первый элемент, так как мы не можем найти подпоследовательность длиннее 1, которая может заканчиваться
    // индексом 0.
    for (let i = 0; i < nums.length; i++) {
        // Теперь давайте сравним его с каждым элементом перед ним.
        for (let j = 0; j < i; j++) {
            // ПРОВЕРКА: Быстрая проверка того, что j может быть частью строго возрастающей подпоследовательности,
            // заканчивающейся на i.
            if (nums[j] < nums[i]) {

                // СЛУЧАЙ 1: Соединение i с подпоследовательностью, заканчивающейся на j, делает подпоследовательность
                // больше, чем любая другая подпоследовательность, закончившаяся на i. Это увеличивает
                // подпоследовательность, заканчивающуюся на j всего на 1, отсюда и сложение.
                if (tracker[i] < tracker[j]+1) {
                    // мы обновляем самую длинную подпоследовательность, заканчивающуюся на i
                    tracker[i] = tracker[j]  + 1;

                    // Новая самая длинная подпоследовательность создается для любого пути, который может иметь
                    // закончилось на j-м индексе. Следовательно, мы устанавливаем частоту i так же, как частоту j.
                    frequency[i] = frequency[j];
                } else if (tracker[i] === tracker[j]  + 1) {

                    // СЛУЧАЙ 2: Это означает, что мы не впервые находим j, который помогает сделать
                    // самая длинная подпоследовательность, заканчивающаяся на i. Поскольку длина самой
                    // подпоследовательности не изменилась, никаких обновлений в массиве трекера делать не нужно.

                    // Но это еще нужно отслеживать в нашем массиве частот. Поэтому мы добавляем частоту
                    // все самые длинные подпоследовательности, которые заканчиваются на j, в наш трекер частоты для i.
                    frequency[i] += frequency[j];
                }
            }
        }
    }

    // FINALLY:
    // 1. Мы находим самый длинный путь, который когда-либо видели, используя наш массив трекеров.
    const longestPath = Math.max(...tracker);

    // 2. Для каждого индекса/элемента, у которого заканчивается самый длинный путь, мы добавляем его частоту к нашему
    // общему результату.
    let result = 0;
    for (let k = 0; k < nums.length; k++) {
        if (tracker[k] === longestPath) result += frequency[k];
    }

    return result;
};

console.log(findNumberOfLIS([1,3,5,4,7])); // 2