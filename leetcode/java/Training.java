import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class Training {
    public static void main(String[] args) {
        List<String> list = new ArrayList<>();
        list.add("D");
        list.add("A");
        list.add("B");
        Collections.sort(list, (a, b) -> a.compareTo(b));
        System.out.println(list);
    }
}
