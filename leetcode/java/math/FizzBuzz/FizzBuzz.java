package leetcode.java.math.FizzBuzz;

import java.util.ArrayList;
import java.util.List;

/* 412. Fizz Buzz
https://leetcode.com/problems/fizz-buzz/description/

Дано целое число n, вернуть ответ в виде массива строк (индексированный на 1), где:

answer[i] == "FizzBuzz", если i делится на 3 и 5.
answer[i] == "Fizz", если i делится на 3.
answer[i] == "Buzz", если i делится на 5.
answer[i] == i (в виде строки), если ни одно из вышеперечисленных условий не выполняется.

Input: n = 3
Output: ["1","2","Fizz"]
*/

public class FizzBuzz {
    public static void main(String[] args) {
        System.out.println(fizzBuzz(3)); // [1, 2, Fizz]
    }

    // fizzBuzz принимает целое число n и возвращает список строк в соответствии с условием задачи.
    // time: O(n), space: O(n)
    private static List<String> fizzBuzz(int n) {
        List<String> output = new ArrayList<>(); // Инициализация результирующего списка строк

        for (int i = 1; i <= n; i++) {
            if (i % 3 == 0 && i % 5 == 0) {
                output.add("FizzBuzz");
            } else if (i % 3 == 0 && i % 5 != 0) {
                output.add("Fizz");
            } else if (i % 5 == 0 && i % 3 != 0) {
                output.add("Buzz");
            } else {
                output.add(String.valueOf(i));
            }
        }

        return output;
    }
}
