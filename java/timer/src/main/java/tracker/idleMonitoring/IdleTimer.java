package tracker.idleMonitoring;

import org.jnativehook.GlobalScreen;
import org.jnativehook.NativeHookException;
import tracker.Entry;

import java.util.logging.Level;
import java.util.logging.Logger;

/**
 * Class IdleTimer
 *
 * @author Axel Nilsson (axnion)
 */
public class IdleTimer implements Runnable {
    private long startTime;
    private KeyboardListener keyboardListener;
    private boolean idle;
    private Entry entry;

    public IdleTimer(Entry entry) {
        keyboardListener = new KeyboardListener(this);
        startTime = System.currentTimeMillis();
        idle = false;
        this.entry = entry;

        Logger logger = Logger.getLogger(GlobalScreen.class.getPackage().getName());
        logger.setLevel(Level.WARNING);

        try {
            GlobalScreen.registerNativeHook();
        }
        catch(NativeHookException expt) {
            System.err.println("There was a problem registering the native hook.");
            System.err.println(expt.getMessage());

            System.exit(1);
        }

        GlobalScreen.addNativeKeyListener(keyboardListener);
    }

    public void run() {
        entry.start();
        while(true) {
            if (elapsedTime() > 10000 && !idle) {
                System.out.println("OMG HE'S GONE!");
                idle = true;
            }

            try {
                Thread.sleep(1000);
            }
            catch(InterruptedException expt) {
                System.exit(1);
            }
        }
    }


    void reset() {
        if (idle) {
            System.out.println("Fuck, you are back...");
        }

        startTime = System.currentTimeMillis();
        idle = false;
    }

    private long elapsedTime() {
        return System.currentTimeMillis() - startTime;
    }

    private void notifyEntry() {
        entry.setIdleState(idle);
    }
}
