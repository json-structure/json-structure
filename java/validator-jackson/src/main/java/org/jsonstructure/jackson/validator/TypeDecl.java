package org.jsonstructure.jackson.validator;

import java.math.BigDecimal;
import java.util.HashSet;
import java.util.Iterator;
import java.util.Map;
import java.util.Set;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.node.ArrayNode;
import com.fasterxml.jackson.databind.node.MissingNode;
import com.fasterxml.jackson.databind.node.ObjectNode;

import org.jsonstructure.jackson.validator.error.CompositeError;
import org.jsonstructure.jackson.validator.error.ValidationError;
import org.jsonstructure.jackson.validator.loanword.Slice;

import static org.jsonstructure.jackson.validator.error.ValidationError.errorAt;

public class TypeDecl {

    // required

    @Nonnull
    @JsonProperty("type")
    public String type;

    // properties available to all type declarations

    @Nonnull
    @JsonProperty("default")
    public JsonNode defaultValue = MissingNode.getInstance();

    @Nullable
    @JsonProperty("enum")
    public JsonNode[] enumValues;

    @Nonnull
    @JsonIgnore
    final Set<JsonNode> enumSet = new HashSet<>();

    // common to primitive types

    @Nullable
    @JsonProperty("format")
    public String format;

    @Nullable
    @JsonProperty("nullable")
    public Boolean nullable;

    // number

    @Nullable
    @JsonProperty("multipleOf")
    public BigDecimal multipleOf;

    @Nullable
    @JsonProperty("minimum")
    public BigDecimal minimum;

    @Nullable
    @JsonProperty("maximum")
    public BigDecimal maximum;

    @Nullable
    @JsonProperty("exclusiveMinimum")
    public BigDecimal exclusiveMinimum;

    @Nullable
    @JsonProperty("exclusiveMaximum")
    public BigDecimal exclusiveMaximum;

    // string

    @Nullable
    @JsonProperty("pattern")
    public String pattern;

    @Nullable
    @JsonProperty("minLength")
    public Integer minLength;

    @Nullable
    @JsonProperty("maxLength")
    public Integer maxLength;

    // struct

    @Nullable
    @JsonProperty("fields")
    public Map<String, TypeDecl> fields;

    // array, set, and map types

    @Nullable
    @JsonProperty("items")
    public TypeDecl items;

    @Nullable
    @JsonProperty("minItems")
    public Integer minItems;

    @Nullable
    @JsonProperty("maxItems")
    public Integer maxItems;

    // union

    @JsonProperty("types")
    public Map<String, TypeDecl> types;

    @Nullable
    public ValidationError validateDecl(@Nonnull Structure structure, @Nonnull Slice<String> scope) {
        CompositeError errors = new CompositeError();
        if (type.length() == 0) {
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

        errors.add(validateEnum(value, scope));

        switch (type) {
            case "boolean":
                errors.add(validateBoolean(value, scope));
                break;
            case "integer":
                errors.add(validateInteger(value, scope));
                break;
            case "number":
                errors.add(validateNumber(value, scope));
                break;
            case "string":
                errors.add(validateString(value, scope));
                break;
            case "json":
                // no validation
                break;
            case "struct":
                errors.add(validateStruct(value, structure, scope));
                break;
            case "array":
                errors.add(validateArray(value, structure, scope));
                break;
            case "set":
                errors.add(validateSet(value, structure, scope));
                break;
            case "map":
                errors.add(validateMap(value, structure, scope));
                break;
            case "union":
                //errors.add(validateUnion(value, structure, scope));
                break;
            default:
                errors.add(errorAt("Internal error. Unhandled primitive type '" + type + "'", scope));
        }
        return errors.simplify();
    }

    @Nullable
    private ValidationError permissible(@Nonnull String name, @Nonnull String type,
                                        @Nonnull PrimitiveTypes.PermissibleFields fields,
                                        boolean observed, @Nonnull Slice<String> scope) {
        boolean allowed = fields.contains(name);
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
                JsonNode value = enumValues[i];
                Slice<String> iScope = scope.append("enum", Integer.toString(i));
                if (value.isNull()) {
                    errors.add(errorAt("null enum value is not permitted", iScope));
                } else if (!enumSet.add(value)) {
                    errors.add(errorAt("duplicate enum value", iScope));
                } else {
                    errors.add(validate(value, structure, iScope));
                }
            }
        }
        if (!defaultValue.isMissingNode()) {
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

    @Nullable
    private ValidationError validateNumber(@Nonnull JsonNode value,
                                           @Nonnull Slice<String> scope) {
        CompositeError errors = new CompositeError();
        if (!value.isNumber()) {
            return errorAt("JSON value is not a number", scope);
        }
        BigDecimal dec = value.decimalValue();
        if (minimum != null && minimum.compareTo(dec) > 0) {
            errors.add(errorAt(String.format("%s is less than minimum %s", dec, minimum), scope));
        }
        if (exclusiveMinimum != null && exclusiveMinimum.compareTo(dec) >= 0) {
            errors.add(errorAt(String.format("%s is less than or equal to exclusiveMinimum %s", dec, exclusiveMinimum), scope));
        }
        if (maximum != null && maximum.compareTo(dec) < 0) {
            errors.add(errorAt(String.format("%s is greater than maximum %s", dec, maximum), scope));
        }
        if (exclusiveMaximum != null && exclusiveMaximum.compareTo(dec) <= 0) {
            errors.add(errorAt(String.format("%s is greater than or equal to exclusiveMinimum %s", dec, exclusiveMaximum), scope));
        }
        if (multipleOf != null && dec.remainder(multipleOf).compareTo(BigDecimal.ZERO) != 0) {
            errors.add(errorAt(String.format("%s is not a multiple of %s", dec, multipleOf), scope));
        }
        return errors.simplify();
    }

    @Nullable
    private ValidationError validateInteger(@Nonnull JsonNode value,
                                            @Nonnull Slice<String> scope) {
        ValidationError errors = validateNumber(value, scope);
        if (errors != null) {
            return errors;
        }
        BigDecimal dec = value.decimalValue();
        if (dec.setScale(0, BigDecimal.ROUND_FLOOR).compareTo(dec) != 0) {
            return errorAt(String.format("%s is not an integer", dec), scope);
        }
        return null;
    }

    @Nullable
    private ValidationError validateString(@Nonnull JsonNode value,
                                           @Nonnull Slice<String> scope) {
        CompositeError errors = new CompositeError();
        if (!value.isTextual()) {
            return errorAt("JSON value is not a string", scope);
        }
        String str = value.textValue();
        if (minLength != null && str.length() < minLength) {
            errors.add(errorAt(String.format("length of string %s is less than minimum length %s", str, minLength), scope));
        }
        if (maxLength != null && str.length() > maxLength) {
            errors.add(errorAt(String.format("length of string %s is greater than maximum length %s", str, maxLength), scope));
        }
        return errors;
    }

    @Nullable
    private ValidationError validateStruct(@Nonnull JsonNode value,
                                           @Nonnull Structure structure,
                                           @Nonnull Slice<String> scope) {
        CompositeError errors = new CompositeError();
        if (!value.isObject()) {
            return errorAt("JSON value is not an object", scope);
        }
        ObjectNode obj = (ObjectNode) value;
        if (fields == null) {
            return errorAt("struct definition is missing required property 'fields'", scope);
        }
        for (Map.Entry<String, TypeDecl> entry : fields.entrySet()) {
            String key = entry.getKey();
            JsonNode child = obj.get(key);
            if (child != null) {
                Slice<String> newScope = scope.append(key);
                errors.add(entry.getValue().validate(child, structure, newScope));
            } else if (entry.getValue().defaultValue.isMissingNode()) {
                errors.add(errorAt("missing required field '" + key + "'", scope));
            }
        }
        Iterator<String> iterator = obj.fieldNames();
        while (iterator.hasNext()) {
            String name = iterator.next();
            if (fields.get(name) == null) {
                errors.add(errorAt("unrecognized field '" + name + "'", scope));
            }
        }
        return errors;
    }

    @Nullable
    private ValidationError validateArray(@Nonnull JsonNode value,
                                          @Nonnull Structure structure,
                                          @Nonnull Slice<String> scope) {
        CompositeError errors = new CompositeError();
        if (!value.isArray()) {
            return errorAt("JSON value is not an array", scope);
        }
        ArrayNode arr = (ArrayNode) value;
        if (items == null) {
            return errorAt("array definition is missing required property 'items'", scope);
        }
        if (minItems != null && arr.size() < minItems) {
            errors.add(errorAt(String.format("size %d of array is less than minimum items %s", arr.size(), minItems), scope));
        }
        if (maxItems != null && arr.size() > maxItems) {
            errors.add(errorAt(String.format("size %d of array is greater than maximum items %s", arr.size(), maxItems), scope));
        }
        Iterator<JsonNode> elements = arr.elements();
        for(int i = 0; elements.hasNext(); i++) {
            JsonNode child = elements.next();
            Slice<String> newScope = scope.append(Integer.toString(i));
            errors.add(items.validate(child, structure, newScope));
        }
        return errors;
    }

    private ValidationError validateSet(@Nonnull JsonNode value,
                                        @Nonnull Structure structure,
                                        @Nonnull Slice<String> scope) {
        CompositeError errors = new CompositeError();
        if (!value.isArray()) {
            return errorAt("JSON value is not a set", scope);
        }
        ArrayNode arr = (ArrayNode) value;
        if (items == null) {
            return errorAt("set definition is missing required property 'items'", scope);
        }
        if (minItems != null && arr.size() < minItems) {
            errors.add(errorAt(String.format("size %d of array is less than minimum items %s", arr.size(), minItems), scope));
        }
        if (maxItems != null && arr.size() > maxItems) {
            errors.add(errorAt(String.format("size %d of array is greater than maximum items %s", arr.size(), maxItems), scope));
        }
        HashSet<JsonNode> unique = new HashSet<>();
        Iterator<JsonNode> elements = arr.elements();
        for(int i = 0; elements.hasNext(); i++) {
            JsonNode child = elements.next();
            Slice<String> newScope = scope.append(Integer.toString(i));
            if (!unique.add(child)) {
                errors.add(errorAt("set has duplicate value", newScope));
            } else{
                errors.add(items.validate(child, structure, newScope));
            }
        }
        return errors;
    }

    private ValidationError validateMap(@Nonnull JsonNode value,
                                        @Nonnull Structure structure,
                                        @Nonnull Slice<String> scope) {
        CompositeError errors = new CompositeError();
        if (!value.isObject()) {
            return errorAt("JSON value is not a map", scope);
        }
        ObjectNode obj = (ObjectNode) value;
        if (items == null) {
            return errorAt("map definition is missing required property 'items'", scope);
        }
        if (minItems != null && obj.size() < minItems) {
            errors.add(errorAt(String.format("size %d of map is less than minimum items %s", obj.size(), minItems), scope));
        }
        if (maxItems != null && obj.size() > maxItems) {
            errors.add(errorAt(String.format("size %d of map is greater than maximum items %s", obj.size(), maxItems), scope));
        }
        Iterator<Map.Entry<String,JsonNode>> iterator = obj.fields();
        while (iterator.hasNext()) {
            Map.Entry<String,JsonNode> entry = iterator.next();
            Slice<String> newScope = scope.append(entry.getKey());
            errors.add(items.validate(entry.getValue(), structure, newScope));
        }
        return errors;
    }

    private ValidationError validateEnum(@Nonnull JsonNode value,
                                         @Nonnull Slice<String> scope) {
        if (enumValues == null) {
            return null;
        }
        if (!enumSet.contains(value)) {
            return errorAt("value not found in 'enum' set", scope);
        }
        return null;
    }

}
