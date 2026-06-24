public class Animals {
    public static void main(String[] args) {
        Pet pet = new Cat();
        pet.play();
        pet.eat();
        // Такого метода в интерфейсе Pet нет - произойдет ошибка
        //pet.giveMilk();

        Mammal mammal = new Cat();
        mammal.giveMilk();
        mammal.eat();
        // Такого метода в интерфейсе Mammal нет - произойдет ошибка
        //mammal.play();

        Cat cat = new Cat();
        cat.play();

        Rodent capybara = new Capybara();
        capybara.crunch();
        capybara.sleep();
        // Такого метода в Rodent нет - произойдет ошибка
        //capybara.dive();
    }
}

class Cat implements Pet, Mammal {
    @Override
    public int getPawsCount() {
        return 4;
    }

    @Override
    public void makeNoise() {
        System.out.println("Мяу!");
    }

    @Override
    public void play() {
        System.out.println("Играю с мячиком.");
    }

    // Метод будет переопределять и метод из интерфейса Pet, и метод из интерфейса Mammal
    @Override
    public void eat() {
        System.out.println("Люблю покушать рыбку.");
    }

    @Override
    public void giveMilk() {
        System.out.println("Кормлю котят молоком.");
    }
}

interface Mammal {
    // У интерфейса Mammal есть аналогичный метод
    void eat();
    void giveMilk();
}

interface Pet {
    int getPawsCount();
    void makeNoise();
    void play();
    //Добавляем метод eat()
    void eat();
}

class Rodent {
    public void crunch() {
        System.out.println("Грызун грызёт что угодно.");
    }

    public void sleep() {
        System.out.println("Грызун уснул.");
    }
}

class Capybara extends Rodent {

    @Override
    public void crunch() {
        System.out.println("Капибара любит грызть кукурузу.");
    }

    @Override
    public void sleep() {
        System.out.println("Капибара крепко спит.");
    }

    public void dive() {
        System.out.println("Капибара умеет нырять.");
    }
}
