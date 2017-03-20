package org.jsonstructure.jackson.validator.loanword;

import java.util.Arrays;
import java.util.List;

import javax.annotation.Nonnegative;
import javax.annotation.Nonnull;

public class Slice<T> {

    @Nonnull
    private final Object[] data;

    @Nonnegative
    private final int length;

    public static <T> Slice<T> empty() {
        return new Slice<>();
    }

    public static <T> Slice<T> create(Object... elements) {
        return new Slice<>(elements, elements.length);
    }

    Slice() {
        this.data = new Object[0];
        this.length = 0;
    }

    Slice(@Nonnull Object[] data, @Nonnegative int length) {
        this.data = data;
        this.length = length;
    }

    public Slice<T> append(Object... elements) {
        Object[] newData = data;
        int total = length + elements.length;
        if (total > data.length) {
            int newSize = (total * 3) / 2 + 1;
            newData = new Object[newSize];
            System.arraycopy(data, 0, newData, 0, length);
        }
        System.arraycopy(elements, 0, newData, length, elements.length);
        return new Slice<>(newData, total);
    }

    public List<T> toList() {
        return (List<T>) Arrays.asList(data).subList(0, length);
    }

}
