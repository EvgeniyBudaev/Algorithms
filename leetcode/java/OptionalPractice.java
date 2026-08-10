import java.util.Collection;
import java.util.List;
import java.util.Optional;
import java.util.Set;

public class OptionalPractice {
    public static void main(String[] args) {
        /* Методы для создания Optional-объекта */
        // Метод Optional.of
        Optional<Integer> optionalNumber = Optional.of(123); // Передаваемое значение не должно быть null
        System.out.println(optionalNumber); // Optional[123]

        // Метод Optional.ofNullable
        String someString = "Hello, world";
        Optional<String> optionalString = Optional.ofNullable(someString);
        System.out.println(optionalString); // Optional[Hello, world]
        String someStringNullable = null; // Значение может прийти из другого метода и равняться null
        Optional<String> optionalStringNullable = Optional.ofNullable(someStringNullable);
        System.out.println(optionalStringNullable); // Optional.empty

        // Метод Optional.empty
        Collection<Integer> numbers = List.of(1, 2, 3);
        System.out.println(findFirstGreaterThan(1, numbers)); // Optional[2]
        System.out.println(findFirstGreaterThan(5, numbers)); // Optional.empty

        /* Методы для чтения Optional */
        // Методы isPresent и isEmpty
        Optional<Double> optionalPrice = Optional.of(123.4);
        System.out.println(optionalPrice.isPresent()); // true
        System.out.println(optionalPrice.isEmpty()); // false

        // Метод get
        Optional<Integer> optNumber = findFirstGreaterThan(5, Set.of(1, 2, 3, 4));
        // Проверяем, что объект класса Optional содержит значение.
        // Без проверки будет NoSuchElementException
        if (optNumber.isPresent()) {
            // вызываем метод get, чтобы получить значение, содержащееся в Optional
            System.out.println("Найденное число равно: " + optNumber.get());
        }

        // Методы orElseThrow, orElse и orElseGet
        // Вызываем метод, возвращающий Optional<Integer>, и затем распаковываем его
        // Если Optional пуст, выкидываем специально созданное исключение,
        // не обобщённое NoSuchElementException, которое генерируется при использовании get
        Integer number1 = findFirstGreaterThan(5, Set.of(1, 2, 3, 4))
                .orElseThrow(() -> new RuntimeException("Число больше 5 не найдено"));
        // Вызываем метод, возвращающий Optional<Integer> и затем распаковываем его
        // Если Optional пуст, возвращаем значение по умолчанию
        Integer number2 = findFirstGreaterThan(5, Set.of(1, 2, 3, 4))
                .orElse(100500);
        // Вызываем метод, возвращающий Optional<Integer> и затем распаковываем его
        // Если Optional пуст, то будет вычислено значение по умолчанию в методе 
        // someHeavyComputation(), но мы не хотим, чтобы этот метод запустился сразу (что
        // произошло бы, если бы использовался orElse), а только если Optional пуст
        // Оборачиваем его в лямбду, которая будет вызвана только после проверки
        Integer number3 = findFirstGreaterThan(5, Set.of(1, 2, 3, 4))
                .orElseGet(() -> someHeavyComputation());

        // ifPresent
        // Вызываем метод, возвращающий Optional<Integer>, и передаём лямбду
        // Лямбду нужно выполнить, если Optional содержит значение
        // В качестве лямбды используем ссылку на метод, который выводит значение на экран
        findFirstGreaterThan(5, Set.of(1, 2, 3, 4))
                .ifPresent(System.out::println);

        // ifPresentOrElse
        // Вызываем метод, возвращающий Optional<Integer>, и передаём две лямбды
        // Первая выполнится в случае, если Optional не пуст,
        //  вторая — если пуст
        findFirstGreaterThan(5, Set.of(1, 2, 3, 4))
                .ifPresentOrElse(
                        number -> System.out.println("Найденное число равно " + number),
                        () -> System.out.println("Число не найдено")
                );
    }

    public static Optional<Integer> findFirstGreaterThan(int n, Collection<Integer> numbers) {
        for (Integer num : numbers) {
            if (num > n) return Optional.of(num); // число больше n найдено
        }
        return Optional.empty(); // в переданных числах, нет числа больше чем n
    }

    public static Integer someHeavyComputation() {
        // Имитация тяжёлых вычислений
        return 42;
    }
}
