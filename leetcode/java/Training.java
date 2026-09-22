

public class Training {
    public static void main(String[] args) {
        Training t1 = new Training();
        Training t2 = new Training();
        Training r = t1.go(t2);
        System.out.println("t2 before: " + t2); // Training@2a139a55
        System.out.println("r: " + r); // null
        System.out.println("t2 after: " + t2); // Training@2a139a55
    }

    public Training go(Training tr) {
        tr = null;
        return tr;
    }
}
