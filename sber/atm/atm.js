const atmStore = [
    ["1000", 2],
    ["500", 1],
    ["100", 9],
    ["50", 99],
]; // total 8350

const getPossibleMoney = (atm, total) => {
    return atm.reduce((acc, item) => {
        const banknote = item[0];
        const quantityBanknotesInATM = item[1];
        const possibleQuantityForGet = Math.min(quantityBanknotesInATM, Math.floor(total / Number(banknote)));
        total -= possibleQuantityForGet * Number(banknote);
        acc[banknote] = possibleQuantityForGet;
        return acc;
    }, {});
};

console.log(getPossibleMoney(atmStore, 4565)); // {50: 23, 100: 9, 500: 1, 1000: 2} -> 4550

/* --- если в банкомате бесконечное кол-во купюр --- */

const atm = [
    ["1000", 0],
    ["500", 0],
    ["100", 0],
    ["50", 0],
];

const getMoney = (atm, total) => {
    return atm.reduce((acc, item) => {
        const banknote = item[0];
        const possibleQuantityForGet = Math.floor(total / banknote);
        total -= possibleQuantityForGet * Number(banknote);
        acc[banknote] = possibleQuantityForGet;
        return acc;
    }, {});
};

console.log(getMoney(atm, 4565)); // {50: 1, 100: 0, 500: 1, 1000: 4} -> 4550