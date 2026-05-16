package leetcode.java.intervals.PointsAndSegments;

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.StringTokenizer;

/* Точки и отрезки

Дано n отрезков на числовой прямой и m точек на этой же прямой. Для каждой из данных точек определите,
скольким отрезкам они принадлежат. Точка x считается принадлежащей отрезку с концами a и b, если выполняется двойное
неравенство min (a, b) ≤ x ≤ max (a, b) .

Входные данные
Первая строка содержит два целых числа n (1≤ n ≤ 10 в степени 5 ) – число отрезков и m (1≤ m ≤10 в степени 5 ) – число точек.
В следующих n строках по два целых числи a i и b i – координаты концов соответствующего отрезка(например, может быть
пара 5 2). В последней строке m целых чисел – координаты точек. Все числа по модулю не превосходят 10 9

Выходные данные
В выходной файл выведите m чисел – для каждой точки количество отрезков, в которых она содержится.

Example 1:
Input:
3 2
0 5
-3 2
7 10
1 6

Output: 2 0
*/

public class PointsAndSegments {
    // time: O((n+m)*log(n+m)), space: O(n+m)
    public static void main(String[] args) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter writer = new BufferedWriter(new OutputStreamWriter(System.out));

        StringTokenizer st = new StringTokenizer(reader.readLine());
        int n = Integer.parseInt(st.nextToken()); // количество отрезков
        int m = Integer.parseInt(st.nextToken()); // количество точек

        List<Point> points = new ArrayList<>(2 * n + m);

        // Читаем отрезки
        for (int i = 0; i < n; i++) {
            st = new StringTokenizer(reader.readLine());
            int a = Integer.parseInt(st.nextToken());
            int b = Integer.parseInt(st.nextToken());

            // Сортируем концы отрезка
            if (a > b) {
                int tmp = a;
                a = b;
                b = tmp;
            }
            points.add(new Point(a, 1, -1)); // начало отрезка (+1)
            points.add(new Point(b, -1, -1)); // конец отрезка (-1)
        }

        // Читаем точки-запросы
        int[] inPoints = new int[m];
        for (int i = 0; i < m; i++) {
            inPoints[i] = Integer.parseInt(reader.readLine().trim());
            points.add(new Point(inPoints[i], 0, i)); // запрос (typ = 0)
        }

        // Сортируем точки и отрезки:
        // 1. По координате возрастания
        // 2. При равенстве координат: начала (+1) → запросы (0) → концы (-1)
        Collections.sort(points, (p1, p2) -> {
            if (p1.coord == p2.coord) {
                // Порядок: +1 (start) > 0 (query) > -1 (end)
                return Integer.compare(p2.typ, p1.typ);
            }
            return Integer.compare(p1.coord, p2.coord);
        });

        int[] result = new int[m]; // результаты для каждой точки-запроса
        int current = 0;           // текущее количество активных отрезков

        // Обрабатываем события в порядке сортировки
        for (Point p : points) {
            if (p.typ == 0) {
                // Это точка-запрос: записываем текущее количество отрезков
                result[p.index] = current;
            } else {
                // Это начало или конец отрезка: обновляем счётчик
                current += p.typ;
            }
        }

        // Выводим результаты
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < m; i++) {
            if (i > 0) sb.append(' ');
            sb.append(result[i]);
        }
        sb.append('\n');
        writer.write(sb.toString());

        writer.flush();
        reader.close();
        writer.close();
    }

    // Point - класс для хранения точек и отрезков.
    private static class Point {
        int coord; // координата
        int typ;   // тип точки: +1 для начала отрезка, -1 для конца, 0 для запроса
        int index; // индекс точки (для запросов)

        Point(int coord, int typ, int index) {
            this.coord = coord;
            this.typ = typ;
            this.index = index;
        }
    }
}
