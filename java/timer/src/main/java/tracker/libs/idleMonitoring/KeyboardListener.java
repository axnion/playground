package tracker.libs.idleMonitoring;

import org.jnativehook.keyboard.NativeKeyEvent;
import org.jnativehook.keyboard.NativeKeyListener;

/**
 * Class KeyboardListener
 *
 * @author Axel Nilsson (axnion)
 */
public class KeyboardListener implements NativeKeyListener {
    private IdleTimer timer;

    public KeyboardListener(IdleTimer timer) {
        this.timer = timer;
    }

    public void nativeKeyPressed(NativeKeyEvent e) {
        timer.reset();
    }

    public void nativeKeyReleased(NativeKeyEvent e) {
        timer.reset();
    }

    public void nativeKeyTyped(NativeKeyEvent e) {
        timer.reset();
    }
}
