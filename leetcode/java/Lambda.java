import java.nio.file.Paths;
import java.time.LocalDateTime;
import java.util.function.Consumer;
import java.util.function.Function;
import java.util.function.Predicate;
import java.util.function.Supplier;

public class Lambda {
    public static void main(String[] args) {
        // Создаём реализацию FileNameTransformer с помощью лямбда-функции 
        FileNameTransformer transformer = filePath -> Paths.get(filePath).getFileName().toString();
        // Вызываем сохранённый экземпляр лямбда-функции
        String result = transformer.getFileName("/home/bigbrother/downloads/movie.mp4");
        System.out.println(result);

        // Function<T, R> с методом R apply(T t)
        Function<Integer, String> intToString = num -> String.valueOf(num);
        System.out.println(intToString.apply(1000_0000));

        // Predicate<T> с методом boolean test(T t)
        Predicate<Integer> isEven = num -> num % 2 == 0;
        if (isEven.test(12345)) {
            System.out.println("Число чётное");
        } else {
            System.out.println("Число нечётное");
        }

        // Consumer<T> с методом void accept(T t)
        Consumer<Double> outputDoubleConsumer =
                num -> System.out.println(String.format("Передано число %.2f", num));
        outputDoubleConsumer.accept(1234.5678);

        // Supplier<T> с методом T get()
        Supplier<LocalDateTime> currentDateTimeSupplier = () -> LocalDateTime.now();
        System.out.println("Текущая дата и время: " + currentDateTimeSupplier.get());
    }
}

@FunctionalInterface
interface FileNameTransformer {
    String getFileName(String filePath);
} 
