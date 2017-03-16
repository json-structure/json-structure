package org.jsonstructure.jackson.validator.loanword;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

public class Result<T,E> {

    @Nullable
    private final T ok;

    @Nullable
    private final E err;

    public static<T,E> Result<T, E> ok(@Nullable T value) {
        return new Result<>(value, null);
    }

    public static<T,E> Result<T, E> err(@Nonnull E error) {
        return new Result<>(null, error);
    }


    Result(@Nullable T ok, @Nullable E err) {
        assert((ok == null) || (err == null));
        this.ok = ok;
        this.err = err;
    }

    @Nullable
    public T getOk() {
        return ok;
    }

    @Nullable
    public E getErr() {
        return err;
    }

    public boolean isOK() {
        return err == null;
    }

    public boolean isError() {
        return err != null;
    }

}
