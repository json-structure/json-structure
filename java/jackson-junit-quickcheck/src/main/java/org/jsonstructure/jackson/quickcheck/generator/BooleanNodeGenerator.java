package org.jsonstructure.jackson.quickcheck.generator;

import java.util.ArrayList;
import java.util.List;

import javax.annotation.Nonnull;

import com.fasterxml.jackson.databind.node.BooleanNode;
import com.pholser.junit.quickcheck.generator.GenerationStatus;
import com.pholser.junit.quickcheck.generator.Generator;
import com.pholser.junit.quickcheck.generator.java.lang.BooleanGenerator;
import com.pholser.junit.quickcheck.random.SourceOfRandomness;

public class BooleanNodeGenerator extends Generator<BooleanNode> {

    @Nonnull
    final BooleanGenerator gen;

    private BooleanNodeGenerator() {
        super(BooleanNode.class);
        this.gen = new BooleanGenerator();
    }

    public static BooleanNodeGenerator create() {
        return new BooleanNodeGenerator();
    }

    @Override
    public BooleanNode generate(SourceOfRandomness random, GenerationStatus status) {
        return BooleanNode.valueOf(gen.generate(random, status));
    }

    @Override
    public List<BooleanNode> doShrink(SourceOfRandomness random, BooleanNode larger) {
        List<Boolean> input = gen.doShrink(random, larger.booleanValue());
        List<BooleanNode> output = new ArrayList<>(input.size());
        input.forEach(x -> output.add(BooleanNode.valueOf(x)));
        return output;
    }

}
