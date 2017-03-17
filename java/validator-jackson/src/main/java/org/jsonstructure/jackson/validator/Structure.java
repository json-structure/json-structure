package org.jsonstructure.jackson.validator;

import java.io.IOException;
import java.io.InputStream;
import java.util.Map;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

import org.jsonstructure.jackson.validator.error.CompositeError;
import org.jsonstructure.jackson.validator.error.ValidationError;
import org.jsonstructure.jackson.validator.loanword.Result;
import org.jsonstructure.jackson.validator.loanword.Slice;

import static org.jsonstructure.jackson.validator.error.ValidationError.errorAt;

public class Structure {

    final StructDef definition;
    boolean initialized;

    @Nonnull
    public static Result<Structure, ValidationError> create(@Nonnull InputStream inputStream)
            throws IOException {

        Result<StructDef, ValidationError> child = StructDef.create(inputStream);
        if (child.isError()) {
            // Silence false positive NPE inspection
            assert(child.getErr() != null);
            return Result.err(child.getErr());
        }
        Structure structure = new Structure(child.getOk());
        ValidationError error = structure.validateStructure();
        if (error != null) {
            return Result.err(error);
        }
        return Result.ok(structure);
    }

    Structure(StructDef definition) {
        this.definition = definition;
    }

    public void resetValidation() {
        initialized = false;
    }

    public ValidationError validateStructure() {
        if (initialized) {
            return null;
        }
        ValidationError error = doValidateStructure();
        if (error == null) {
            initialized = true;
        }
        return error;
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


}
