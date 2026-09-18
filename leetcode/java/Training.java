public class Training {
    public static void main(String[] args) {
        new Training().go();
    }

    void go() {
        Mammal m = new Zebra();
        System.out.println(m.name + " " + m.makeNoise() + " " + m.foo());

        Zebra z = new Zebra();
        System.out.println(z.name + " " + z.makeNoise() + " " + z.foo());
    }

    class Mammal {
        String name = "furry";

        String makeNoise() {
            return "generic noise";
        }

        static String foo() {
            return "foo";
        }
    }

    class Zebra extends Mammal {
        String name = "stripes";

        String makeNoise() {
            return "bray";
        }

        static String foo() {
            return "baz";
        }
    }

}
