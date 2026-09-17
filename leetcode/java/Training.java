import java.util.HashMap;
import java.util.Map;
import java.util.TreeMap;

public class Training {
    public static void main(String[] args) {
        Map<Object, Object> map = new HashMap<>();
        map.put(new Object(), new Object());
        map.put(new Object(), new Object());
        TreeMap<Object, Object> treeMap = new TreeMap<>();
        treeMap.putAll(map);
        System.out.println(treeMap.size());
    }
}

