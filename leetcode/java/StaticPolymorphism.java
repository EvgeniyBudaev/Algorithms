public class StaticPolymorphism {
    public static void main(String[] args) {
        Printer printer = new Printer();
        // Будет выбран метод, принимающий тип int
        printer.print(10);
        // Будет выбран метод, принимающий тип String
        printer.print("Привет!");
        // Будет выбран метод, принимающий тип Object
        printer.print(printer);
    }
}

class Printer {

    public void print(int i) {
        System.out.println(i + " - это число.");
    }

    public void print(String s) {
        System.out.println(s + " - это строка.");
    }

    public void print(Object o) {
        System.out.println(o.toString() + " - это объект.");
    }

    @Override
    public String toString() {
        return "Я - объект принтера!";
    }
}
