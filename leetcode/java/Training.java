import java.io.BufferedWriter;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.time.Duration;

public class Training {
    public static void main(String[] args) throws InterruptedException {
        SumThread thread = new SumThread();
        thread.start();

        Thread.sleep(Duration.ofSeconds(1));
        thread.interrupt(); 
    }
}

class SumThread extends Thread {

    @Override
    public void run() {
        try (BufferedWriter writer = Files.newBufferedWriter(Path.of("result.txt"))) {
            long total = 0;
            for (int i = 0; i < 1_000_000_000; i++) {
                total += i;

                if (i % 100 == 0) { // Каждые 100 итераций записываем результат в файл
                    writer.write(total + "\n");
                    writer.flush();

                    if (isInterrupted()) {
                        // Если поток прервали, завершаем работу
                        break;
                    }
                }
            }
        } catch (IOException e) {
            throw new RuntimeException(e.getMessage(), e);
        }
    }

} 
