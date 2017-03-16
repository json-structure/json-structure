package org.jsonstructure.jackson.validator;

import java.io.IOException;
import java.io.InputStream;
import java.util.Map;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

import org.jsonstructure.jackson.validator.error.CompositeError;
import org.jsonstructure.jackson.validator.error.JSONStructureError;
import org.jsonstructure.jackson.validator.loanword.Once;
import org.jsonstructure.jackson.validator.loanword.Result;
import org.jsonstructure.jackson.validator.loanword.Slice;

import static org.jsonstructure.jackson.validator.error.JSONStructureError.errorAt;

public class JSONStructure {

    final JSONStructureDefinition definition;
    final Once<JSONStructureError> validateOnce;

    @Nonnull
    public static Result<JSONStructure, JSONStructureError> create(@Nonnull InputStream inputStream)
            throws IOException {

        Result<JSONStructureDefinition, JSONStructureError> child = JSONStructureDefinition.create(inputStream);
        if (child.isError()) {
            // Silence false positive NPE inspection
            assert(child.getErr() != null);
            return Result.err(child.getErr());
        }
        JSONStructure structure = new JSONStructure(child.getOk());
        JSONStructureError error = structure.validateStructure();
        if (error != null) {
            return Result.err(error);
        }
        return Result.ok(structure);
    }

    JSONStructure(JSONStructureDefinition definition) {
        this.definition = definition;
        this.validateOnce = new Once<>(this::doValidateStructure);
    }

    @Nullable
    public JSONStructureError validateStructure() {
        return validateOnce.value();
    }

    @Nullable
    private JSONStructureError doValidateStructure() {
        JSONStructureError error = validateStructureTopLevel();
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
    private JSONStructureError validateStructureTopLevel() {
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

        return errors;
    }

    @Nullable
    private JSONStructureError validateStructureDecls() {
        CompositeError errors = new CompositeError();

        for (Map.Entry<String, TypeDecl> entry : definition.types.entrySet()) {
            Slice<String> scope = Slice.create("types", entry.getKey());
            errors.add(entry.getValue().validateDecl(this, scope));
        }
        assert(definition.main != null);
        Slice<String> scope = Slice.create("main");
        errors.add(definition.main.validateDecl(this, scope));
        return errors;
    }

    @Nullable
    private JSONStructureError validateEmbedded() {
        CompositeError errors = new CompositeError();

        for (Map.Entry<String, TypeDecl> entry : definition.types.entrySet()) {
            Slice<String> scope = Slice.create("types", entry.getKey());
            errors.add(entry.getValue().validateEmbedded(this, scope));
        }
        assert(definition.main != null);
        Slice<String> scope = Slice.create("main");
        errors.add(definition.main.validateEmbedded(this, scope));
        return errors;
    }


}
