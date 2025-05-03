// Приближенный алгоритм
// Задача о покрытии множества
// Нужно найти минимальный набор станций, который бы покрывал все штаты
let statesNeeded = new Set(["mt", "wa", "or", "id", "nv", "ut", "ca", "az"]); // Список штатов

const stations = {}; // Список станций
// Ключчи - названия станций, значения - сокращенные обозначения штатов, входящих в зону охвата
stations.kone = new Set(["id", "nv", "ut"]);
stations.ktwo = new Set(["wa", "id", "mt"]);
stations.kthree = new Set(["or", "nv", "ca"]);
stations.kfour = new Set(["nv", "ut"]);
stations.kfive = new Set(["ca", "az"]);

const finalStations = new Set(); // Итоговой набор станций, покрывающий все штаты

while (statesNeeded.size) {
    // bestStation - станция, которая покрывает больше всего штатов, не входящих в зону покрытия
    let bestStation = null;
    // statesCovered - список штатов, обслуживаемые этой станцией, и которые еще не входят в зону покрытия
    let statesCovered = new Set();
    Object.keys(stations).forEach(station => {
        const states = stations[station];
        const covered = new Set([...statesNeeded].filter(x => states.has(x)));
        if (covered.size > statesCovered.size) {
            bestStation = station;
            statesCovered = covered;
        }
    });
    statesNeeded = new Set([...statesNeeded].filter(x => !statesCovered.has(x)));
    finalStations.add(bestStation);
}

console.log(finalStations); // Set { 'kone', 'ktwo', 'kthree', 'kfive' }