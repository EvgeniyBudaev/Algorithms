import java.util.LinkedHashMap;
import java.util.Map;

public class Training {
    public static void main(String[] args) {
        char result = findFirstUniqueLetter("summer time sadness");
        System.out.println(result); // 'u'
    }

    private static Character findFirstUniqueLetter(String str) {
        if (str == null || str.isEmpty()) {
            throw new IllegalArgumentException("String is null or empty");
        }

        Map<Character, Integer> freq = new LinkedHashMap<>();
        for (char c : str.toCharArray()) {
            freq.put(c, freq.getOrDefault(c, 0) + 1);
        }

        for (Map.Entry<Character, Integer> entry : freq.entrySet()) {
            if (entry.getValue() == 1) {
                return entry.getKey();
            }
        }

        throw new RuntimeException("No unique character found");
    }
}
