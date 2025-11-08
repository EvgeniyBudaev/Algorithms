package array.ReverseBits;

/* 190. Reverse Bits
https://leetcode.com/problems/reverse-bits/description/

Обратные биты заданного 32-битного целого числа без знака.

Примечание:

Обратите внимание, что в некоторых языках, таких как Java, нет типа целого числа без знака. В этом случае и входные, и
выходные данные будут заданы как тип целого числа со знаком. Они не должны влиять на вашу реализацию, поскольку
внутреннее двоичное представление целого числа одинаково, независимо от того, является ли оно знаковым или беззнаковым.
В Java компилятор представляет целые числа со знаком, используя нотацию дополнения до 2. Поэтому в примере 2 выше
входные данные представляют целое число со знаком -3, а выходные данные представляют целое число со знаком -1073741825.

Input: n = 00000010100101000001111010011100
Output:    964176192 (00111001011110000010100101000000)
Explanation: The input binary string 00000010100101000001111010011100 represents the unsigned integer 43261596, so return 964176192 which its binary representation is 00111001011110000010100101000000.
*/

public class ReverseBits {
    public static void main(String[] args) {
        int num = 0b00000010100101000001111010011100; // Объявляем переменную для записи числа
        System.out.println(reverseBits(num)); // 964176192
    }

    // reverseBits - Возвращает обратное число без знака.
    // time: O(n), где n - длина числа, space: O(n)
    private static int reverseBits(int n) {
        StringBuilder revNumString = new StringBuilder(); // Объявляем переменную для записи обратного числа

        // Преобразуем число в двоичную строку с ведущими нулями
        String binaryString = String.format("%32s", Integer.toBinaryString(n)).replace(' ', '0');

        // Перебираем символы числа
        for (char c : binaryString.toCharArray()) {
            revNumString.insert(0, c); // Добавляем символ в начало строки
        }

        // Преобразуем обратное число в целое число без знака
        long revNum = Long.parseLong(revNumString.toString(), 2);
        return (int) revNum; // Возвращаем обратное число
    }
}
