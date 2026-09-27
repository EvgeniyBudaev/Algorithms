
import java.util.StringJoiner;


public class Training {
    public static void main(String[] args) {
        StringJoiner joiner = new StringJoiner(".", "prefix-", "-suffix");
        for (String s : "Hello world".split(" ")) {
            joiner.add(s);
        }
        System.out.println(joiner); // prefix-Hello.world-suffix
    }
}
