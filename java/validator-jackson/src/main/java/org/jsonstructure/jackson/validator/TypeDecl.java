package org.jsonstructure.jackson.validator;

import java.math.BigDecimal;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.JsonNode;

import org.jsonstructure.jackson.validator.error.CompositeError;
import org.jsonstructure.jackson.validator.error.ValidationError;
import org.jsonstructure.jackson.validator.loanword.Slice;

import static org.jsonstructure.jackson.validator.error.ValidationError.errorAt;

public class TypeDecl {

    // required

    @Nullable
    final String type;

    // properties available to all type declarations

    @Nullable
    final JsonNode defaultValue;

    @Nullable
    final JsonNode[] enumValues;

    // common to primitive types

    @Nullable
    final String format;

    @Nullable
    final Boolean nullable;

    // number

    @Nullable
    final BigDecimal multipleOf;

    @Nullable
    final BigDecimal minimum;

    @Nullable
    final BigDecimal maximum;

    @Nullable
    final BigDecimal exclusiveMinimum;

    @Nullable
    final BigDecimal exclusiveMaximum;

    // string

    @Nullable
    final Integer minLength;

    @Nullable
    final Integer maxLength;

    // struct

    @Nullable
    final Map<String, TypeDecl> fields;

    // array, set, and map types

    @Nullable
    final TypeDecl items;

    @Nullable
    final Integer minItems;

    @Nullable
    final Integer maxItems;

    // union

    final Map<String, TypeDecl> types;

    @JsonCreator
    public TypeDecl(@JsonProperty("type") String type,
                    @JsonProperty("default") JsonNode defaultValue,
                    @JsonProperty("enum") JsonNode[] enumValues,
                    @JsonProperty("format") String format,
                    @JsonProperty("nullable") Boolean nullable,
                    @JsonProperty("multipleOf") BigDecimal multipleOf,
                    @JsonProperty("minimum") BigDecimal minimum,
                    @JsonProperty("maximum") BigDecimal maximum,
                    @JsonProperty("exclusiveMinimum") BigDecimal exclusiveMinimum,
                    @JsonProperty("exclusiveMaximum") BigDecimal exclusiveMaximum,
                    @JsonProperty("minLength") Integer minLength,
                    @JsonProperty("maxLength") Integer maxLength,
                    @JsonProperty("feilds") Map<String, TypeDecl> fields,
                    @JsonProperty("items") TypeDecl items,
                    @JsonProperty("minItems") Integer minItems,
                    @JsonProperty("maxItems") Integer maxItems,
                    @JsonProperty("types") Map<String, TypeDecl> types) {

        this.type = type;
        this.defaultValue = defaultValue;
        this.enumValues = enumValues;
        this.format = format;
        this.nullable = nullable;
        this.multipleOf = multipleOf;
        this.minimum = minimum;
        this.maximum = maximum;
        this.exclusiveMinimum = exclusiveMinimum;
        this.exclusiveMaximum = exclusiveMaximum;
        this.minLength = minLength;
        this.maxLength = maxLength;
        this.fields = fields;
        this.items = items;
        this.minItems = minItems;
        this.maxItems = maxItems;
        this.types = types;
    }

    @Nullable
    public ValidationError validateDecl(@Nonnull Structure structure, @Nonnull Slice<String> scope) {
        CompositeError errors = new CompositeError();
        if ((type == null) || (type.length() == 0)) {
            return errorAt("missing required property 'type'", scope);
        }
        PrimitiveTypes.PermissibleFields pf = PrimitiveTypes.PERMISSIBLE_FIELDS.get(type);
        TypeDecl decl = structure.definition.types.get(type);
        if ((pf == null) && (decl == null)) {
            return errorAt("unknown type '" + type + "'", scope);
        }
        if ((pf != null) && (decl != null)) {
            return errorAt("unexpected error. type declaration '" + type + "' shadows primitive type", scope);
        }
        if (decl != null) {
            return detectTypeAliasCycle(structure, decl, new HashSet<>(), scope);
        }
        errors.add(permissible("format", type, pf, format != null, scope));
        errors.add(permissible("nullable", type, pf, nullable != null, scope));
        errors.add(permissible("multipleOf", type, pf, multipleOf != null, scope));
        errors.add(permissible("minimum", type, pf, minimum != null, scope));
        errors.add(permissible("maximum", type, pf, maximum != null, scope));
        errors.add(permissible("exclusiveMinimum", type, pf, exclusiveMinimum != null, scope));
        errors.add(permissible("exclusiveMaximum", type, pf, exclusiveMaximum != null, scope));
        errors.add(permissible("pattern", type, pf, false, scope));
        errors.add(permissible("minLength", type, pf, minLength != null, scope));
        errors.add(permissible("maxLength", type, pf, maxLength != null, scope));
        errors.add(permissible("fields", type, pf, fields != null, scope));
        errors.add(permissible("items", type, pf, items != null, scope));
        errors.add(permissible("minItems", type, pf, minItems != null, scope));
        errors.add(permissible("maxItems", type, pf, maxItems != null, scope));
        errors.add(permissible("types", type, pf, types != null, scope));

        errors.add(validateNumberTypeDecl(scope));
        errors.add(validateStringTypeDecl(scope));
        errors.add(validateStructTypeDecl(structure, scope));
        errors.add(validateCollectionTypeDecl(structure, scope));
        errors.add(validateUnionTypeDecl(structure, scope));

        return errors.simplify();
    }

    @Nullable
    public ValidationError validate(@Nullable JsonNode value,
                                    @Nonnull Structure structure,
                                    @Nonnull Slice<String> scope) {
        CompositeError errors = new CompositeError();
        if (!PrimitiveTypes.PRIMITIVE_TYPES.contains(type)) {
            TypeDecl def = structure.definition.types.get(type);
            if (def == null) {
                return errorAt("Unknown type '" + type + "'", scope);
            }
            return def.validate(value, structure, scope);
        }
        if ((value == null) || value.isNull()) {
            if ((nullable != null) && nullable) {
                return null;
            }
            return errorAt("JSON null value when nullable property is false", scope);
        }
        return errors;
    }

    @Nullable
    private ValidationError permissible(@Nonnull String name, @Nonnull String type,
                                        @Nonnull PrimitiveTypes.PermissibleFields fields,
                                        boolean observed, @Nonnull Slice<String> scope) {
        boolean allowed = fields.contains(type);
        if (observed && !allowed) {
            return errorAt("Property " + name + " is not allowed with type " + type, scope);
        }
        return null;
    }

    @Nullable
    private ValidationError validateNumberTypeDecl(@Nonnull Slice<String> scope) {
        CompositeError errors = new CompositeError();
        BigDecimal min = null, max = null;
        if (!"integer".equals(type) && !"number".equals("type")) {
            return null;
        }
        if ((minimum != null) && (exclusiveMinimum != null)) {
            errors.add(errorAt("'minimum' and 'exclusiveMinimum' are both defined", scope));
        }
        if ((maximum != null) && (exclusiveMaximum != null)) {
            errors.add(errorAt("'maximum' and 'exclusiveMaximum' are both defined", scope));
        }
        if (minimum != null) {
            min = minimum;
        } else if (exclusiveMinimum != null) {
            min = exclusiveMinimum;
        }
        if (maximum != null) {
            max = maximum;
        } else if (exclusiveMaximum != null) {
            max = exclusiveMaximum;
        }
        if ((min != null) && (max != null) && min.compareTo(max) > 0) {
            errors.add(errorAt("'maximum' is less than 'minimum'", scope));
        }
        if ((multipleOf != null) && multipleOf.compareTo(BigDecimal.ZERO) < 0) {
            errors.add(errorAt("multipleOf is less than minimum", scope));
        }
        return errors;
    }

    @Nullable
    private ValidationError validateStringTypeDecl(@Nonnull Slice<String> scope) {
        CompositeError errors = new CompositeError();
        if (!"string".equals(type)) {
            return null;
        }
        if ((minLength != null) && (minLength < 0)) {
            errors.add(errorAt("'minLength' must be a non-negative integer", scope));
        }
        if ((maxLength != null) && (maxLength < 0)) {
            errors.add(errorAt("'maxLength' must be a non-negative integer", scope));
        }
        if ((minLength != null) && (maxLength != null) && (minLength > maxLength)) {
            errors.add(errorAt("'maxLength' is less than 'minLength'", scope));
        }
        return errors;
    }

    @Nullable
    private ValidationError validateStructTypeDecl(@Nonnull Structure structure, @Nonnull Slice<String> scope) {
        CompositeError errors = new CompositeError();
        if (!"struct".equals(type)) {
            return null;
        }
        if (fields == null) {
            return errorAt("missing required property 'fields'", scope);
        }
        for (Map.Entry<String, TypeDecl> entry : fields.entrySet()) {
            Slice<String> newScope = scope.append("fields", entry.getKey());
            errors.add(entry.getValue().validateDecl(structure, newScope));
        }
        return errors;
    }

    @Nullable
    private ValidationError validateCollectionTypeDecl(@Nonnull Structure structure, @Nonnull Slice<String> scope) {
        CompositeError errors = new CompositeError();
        if (!"array".equals(type) && !"set".equals(type) && !"map".equals(type)) {
            return null;
        }
        if (items == null) {
            return errorAt("missing required property 'items'", scope);
        }
        Slice<String> newScope = scope.append("items");
        errors.add(items.validateDecl(structure, newScope));
        if ((minItems != null) && (minItems < 0)) {
            errors.add(errorAt("'minItems' must be a non-negative integer", scope));
        }
        if ((maxItems != null) && (maxItems < 0)) {
            errors.add(errorAt("'maxItems' must be a non-negative integer", scope));
        }
        if ((minItems != null) && (maxItems != null) && (minItems > maxItems)) {
            errors.add(errorAt("'maxItems' is less than 'minItems'", scope));
        }
        return errors;
    }

    @Nullable
    private ValidationError validateUnionTypeDecl(@Nonnull Structure structure, @Nonnull Slice<String> scope) {
        CompositeError errors = new CompositeError();
        if (!"union".equals(type)) {
            return null;
        }
        if (types == null) {
            return errorAt("missing required property 'types'", scope);
        } else if (types.isEmpty()) {
            return errorAt("'types' must have at least one entry", scope);
        }
        for (Map.Entry<String, TypeDecl> entry : types.entrySet()) {
            Slice<String> newScope = scope.append("types", entry.getKey());
            errors.add(entry.getValue().validateDecl(structure, newScope));
        }
        return errors;
    }

    @Nullable
    private static ValidationError detectTypeAliasCycle(@Nonnull Structure structure,
                                                        @Nonnull TypeDecl td,
                                                        @Nonnull Set<String> prev,
                                                        @Nonnull Slice<String> scope) {
        String name = td.type;
        TypeDecl decl = structure.definition.types.get(name);
        if (prev.contains(name)) {
            return errorAt("Type alias cycle detected " + prev, scope);
        }
        if (decl == null) {
            return null;
        }
        prev.add(name);
        return detectTypeAliasCycle(structure, decl, prev, scope);
    }

    @Nullable
    ValidationError validateEmbedded(@Nonnull Structure structure, @Nonnull Slice<String> scope) {
        CompositeError errors = new CompositeError();
        if (enumValues != null) {
            for (int i = 0; i < enumValues.length; i++) {
                Slice<String> iScope = scope.append("enum", Integer.toString(i));
                errors.add(validate(enumValues[i], structure, iScope));
            }
        }
        if (defaultValue != null) {
            Slice<String> newScope = scope.append("default");
            errors.add(validate(defaultValue, structure, newScope));
        }
        if (fields != null) {
            for (Map.Entry<String, TypeDecl> entry : fields.entrySet()) {
                Slice<String> newScope = scope.append("fields", entry.getKey());
                errors.add(entry.getValue().validateEmbedded(structure, newScope));
            }
        }
        if (types != null) {
            for (Map.Entry<String, TypeDecl> entry : types.entrySet()) {
                Slice<String> newScope = scope.append("types", entry.getKey());
                errors.add(entry.getValue().validateEmbedded(structure, newScope));
            }
        }
        if (items != null) {
            Slice<String> newScope = scope.append("items");
            errors.add(items.validateEmbedded(structure, newScope));
        }
        return errors;
    }

    @Nullable
    private ValidationError validateBoolean(@Nonnull JsonNode value,
                                            @Nonnull Slice<String> scope) {
        if (!value.isBoolean()) {
            return errorAt("JSON value is not a boolean", scope);
        }
        return null;
    }


}
