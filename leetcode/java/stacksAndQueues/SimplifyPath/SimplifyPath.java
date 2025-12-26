package stacksAndQueues.SimplifyPath;

import java.util.ArrayDeque;
import java.util.Deque;

/* 71. Simplify Path
https://leetcode.com/problems/simplify-path/description/

Учитывая абсолютный путь к файловой системе в стиле Unix, который начинается с косой черты «/», преобразуйте этот путь в
упрощенный канонический путь.

В контексте файловой системы в стиле Unix одна точка '.' означает текущий каталог, двойная точка «..» означает переход
на один уровень каталога вверх, а несколько косых черт, таких как «//», интерпретируются как одна косая черта.
В этой задаче рассматривайте последовательности точек, не подпадающие под действие предыдущих правил (например, «...»),
как допустимые имена файлов или каталогов.

Упрощенный канонический путь должен соответствовать следующим правилам:

Он должен начинаться с одинарной косой черты «/».
Каталоги внутри пути должны быть разделены только одной косой чертой «/».
Он не должен заканчиваться косой чертой «/», если только это не корневой каталог.
Следует исключить любые одиночные или двойные точки, используемые для обозначения текущих или родительских каталогов.
Верните новый путь.

Input: path = "/home/"
Output: "/home"

Input: path = "/home//foo/"
Output: "/home/foo"

Input: path = "/home/user/Documents/../Pictures"
Output: "/home/user/Pictures"

Input: path = "/../"
Output: "/"

Input: path = "/.../a/../b/c/../d/./"
Output: "/.../b/d"
*/

public class SimplifyPath {
    public static void main(String[] args) {
        System.out.println(simplifyPath("/home/"));       // "/home"
        System.out.println(simplifyPath("/home//foo/"));  // "/home/foo"
    }

    // simplifyPath принимает абсолютный путь в стиле Unix и возвращает упрощенный канонический путь.
    // time: O(n), space: O(n), где n - длина пути
    private static String simplifyPath(String path) {
        Deque<String> stack = new ArrayDeque<>(); // Стек для хранения директорий
        String[] directories = path.split("/"); // Разбиваем путь на составляющие директории

        for (String dir : directories) {
            if (dir.equals(".") || dir.isEmpty()) { // Пропускаем текущий каталог и пустые строки
                continue;
            } else if (dir.equals("..")) { // Обрабатываем переход на один уровень каталога вверх
                if (!stack.isEmpty()) {
                    stack.pollLast(); // Удаляем последний элемент
                }
            } else {
                // Добавляем обычную директорию
                stack.offerLast(dir);
            }
        }

        return "/" + String.join("/", stack); // Собираем упрощенный путь из элементов стека
    }
}
