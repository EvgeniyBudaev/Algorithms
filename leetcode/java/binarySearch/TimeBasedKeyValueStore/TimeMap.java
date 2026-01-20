package binarySearch.TimeBasedKeyValueStore;

import java.util.*;

class TimeMap {
    // Хранилище: ключ -> список пар (значение, временная метка)
    private Map<String, List<TimeValue>> store;

    public TimeMap() {
        store = new HashMap<>();
    }

    public void set(String key, String value, int timestamp) {
        store.putIfAbsent(key, new ArrayList<>());
        store.get(key).add(new TimeValue(value, timestamp));
    }

    public String get(String key, int timestamp) {
        List<TimeValue> arr = store.get(key);
        if (arr == null || arr.isEmpty()) {
            return "";
        }

        int low = 0;
        int high = arr.size() - 1;
        String res = "";

        while (low <= high) {
            int mid = (low + high) / 2;
            TimeValue tv = arr.get(mid);

            if (timestamp == tv.timestamp) {
                return tv.value;
            }

            if (timestamp >= tv.timestamp) {
                low = mid + 1;
                res = tv.value;
            } else {
                high = mid - 1;
            }
        }

        return res;
    }

    // Вспомогательный класс для хранения значения и временной метки
    private static class TimeValue {
        String value;
        int timestamp;

        TimeValue(String value, int timestamp) {
            this.value = value;
            this.timestamp = timestamp;
        }
    }
}