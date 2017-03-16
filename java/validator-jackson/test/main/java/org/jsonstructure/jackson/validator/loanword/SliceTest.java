package org.jsonstructure.jackson.validator.loanword;

import java.util.Arrays;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class SliceTest {

    @Test
    public void build() {
        Slice<Integer> s0 = Slice.empty();
        Slice<Integer> s1 = s0.append(1, 2, 3, 4, 5);
        Slice<Integer> s2 = s1.append(6, 7);
        assertEquals(Arrays.asList(1, 2, 3, 4, 5), s1.toList());
        assertEquals(Arrays.asList(1, 2, 3, 4, 5, 6, 7), s2.toList());
        Slice<Integer> s3 = s1.append(8, 9, 10);
        assertEquals(Arrays.asList(1, 2, 3, 4, 5), s1.toList());
        assertEquals(Arrays.asList(1, 2, 3, 4, 5, 8, 9), s2.toList());
        assertEquals(Arrays.asList(1, 2, 3, 4, 5, 8, 9, 10), s3.toList());
    }
}
