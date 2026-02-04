package leetcode.java.hashing.FindPlayersWithZeroOrOneLosses;

import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/* 2225. Find Players With Zero or One Losses
https://leetcode.com/problems/find-players-with-zero-or-one-losses/description/

Вам дан целочисленный массив matchs, где match[i] = [winneri,loseri] указывает, что игрок-победитель победил
игрока-проигравшего в матче.
Вернуть ответ в виде списка размером 2, где:

ответ[0] — список всех игроков, которые не проиграли ни одного матча.
ответ[1] — список всех игроков, проигравших ровно один матч.
Значения в двух списках должны возвращаться в порядке возрастания.

Примечание:
Следует учитывать только игроков, сыгравших хотя бы один матч.
Тестовые случаи будут сгенерированы таким образом, чтобы никакие два совпадения не имели одинаковый результат.

Input: matches = [[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]]
Output: [[1,2,10],[4,5,7,8]]
Explanation:
Players 1, 2, and 10 have not lost any matches.
Players 4, 5, 7, and 8 each have lost one match.
Players 3, 6, and 9 each have lost two matches.
Thus, answer[0] = [1,2,10] and answer[1] = [4,5,7,8].

Input: matches = [[2,3],[1,3],[5,4],[6,4]]
Output: [[1,2,5,6],[]]
Explanation:
Players 1, 2, 5, and 6 have not lost any matches.
Players 3 and 4 each have lost two matches.
Thus, answer[0] = [1,2,5,6] and answer[1] = [].
*/

public class FindPlayersWithZeroOrOneLosses {
    public static void main(String[] args) {
        int[][] matches = {
            {1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7},
            {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}
        };
        System.out.println(findWinners(matches)); // [[1, 2, 10], [4, 5, 7, 8]]
    }

    // findWinners - функция для нахождения игроков, которые не проиграли ни одного матча или с одним проигрышем.
    // time: O(n), space: O(n), где n - количество матчей
    private static List<List<Integer>> findWinners(int[][] matches) {
        Map<Integer, Integer> playerLosses = new HashMap<>(); // Карта для подсчета проигрышей для каждого игрока

        // Подсчёт проигрышей для каждого игрока
        for (int[] match : matches) { // Проходим по матчам
            int winner = match[0]; // Победитель
            int loser = match[1];  // Проигравший

            // Инициализируем победителя с 0 проигрышами, если его ещё нет в карте
            playerLosses.putIfAbsent(winner, 0);

            // Увеличиваем счётчик проигрышей проигравшего
            playerLosses.put(loser, playerLosses.getOrDefault(loser, 0) + 1);
        }

        // Подготовка списков для ответа
        List<Integer> zeroLosses = new ArrayList<>(); // Игроки с 0 поражениями
        List<Integer> oneLoss = new ArrayList<>();    // Игроки с 1 поражением

        // Классификация игроков по количеству проигрышей
        for (Map.Entry<Integer, Integer> entry : playerLosses.entrySet()) {
            int player = entry.getKey();
            int losses = entry.getValue();

            if (losses == 0) {
                zeroLosses.add(player);
            } else if (losses == 1) {
                oneLoss.add(player);
            }
        }

        // Сортировка результатов
        Collections.sort(zeroLosses); // Сортируем игроков с 0 проигрышами
        Collections.sort(oneLoss); // Сортируем игроков с 1 проигрышем

        // Формирование итогового ответа
        List<List<Integer>> result = new ArrayList<>();
        result.add(zeroLosses);
        result.add(oneLoss);

        return result;
    }
}
