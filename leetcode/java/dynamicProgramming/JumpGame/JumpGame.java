package leetcode.java.dynamicProgramming.JumpGame;

/* 55. Jump Game
https://leetcode.com/problems/jump-game/description/

Вам дан целочисленный массив чисел. Изначально вы находитесь в первом индексе массива, и каждый элемент массива
представляет максимальную длину прыжка в этой позиции.
Возвращайте true, если вы можете достичь последнего индекса, или false в противном случае.

Input: nums = [2,3,1,1,4]
Output: true
Пояснение: Перейти на 1 шаг от индекса 0 к 1, затем на 3 шага до последнего индекса.

Input: nums = [3,2,1,0,4]
Output: false
Пояснение: Несмотря ни на что, вы всегда будете достигать индекса 3. Его максимальная длина перехода равна 0, что делает
невозможным достижение последнего индекса
*/

public class JumpGame {
    public static void main(String[] args) {
        
    }

    // canJump проверяет, можно ли добраться до последнего элемента массива.
    // time: O(n), space: O(1)
    private static boolean canJump(int[] nums) {
        // Если массив пустой или содержит один элемент - сразу возвращаем true
        if (nums.length <= 1) {
            return true;
        }

        int maximum = nums[0]; // maximum - максимальная позиция, до которой можно добраться

        for (int i = 0; i < nums.length; i++) {
            // Если текущая позиция недостижима и здесь стоит 0 - возвращаем false
            if (maximum <= i && nums[i] == 0) {
                return false;
            }

            // Обновляем максимальную достижимую позицию
            if (i + nums[i] > maximum) {
                maximum = i + nums[i];
            }

            // Если можем достичь последнего элемента - возвращаем true
            if (maximum >= nums.length - 1) {
                return true;
            }
        }

        return false; // Если не удалось добраться до конца, возвращаем false
    }
}
