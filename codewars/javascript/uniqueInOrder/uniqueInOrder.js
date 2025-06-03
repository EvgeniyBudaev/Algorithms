/* Функция unique_in_order, которая принимает в качестве аргумента
 последовательность и возвращает список элементов без каких-либо элементов с
  одинаковым значением рядом друг с другом и с сохранением исходного порядка
   элементов. */

var uniqueInOrder=function(iterable){
  return [...iterable].filter((v,i,a)=>v!==a[i-1]);
}

console.log(uniqueInOrder('AAAABBBCCDAABBB'));
