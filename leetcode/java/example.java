class Example {
    public static void main(String[] args) {
        // Java автоматически привела объект класса RussianPassport в типу интерфейса Document
        Document document = new RussianPassport("1111", "567899");
        System.out.println(document.getDocumentNumber());

        String name = "Катя";
        Employee employee = new SoftwareDeveloper(name);
        employee.work();
        employee.getSalary();

        Person person = new SoftwareDeveloper(name);
        person.eat();
        person.sleep();

        SoftwareDeveloper softwareDeveloper = new SoftwareDeveloper(name);
        softwareDeveloper.doDebug();
        softwareDeveloper.think();
        softwareDeveloper.eat();
        softwareDeveloper.sleep();
        softwareDeveloper.work();
        softwareDeveloper.getSalary();
    }
}

interface Document {
    public String getDocumentNumber();
}

class RussianPassport implements Document {
    private final String series;
    private final String number;

    public RussianPassport(String series, String number) {
        this.series = series;
        this.number = number;
    }

    @Override
    public String getDocumentNumber() {
        return series + " " + number;
    }

    public String getSeries() {
        return series;
    }

    public String getNumber() {
        return number;
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
}
class SoftwareDeveloper extends Person implements Employee {

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
}