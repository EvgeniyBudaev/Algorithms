public class DynamicPolymorphism {
    public static void main(String[] args) {
        Actor actor = new Hamlet();
        actor.play(); // Актёр вышел на сцену и начал монолог: "Быть или не быть, вот в чём вопрос."
        System.out.println(actor.sayLine()); // "Быть или не быть, вот в чём вопрос."
    }
}

class Actor {

    public void play() {
        System.out.println("Актёр вышел на сцену и начал монолог: \"" + this.sayLine() + "\"");
    }

    public String sayLine() {
        return "Провала прошу, провала.";
    }

}
class Hamlet extends Actor {

    @Override
    public String sayLine() {
        return "Быть или не быть, вот в чём вопрос.";
    }

}
