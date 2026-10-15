import java.util.ArrayList;
import java.util.List;

public class Training {
    public static void main(String[] args) {
        Apple a = new Apple();
        Banana b = new Banana();
        List<Fruit> apples = new ArrayList<>();
        apples.add(a);
        List<Fruit> bananes = new ArrayList<>();
        bananes.add(b);
        Bag<Fruit> bag = new Bag<>(apples);
        bag.print();
    }
}

class Bag<T extends Fruit> {
    private final List<T> list;

    public Bag(List<T> list) {
        this.list = list;
    }

    public void print() {
        for (T el : list) {
            System.out.println(el.getName());
        }
    }
}

abstract class Fruit {
    private final String name;
    private final Double price;

    protected Fruit(String name, Double price) {
        this.name = name;
        this.price = price;
    }

    public String getName() {
        return name;
    }

    public Double getPrice() {
        return price;
    }
}

class Apple extends Fruit {
    public Apple() {
        super("Яблоко", 10.0);
    }
}

class Banana extends Fruit {
    public Banana() {
        super("Банан", 15.50);
    }
}