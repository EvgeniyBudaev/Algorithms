const _readline = require('readline');

const _reader = _readline.createInterface({
	input: process.stdin
});

const _inputLines = [];
let _curLine = 0;

_reader.on('line', line => {
	_inputLines.push(line);
});

// Когда ввод закончится, будет вызвана функция solve.
process.stdin.on('end', solve);

function solve() {
	var splitted = _inputLines[0].trim(" ").split(" ").map(num => Number(num));
	const a = splitted[0];
	const b = splitted[1];

	const answer = a + b;

	console.log(answer);
}
