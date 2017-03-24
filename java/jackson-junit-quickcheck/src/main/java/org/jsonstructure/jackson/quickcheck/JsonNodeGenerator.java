package org.jsonstructure.jackson.quickcheck;

import javax.annotation.Nonnull;

import com.fasterxml.jackson.databind.JsonNode;
import com.pholser.junit.quickcheck.generator.GenerationStatus;
import com.pholser.junit.quickcheck.generator.Generator;
import com.pholser.junit.quickcheck.random.SourceOfRandomness;

import org.jsonstructure.jackson.quickcheck.generator.ArrayNodeGenerator;
import org.jsonstructure.jackson.quickcheck.generator.BooleanNodeGenerator;
import org.jsonstructure.jackson.quickcheck.generator.UnionNodeGenerator;
import org.jsonstructure.jackson.validator.PrimitiveTypes;
import org.jsonstructure.jackson.validator.Structure;
import org.jsonstructure.jackson.validator.TypeDecl;

public class JsonNodeGenerator extends Generator<JsonNode> {

    @Nonnull
    final Generator<? extends JsonNode> delegate;

    JsonNodeGenerator(@Nonnull Generator<? extends JsonNode> delegate) {
        super(JsonNode.class);
        this.delegate = delegate;
    }

    public static Generator<JsonNode> create(Structure structure, TypeDecl decl) {
        return new JsonNodeGenerator(buildDelegate(structure, decl));
    }

    public static Generator<? extends JsonNode> buildDelegate(Structure structure, TypeDecl decl) {
        String type = decl.type;
        if (type == null) {
            throw new NullPointerException("type name must be non-null");
        }
        if (!PrimitiveTypes.PRIMITIVE_TYPES.contains(type)) {
            TypeDecl typeDecl = structure.definition.types.get(type);
            if (typeDecl == null) {
                throw new IllegalArgumentException("type alias " + type + " is undefined");
            }
            return create(structure, typeDecl);
        }
        switch (type) {
            case "boolean":
                return BooleanNodeGenerator.create();
            case "array":
                return ArrayNodeGenerator.create(structure, decl);
            case "union":
                return UnionNodeGenerator.create(structure, decl);
            default:
                throw new IllegalStateException("unhandled primitive type " + type);
        }
    }

    @Override
    public JsonNode generate(SourceOfRandomness random, GenerationStatus status) {
        return delegate.generate(random, status);
    }


}
