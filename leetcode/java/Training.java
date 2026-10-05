
public class Training {
    public static void main(String[] args) {
        A a = new B();
        a.F();
    }
}

    class A {
        public static void F() {
            System.out.println("A");
        };
    }

    class B extends A {
        public static void F() {
            System.out.println("B");
        };
    }
