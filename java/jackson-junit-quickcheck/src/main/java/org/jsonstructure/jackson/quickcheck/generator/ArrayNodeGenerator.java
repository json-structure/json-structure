package org.jsonstructure.jackson.quickcheck.generator;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

import javax.annotation.Nonnull;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.node.ArrayNode;
import com.fasterxml.jackson.databind.node.JsonNodeFactory;
import com.pholser.junit.quickcheck.generator.GenerationStatus;
import com.pholser.junit.quickcheck.generator.Generator;
import com.pholser.junit.quickcheck.generator.java.util.ArrayListGenerator;
import com.pholser.junit.quickcheck.random.SourceOfRandomness;

import org.jsonstructure.jackson.quickcheck.JsonNodeGenerator;
import org.jsonstructure.jackson.validator.Structure;
import org.jsonstructure.jackson.validator.TypeDecl;

public class ArrayNodeGenerator extends Generator<ArrayNode> {

    @Nonnull
    final ArrayListGenerator gen;

    @Nonnull
    private final JsonNodeFactory factory;

    @Nonnull
    final TypeDecl typeDecl;

    private ArrayNodeGenerator(@Nonnull ArrayListGenerator gen,
                               @Nonnull JsonNodeFactory factory,
                               @Nonnull TypeDecl typeDecl) {
        super(ArrayNode.class);
        this.gen = gen;
        this.factory = factory;
        this.typeDecl = typeDecl;
    }

    public static ArrayNodeGenerator create(@Nonnull Structure structure,
                                            @Nonnull TypeDecl typeDecl) {
        ArrayListGenerator gen = new ArrayListGenerator();
        Generator<? extends JsonNode> delegate = JsonNodeGenerator.buildDelegate(structure, typeDecl.items);
        gen.addComponentGenerators(Collections.singletonList(delegate));
        if ((typeDecl.minItems != null) || (typeDecl.maxItems != null)) {
            gen.configure(new SizeImpl(typeDecl.minItems, typeDecl.maxItems));
        }
        return new ArrayNodeGenerator(gen, JsonNodeFactory.instance, typeDecl);
    }

    @Override
    public ArrayNode generate(SourceOfRandomness random, GenerationStatus status) {
        ArrayNode result = new ArrayNode(factory);
        List<JsonNode> items = gen.generate(random, status);
        result.addAll(items);
        return result;
    }

    @Override
    public List<ArrayNode> doShrink(SourceOfRandomness random, ArrayNode larger) {
        List<ArrayNode> result = new ArrayList<>();
        ArrayList<JsonNode> convert = new ArrayList<>();
        larger.forEach(convert::add);
        List<ArrayList> items = gen.doShrink(random, convert);
        for (ArrayList list : items) {
            ArrayNode node = new ArrayNode(factory);
            node.addAll(list);
            result.add(node);
        }
        return result;
    }

}
