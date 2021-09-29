// [null, 1, 2, null, undefined, true, {}, () => {}, 3].filter(_.identity);
// => [1, 2, true, Object {}, function(), 3]

const list = [null, 1, 2, null, undefined, true, {}, () => {}, 3];

function identity(value) {
    return value;
}

console.log(list.filter(identity)); // => [1, 2, true, Object {}, function(), 3]