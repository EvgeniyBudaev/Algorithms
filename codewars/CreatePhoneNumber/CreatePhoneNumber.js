const numbers= [1, 2, 3, 4, 5, 6, 7, 8, 9, 0];

function createPhoneNumber(numbers){
  return `(${numbers.slice(0,3).join('')}) ${numbers.slice(3,6).join('')}-${numbers.slice(6).join('')}`;
}

// function createPhoneNumber(numbers){
// 	let format = '(xxx) xxx-xxxx';
//
// 	for(let i = 0; i < numbers.length; i++)
// 	{
// 		format = format.replace('x', numbers[i]);
// 	}
//
// 	return format;
// }

console.log(createPhoneNumber(numbers)); // "(123) 456-7890"
