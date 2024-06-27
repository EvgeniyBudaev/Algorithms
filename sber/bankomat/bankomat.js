const atm = [
    ["1000", 2],
    ["500", 1],
    ["100", 9],
    ["50", 99],
];

const getMoney = (atm, total) => {
    return atm.reduce((acc, item) => {
        const banknote = item[0];
        const quantityBanknotesInATM = item[1];
        const possibleQuantityForGet = Math.min(quantityBanknotesInATM, Math.floor(total / banknote));
        total -= possibleQuantityForGet * Number(banknote);
        acc[banknote] = possibleQuantityForGet;
        return acc;
    }, {});
};

console.log(getMoney(atm, 4565)); // {50: 23, 100: 9, 500: 1, 1000: 2}