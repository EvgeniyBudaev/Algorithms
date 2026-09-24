import java.util.Comparator;
import java.util.Set;
import java.util.TreeSet;

public class Training {
    public static void main(String[] args) {
        Comparator<Film> comparator = (i1, i2) -> i1.rating - i2.rating;
        Set<Film> films = new TreeSet<>(comparator);

        for (Film film : films) {
            System.out.println(film);
        }
    }

    public class Film {
        public String title;
        public String directorName;
        public int rating;
    } 
}
