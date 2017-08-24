package tracker;

import static spark.Spark.*;

import tracker.idleMonitoring.IdleTimer;

/**
 * Class tracker.App
 *
 * @author Axel Nilsson (axnion)
 */
public class App {
    private IdleTimer idleTimer;


    public static void main(String[] args) {
        get("/start", (req, res) -> {
            return "";
        });



        IdleTimer timer = new IdleTimer(new Entry());
        timer.run();
    }

    private void startTimer() {

    }

    private void stopTimer() {

    }
}
