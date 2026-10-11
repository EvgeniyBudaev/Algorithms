
public class Training {
    public static void main(String[] args) {
        B b = new B();
        A a = new B();

        b.f(); // A
        a.f(); // A
    }
}


class A {
    public void f() {
        System.out.println("A");
    }
}

class B extends A {
    // @Override
    // public void f() {
    //     System.out.println("B");
    // }
}