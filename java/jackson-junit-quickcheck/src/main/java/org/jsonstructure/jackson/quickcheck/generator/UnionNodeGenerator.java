package org.jsonstructure.jackson.quickcheck.generator;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

import javax.annotation.Nonnull;

import com.fasterxml.jackson.databind.JsonNode;
import com.pholser.junit.quickcheck.generator.GenerationStatus;
import com.pholser.junit.quickcheck.generator.Generator;
import com.pholser.junit.quickcheck.internal.Weighted;
import com.pholser.junit.quickcheck.internal.generator.CompositeGenerator;
import com.pholser.junit.quickcheck.random.SourceOfRandomness;

import org.jsonstructure.jackson.quickcheck.JsonNodeGenerator;
import org.jsonstructure.jackson.validator.Structure;
import org.jsonstructure.jackson.validator.TypeDecl;

public class UnionNodeGenerator extends Generator<JsonNode> {

    @Nonnull
    final CompositeGenerator gen;

    private UnionNodeGenerator(@Nonnull CompositeGenerator gen) {
        super(JsonNode.class);
        this.gen = gen;
    }

    public static UnionNodeGenerator create(@Nonnull Structure structure,
                                            @Nonnull TypeDecl typeDecl) {
        List<Weighted<Generator<?>>> composed = new ArrayList<>();
        for (Map.Entry<String, TypeDecl> entry : typeDecl.types.entrySet()) {
            Generator<?> typeGen = JsonNodeGenerator.create(structure, entry.getValue());
            Weighted<Generator<?>> weighted = new Weighted<>(typeGen, 1);
            composed.add(weighted);
        }
        CompositeGenerator gen = new CompositeGenerator(composed);
        return new UnionNodeGenerator(gen);
    }

    @Override
    public JsonNode generate(SourceOfRandomness random, GenerationStatus status) {
        return (JsonNode) gen.generate(random, status);
    }

    @Override
    public List<JsonNode> doShrink(SourceOfRandomness random, JsonNode larger) {
        return (List<JsonNode>) ((List<?>) gen.doShrink(random, larger));
    }
}
