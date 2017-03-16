package org.jsonstructure.jackson.validator.loanword;

import java.util.concurrent.CountDownLatch;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.function.Supplier;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

public class Once<T> {

    @Nonnull
    private final AtomicBoolean started;

    @Nonnull
    private final CountDownLatch finished;

    @Nonnull
    private final Supplier<T> supplier;

    private volatile T result;

    private volatile RuntimeException exception;

    public Once(@Nonnull Supplier<T> supplier) {
        this.started = new AtomicBoolean();
        this.finished = new CountDownLatch(1);
        this.supplier = supplier;
    }

    private boolean isFirst() {
        return started.compareAndSet(false, true);
    }

    private void taskComplete() {
        finished.countDown();
    }

    private void await() {
        boolean interrupted = false;
        try {
            while (true) {
                try {
                    finished.await();
                    return;
                } catch (InterruptedException e) {
                    interrupted = true;
                }
            }
        } finally {
            if (interrupted) {
                Thread.currentThread().interrupt();
            }
        }
    }

    @Nullable
    public T value() {
        boolean first = isFirst();
        if (first) {
            try {
                result = supplier.get();
            } catch (RuntimeException ex) {
                exception = ex;
            } finally {
                taskComplete();
            }
        } else {
            await();
        }
        if (exception != null) {
            throw exception;
        }
        return result;
    }

}
