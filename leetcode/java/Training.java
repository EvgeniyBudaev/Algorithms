


public class Training {
    private static final String world = "world";

    public static void main(String[] args) {
        StaticTest st = new StaticTest();
        st.foo();
    }

    public static class StaticTest {
        public void foo() {
            System.out.println(world);
        }
    }
}
