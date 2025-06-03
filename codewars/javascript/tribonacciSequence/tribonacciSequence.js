function tribonacci(signature,n){
  if (n === 0) return [];
  if (n === 1) return [signature[0]];
  if (n === 2) return [signature[0], signature[1]];
  if (n === 3) return [signature[0], signature[1], signature[2]];
  let tribonacciArray = [signature[0], signature[1], signature[2]];
  for (let i = 3; i < n; i++) {
	tribonacciArray.push(tribonacciArray[i - 1] + tribonacciArray[i - 2] + tribonacciArray[i - 3]);
  }
  return tribonacciArray;
}

console.log(tribonacci([1,1,1],10)); // [1,1,1,3,5,9,17,31,57,105]
