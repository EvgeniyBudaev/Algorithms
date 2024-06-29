// Данные о фруктах
const fruitsData = [
    { size: 5, weight: 150, type: 'orange' },
    { size: 6, weight: 180, type: 'grapefruit' },
    // Добавьте больше данных о фруктах здесь
];

// Функция для вычисления евклидова расстояния
function euclideanDistance(pointA, pointB) {
    return Math.sqrt(
        Math.pow(pointA.size - pointB.size, 2) +
        Math.pow(pointA.weight - pointB.weight, 2)
    );
}

// Функция для определения k ближайших соседей
function findKNearestNeighbors(newPoint, k) {
    let distances = [];
    fruitsData.forEach(fruit => {
        const distance = euclideanDistance(newPoint, fruit);
        distances.push({ fruit, distance });
    });

    // Сортировка по возрастанию расстояния
    distances.sort((a, b) => a.distance - b.distance);

    // Выбираем первые k элементов
    return distances.slice(0, k).map(item => item.fruit);
}

// Функция для классификации нового фрукта
function classifyFruit(newPoint, k) {
    const neighbors = findKNearestNeighbors(newPoint, k);
    let orangeCount = 0;
    let grapefruitCount = 0;

    neighbors.forEach(fruit => {
        if (fruit.type === 'orange') {
            orangeCount++;
        } else if (fruit.type === 'grapefruit') {
            grapefruitCount++;
        }
    });

    return orangeCount > grapefruitCount? 'orange' : 'grapefruit';
}

// Пример использования
const newFruit = { size: 7, weight: 200 };
console.log(`The new fruit is classified as: ${classifyFruit(newFruit, 3)}`); // grapefruit