package org.jsonstructure.jackson.quickcheck.generator;

import java.lang.annotation.Annotation;

import com.pholser.junit.quickcheck.generator.Size;

public class SizeImpl implements Size {

    private final int min;

    private final int max;

    public SizeImpl(Integer min, Integer max) {
        this.min = (min == null) ? 0 : min;
        this.max = (max == null) ? Integer.MAX_VALUE : max;
    }

    @Override
    public int min() {
        return min;
    }

    @Override
    public int max() {
        return max;
    }

    @Override
    public Class<? extends Annotation> annotationType() {
        return null;
    }
}
