
public class Training {
    public static void main(String[] args) {
        Rodent capybara = new Capybara();
        capybara.crunch();
        capybara.sleep();
    }
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