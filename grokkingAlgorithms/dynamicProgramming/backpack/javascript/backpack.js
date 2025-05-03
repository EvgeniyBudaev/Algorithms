// https://www.youtube.com/watch?v=5raYU-45iq8&ab_channel=EnglishPractice4Every1
// Задача о рюкзаке
// Требуется подобрать набор товаров максимальной стоимости, которые можно сложить в рюкзак

const items = [
    {name: 'tapeRecorder', weight: 4, price: 3000},
    {name: 'laptop', weight: 3, price: 2000},
    {name: 'guitar', weight: 1, price: 1500},
];
const capacity = 4;

const getGoodsWithMaxPrice = (items, capacity) => {
    let dynamicProgramming = new Array(items.length + 1)
        .fill(0).map(() => new Array(capacity + 1).fill(0));

    for (let i = 1; i <= items.length; i++) {
        for (let j = 1; j <= capacity; j++) {
            const {weight, price} = items[i - 1];
            if (weight <= j) {
                dynamicProgramming[i][j] =
                    Math.max(price + dynamicProgramming[i - 1][j - weight], dynamicProgramming[i - 1][j]);
            } else {
                dynamicProgramming[i][j] = dynamicProgramming[i - 1][j];
            }
        }
    }

    let selectedItems = [];
    let i = items.length;
    let j = capacity;

    while (i > 0 && j > 0) {
        if (dynamicProgramming[i][j] !== dynamicProgramming[i - 1][j]) {
            selectedItems.push(items[i - 1].name);
            j -= items[i - 1].weight;
        }
        i--;
    }

    return {
        maxPrice: dynamicProgramming[items.length][capacity],
        selectedItems: selectedItems.reverse(),
    };
};

console.log(getGoodsWithMaxPrice(items, capacity));