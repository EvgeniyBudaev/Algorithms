console.log('1. Synchronous code');

queueMicrotask(() => {
  console.log('2. queueMicrotask');
});

Promise.resolve().then(() => {
  console.log('3. Promise');
});

setTimeout(() => {
  console.log('4. setTimeout');
}, 0);

console.log('5. Synchronous code');