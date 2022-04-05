/* Функция rgb приводит к возврату шестнадцатеричного представления. */

function rgb(r, g, b){
	return toHex(r)+toHex(g)+toHex(b);
}

function toHex(d) {
	if (d < 0 ) { return "00"; }
	if (d > 255 ) { return "FF"; }
	return  ("0"+(Number(d).toString(16))).slice(-2).toUpperCase();
}

// function rgb(r, g, b){
//   return '#' + [r, g, b].map(x => {
// 	let hex = x.toString(16);
// 	  if (hex.length < 2) {
// 		  hex = "0" + hex;
// 	  }
// 	  if (x < 0) {
// 		  hex = "00";
// 	  }
// 	  if (x > 255) {
// 		  hex = "FF";
// 	  }
// 	  return hex.toUpperCase();
//   }).join('');
// }

// console.log(rgb(255, 255, 255)); // #ffffff
// console.log(rgb(255, 255, 300)); // #ffffff
// console.log(rgb(0, 0, 0)); // #000000
console.log(rgb(148, 0, 211)); // #9400d3
// console.log(rgb(0, 0, -20)); // #000000
