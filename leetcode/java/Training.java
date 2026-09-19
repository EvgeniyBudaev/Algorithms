
public class Training {
    public static void main(String[] args) {
        String prefix = "Hello";

        Greeting task = () -> {
            System.out.println(prefix);
        };

        task.sayHello();
    }

    @FunctionalInterface
    interface Greeting {
        void sayHello();
    }
}
