

public class Training {
    public static void main(String[] args) {
        Float wrapper = null;
        float primitive = wrapper; // NullPointerException
        System.out.println(primitive); 
    }
}
