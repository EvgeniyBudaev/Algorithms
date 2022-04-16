/*
Марсоход проезжает на каждой итерации на 1метр
l-налево, r-направо, t-вверх, b-вниз.
Если марсоход поехал налево, а затем направо, то он остался на той же позиции и
в итоге нужно вернуть пустой массив. В ином случае нужно вернуть массив
элементов, в какую сторону поехал марсоход.
*/

// ['l', 'r'] => []
// ['l', 'r', 't', 'b'] => []
// ['l', 'l', 'l', 'r', 't', 'b'] => ['l', 'l']

const moveOptions = {
  l: 'r',
  r: 'l',
  t: 'b',
  b: 't',
};

const getCommon = (list) => {
  const common = [];

  for(let i = 0; i < list.length; i++) {
    const lastItem = common[common.length - 1];

    if (lastItem !== moveOptions[list[i]]) {
      common.push(list[i]);
    } else {
      common.pop();
    }
  }

  return common;
};

// console.log(getCommon(['l', 'r']));
// console.log(getCommon(['l', 'r', 't', 'b']));
// console.log(getCommon(['l', 'l']));
console.log(getCommon(['l', 'l', 'l', 'r', 't', 'b'])); // ['l', 'l']
