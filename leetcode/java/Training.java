import java.util.stream.Stream;

public class Training {
    public static void main(String[] args) {
        Stream.of(5, 3, 4).mapToInt(i -> i).max().getAsInt();
    }
}
