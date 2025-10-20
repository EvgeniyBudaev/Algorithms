package array.HammingDistance;

/* 461. Hamming Distance
https://leetcode.com/problems/hamming-distance/description/

Расстояние Хэмминга между двумя целыми числами — это количество позиций, в которых соответствующие биты различны.
Даны два целых числа x и y, верните расстояние Хэмминга между ними..

Input: x = 1, y = 4
Output: 2
Объяснение:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
Стрелки выше указывают на позиции, в которых соответствующие биты отличаются.
*/

public class HammingDistance {
    public static void main(String[] args) {
        System.out.println(hammingDistance(1, 4)); // 2
    }

    // hammingDistance принимает на вход два целых числа и возвращает расстояние Хэмминга между ними.
    // time: O(1) (константное время) - операции над числами занимают константное время
    // space: O(1) (константное пространство) - используем только константные переменные
    private static int hammingDistance(int x, int y) {
        String c = String.format("%32s", Integer.toBinaryString(x)).replace(' ', '0'); // Переводим в двоичное число с 32 знаками 00000000000000000000000000000001
        String g = String.format("%32s", Integer.toBinaryString(y)).replace(' ', '0'); // Переводим в двоичное число с 32 знаками 00000000000000000000000000000100
        int count = 0; // Счетчик разрядов, где биты отличаются

        for (int i = 0; i < c.length(); i++) { // Проходим по каждому разряду числа
            if (c.charAt(i) != g.charAt(i)) { // Если разряды разные, увеличиваем счетчик
                count++;
            }
        }

        return count; // Возвращаем счетчик
    }
}
