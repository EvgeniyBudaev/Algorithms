

public class Training {
    public static void main(String[] args) {
        Training t1 = new Training();
        Training t2 = new Training();
        Training r = t1.go(t2);
        System.out.println(r); // null
    }

    public Training go(Training tr) {
        tr = null;
        return tr;
    }
}
