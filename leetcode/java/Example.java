import java.util.HashMap;

class Example {
    public static void main(String[] args) {
        String name = "Катя";

        System.out.println("Что " + name + " делает как работник:");
        Employee employee = new SoftwareDeveloper(name);
        // employee.work();
        // System.out.println("Получает зарплату за месяц: " + employee.getSalary() + " р.");
        // Таких методов в интерфейсе Employee нет - произойдет ошибка
        // employee.eat();
        // employee.doDebug();

        System.out.println("\nЧто " + name + " делает как человек:");
        Person person = new SoftwareDeveloper(name);
        // person.eat();
        // person.sleep();
        person.sayLine();
        // Таких методов в классе Person нет - произойдет ошибка
        // person.work();
        // person.doDebug();
        person.callStatic();

        System.out.println("\nЧто " + name + " делает как программист:");
        SoftwareDeveloper softwareDeveloper = new SoftwareDeveloper(name);
        // softwareDeveloper.eat();
        // softwareDeveloper.work();
        // softwareDeveloper.getSalary();
        // softwareDeveloper.doDebug();
        // softwareDeveloper.think();
        softwareDeveloper.sayLine();
        softwareDeveloper.callStatic();

        HashMap<String, String> officeTool = new HashMap<>();
        officeTool.put("S234", "Большой степлер");
         officeTool.remove("P34342"); // нет ошибки
         System.out.println(officeTool);
    }
}


interface Employee {
    void work();

    double getSalary();
}

class Person {
    private final String name;

    public Person(String name) {
        this.name = name;
    }

    public void eat() {
        System.out.println("Обедает в кафе.");
    }

    public void sleep() {
        System.out.println("Крепко спит всю ночь.");
    }

    public void sayLine() {
        System.out.println("Person sayLine");
    }

    public static void callStatic() {
        System.out.println("Person callStatic");
    }
}

class SoftwareDeveloper extends Person implements Employee {
    public static final double PI = 3.14;

    public SoftwareDeveloper(String name) {
        super(name);
    }

    @Override
    public void work() {
        System.out.println("Выполняет свою работу - пишет код весь день.");
    }

    @Override
    public double getSalary() {
        return 100500.0;
    }

    public void doDebug() {
        System.out.println("Ищет ошибки в коде.");
    }

    public void think() {
        System.out.println("Обдумывает решение задачи.");
    }

    @Override
    public void sayLine() {
        System.out.println("SoftwareDeveloper sayLine");
    }

    public static void callStatic() {
        System.out.println("SoftwareDeveloper callStatic");
    }
}
