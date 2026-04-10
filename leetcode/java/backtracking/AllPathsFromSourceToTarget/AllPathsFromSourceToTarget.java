package leetcode.java.backtracking.AllPathsFromSourceToTarget;

import java.util.ArrayList;
import java.util.List;

/* 797. All Paths From Source to Target
 https://leetcode.com/problems/all-paths-from-source-to-target/description/

Дан направленный ациклический граф (DAG) из n узлов, помеченных от 0 до n-1, найдите все возможные пути от узла 0 до
узла n-1 и верните их в любом порядке.

Граф задается следующим образом: граф[i] — это список всех узлов, которые вы можете посетить из узла i
(т. е. существует направленное ребро от узла i к узлу граф[i][j]).

Input: graph = [[1,2],[3],[3],[]]
Output: [[0,1,3],[0,2,3]]
Пояснение: Есть два пути: 0 -> 1 -> 3 и 0 -> 2 -> 3.

Input: graph = [[4,3,1],[3,2,4],[3],[4],[]]
Output: [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
*/

public class AllPathsFromSourceToTarget {
    public static void main(String[] args) {
        int[][] graph = {
                {1, 2},
                {3},
                {3},
                {}
        };

        System.out.println(allPathsSourceTarget(graph));
    }

    /**
     * allPathsSourceTarget находит все пути от начального узла (0) до конечного (n-1) в графе
     * time: O(2^N * N) — в худшем случае экспоненциальное число путей длиной до N
     * space: O(N) — глубина стека рекурсии + O(2^N * N) для хранения всех путей в результате
     * где N — количество узлов в графе
     */
    private static List<List<Integer>> allPathsSourceTarget(int[][] graph) {
        List<List<Integer>> result = new ArrayList<>(); // Здесь будем хранить все найденные пути
        int n = graph.length; // Количество узлов в графе
        List<Integer> path = new ArrayList<>(); // Начинаем путь с узла 0
        path.add(0);

        // Запускаем DFS с начального узла
        dfs(0, n - 1, graph, path, result);

        return result;
    }

    /**
     * Внутренняя рекурсивная функция для обхода в глубину (DFS) с бэктрекингом
     * time: O(2^n * n), space: O(n) — глубина рекурсивного стека
     */
    private static void dfs(int node, int target, int[][] graph,
                            List<Integer> currentPath, List<List<Integer>> result) {
        // Если достигли целевого узла, добавляем текущий путь в результат
        if (node == target) {
            // Важно создать копию пути, чтобы избежать изменений при бэктрекинге
            result.add(new ArrayList<>(currentPath));
            return;
        }

        // Рекурсивно посещаем всех соседей текущего узла
        for (int neighbor : graph[node]) {
            // Добавляем соседа в путь и продолжаем поиск
            currentPath.add(neighbor);
            dfs(neighbor, target, graph, currentPath, result);
            // Бэктрекинг: удаляем последний добавленный узел для исследования других ветвей
            currentPath.remove(currentPath.size() - 1);
        }
    }
}
