package tracker.libs;

/**
 * Class tracker.libs.Entry
 *
 * @author Axel Nilsson (axnion)
 */
public class Entry {
    private long startTime;
    private long idleStartTime;
    private boolean idle;

    public void start() {
        startTime = System.currentTimeMillis();
    }

    public void setIdleState(boolean state) {
        idle = state;
    }
}
