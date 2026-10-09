import java.util.ArrayList;
import java.util.List;
import java.util.Objects;
import java.util.stream.Collectors;

public class Training {
    public static void main(String[] args) {
        List<String> list = new ArrayList<>();
        list.add("1");
        list.add("2");
        list.add(null);

        List<String> result = list.stream()
        .filter(Objects::nonNull)
        .collect(Collectors.toList());

        for (String el : result) {
            System.out.println(el);
        }
    }
}
