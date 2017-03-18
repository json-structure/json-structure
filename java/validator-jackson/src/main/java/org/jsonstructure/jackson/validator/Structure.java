package org.jsonstructure.jackson.validator;

import java.io.IOException;
import java.io.InputStream;
import java.util.Map;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

import com.fasterxml.jackson.databind.JsonNode;

import org.jsonstructure.jackson.validator.error.CompositeError;
import org.jsonstructure.jackson.validator.error.ValidationError;
import org.jsonstructure.jackson.validator.loanword.Result;
import org.jsonstructure.jackson.validator.loanword.Slice;

import static org.jsonstructure.jackson.validator.error.ValidationError.errorAt;

public class Structure {

    @Nonnull
    final StructDef definition;

    @Nonnull
    final Options options;

    boolean initialized;
    ValidationError initError;

    Structure(@Nonnull StructDef definition, @Nonnull Options options) {
        this.definition = definition;
        this.options = options;
    }

    @Nonnull
    public static Result<Structure, ValidationError> create(@Nonnull InputStream inputStream, @Nonnull Options options)
            throws IOException {

        Result<StructDef, ValidationError> child = StructDef.create(inputStream);
        return buildStructure(child, options);
    }

    @Nonnull
    public static Result<Structure, ValidationError> createFromNode(@Nonnull JsonNode node, @Nonnull Options options)
            throws IOException {

        Result<StructDef, ValidationError> child = StructDef.createFromNode(node);
        return buildStructure(child, options);
    }

    private static Result<Structure, ValidationError> buildStructure(@Nonnull Result<StructDef, ValidationError> child,
                                                                     @Nonnull Options options) {
        if (child.isError()) {
            // Silence false positive NPE inspection
            assert(child.getErr() != null);
            return Result.err(child.getErr());
        }
        Structure structure = new Structure(child.getOk(), options);
        ValidationError error = structure.validateStructure();
        if (error != null) {
            return Result.err(error);
        }
        return Result.ok(structure);
    }

    public void reset() {
        initialized = false;
    }

    public ValidationError validateStructure() {
        if (!initialized) {
            initError = doValidateStructure();
            initialized = true;
        }
        return initError;
    }

    @Nullable
    private ValidationError doValidateStructure() {
        ValidationError error = validateStructureTopLevel();
        if (error != null) {
            return error;
        }
        error = validateStructureDecls();
        if (error != null) {
            return error;
        }
        error = validateEmbedded();
        if (error != null) {
            return error;
        }
        return null;
    }

    @Nullable
    private ValidationError validateStructureTopLevel() {
        CompositeError errors = new CompositeError();

        for (Map.Entry<String, TypeDecl> entry : definition.types.entrySet()) {
            Slice<String> scope = Slice.create("types", entry.getKey());
            if (entry.getValue() == null) {
                errors.add(errorAt("type declaration must be non-null", scope));
            }
        }
        Slice<String> scope = Slice.create("main");
        if (definition.main == null) {
            errors.add(errorAt("main type declaration must be declared", scope));
        }

        return errors.simplify();
    }

    @Nullable
    private ValidationError validateStructureDecls() {
        CompositeError errors = new CompositeError();

        for (Map.Entry<String, TypeDecl> entry : definition.types.entrySet()) {
            Slice<String> scope = Slice.create("types", entry.getKey());
            errors.add(entry.getValue().validateDecl(this, scope));
        }
        assert(definition.main != null);
        Slice<String> scope = Slice.create("main");
        errors.add(definition.main.validateDecl(this, scope));
        return errors.simplify();
    }

    @Nullable
    private ValidationError validateEmbedded() {
        CompositeError errors = new CompositeError();

        for (Map.Entry<String, TypeDecl> entry : definition.types.entrySet()) {
            Slice<String> scope = Slice.create("types", entry.getKey());
            errors.add(entry.getValue().validateEmbedded(this, scope));
        }
        assert(definition.main != null);
        Slice<String> scope = Slice.create("main");
        errors.add(definition.main.validateEmbedded(this, scope));
        return errors.simplify();
    }

    @Nullable
    public ValidationError validateFromNode(@Nonnull JsonNode node) {
        return definition.main.validate(node, this, Slice.empty());
    }


}
