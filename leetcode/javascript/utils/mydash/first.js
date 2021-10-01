// _.first([1, 2, 3, 4, 5]); // => 1

const arr = [1, 2, 3];

const first = (array) => {
  return (array && array.length) ? array[0] : undefined;
};

console.log(first(arr)); // => 1

function firstYandex(list) {
        if (!Array.isArray(list)) {
            return undefined;
        }

        return list.length ? list[0] : undefined;
}
