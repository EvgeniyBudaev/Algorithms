import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Training {
    public static void main(String[] args) {
        Candy c1 = new Candy("a");
        Candy c2 = new Candy("c");
        Candy c3 = new Candy("b");
        List<Candy> candies = new ArrayList<>(List.of(c1, c2, c3));
        boolean t = candies.remove(c3);
        System.out.println("res: " + t);

        System.out.println("Сортировка по имени");
        candies.sort(Candy::compareByName);
        candies.forEach(i -> System.out.println(i.name));

        Map<String, Integer> map = new HashMap<>();
        map.put("5", 5);
        boolean t2 = map.remove("5", 5);
        System.out.println("res: " + t2);
    }
}

class Candy {
    public String name;

    public Candy (String name) {
        this.name = name;
    }

    public static int compareByName(Candy c1, Candy c2) {
        return c1.name.compareTo(c2.name);
    } 
}
