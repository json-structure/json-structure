package org.jsonstructure.jackson.validator;

import java.math.BigDecimal;
import java.util.Map;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.annotation.JsonPOJOBuilder;
import com.fasterxml.jackson.databind.node.MissingNode;

@JsonPOJOBuilder(buildMethodName = "build", withPrefix = "")
public final class TypeDeclBuilder {
    String type;
    JsonNode defaultValue = MissingNode.getInstance();
    JsonNode[] enumValues;
    String format;
    Boolean nullable;
    BigDecimal multipleOf;
    BigDecimal minimum;
    BigDecimal maximum;
    BigDecimal exclusiveMinimum;
    BigDecimal exclusiveMaximum;
    String patternRaw;
    Integer minLength;
    Integer maxLength;
    Map<String, TypeDecl> fields;
    TypeDecl items;
    Integer minItems;
    Integer maxItems;
    Map<String, TypeDecl> types;

    public TypeDeclBuilder() {
    }

    public TypeDeclBuilder type(String type) {
        this.type = type;
        return this;
    }

    @JsonProperty("default")
    public TypeDeclBuilder defaultValue(JsonNode defaultValue) {
        this.defaultValue = defaultValue;
        return this;
    }

    @JsonProperty("enum")
    public TypeDeclBuilder enumValues(JsonNode[] enumValues) {
        this.enumValues = enumValues;
        return this;
    }

    public TypeDeclBuilder format(String format) {
        this.format = format;
        return this;
    }

    public TypeDeclBuilder nullable(Boolean nullable) {
        this.nullable = nullable;
        return this;
    }

    public TypeDeclBuilder multipleOf(BigDecimal multipleOf) {
        this.multipleOf = multipleOf;
        return this;
    }

    public TypeDeclBuilder minimum(BigDecimal minimum) {
        this.minimum = minimum;
        return this;
    }

    public TypeDeclBuilder maximum(BigDecimal maximum) {
        this.maximum = maximum;
        return this;
    }

    public TypeDeclBuilder exclusiveMinimum(BigDecimal exclusiveMinimum) {
        this.exclusiveMinimum = exclusiveMinimum;
        return this;
    }

    public TypeDeclBuilder exclusiveMaximum(BigDecimal exclusiveMaximum) {
        this.exclusiveMaximum = exclusiveMaximum;
        return this;
    }

    @JsonProperty("pattern")
    public TypeDeclBuilder patternRaw(String patternRaw) {
        this.patternRaw = patternRaw;
        return this;
    }

    public TypeDeclBuilder minLength(Integer minLength) {
        this.minLength = minLength;
        return this;
    }

    public TypeDeclBuilder maxLength(Integer maxLength) {
        this.maxLength = maxLength;
        return this;
    }

    public TypeDeclBuilder fields(Map<String, TypeDecl> fields) {
        this.fields = fields;
        return this;
    }

    public TypeDeclBuilder items(TypeDecl items) {
        this.items = items;
        return this;
    }

    public TypeDeclBuilder minItems(Integer minItems) {
        this.minItems = minItems;
        return this;
    }

    public TypeDeclBuilder maxItems(Integer maxItems) {
        this.maxItems = maxItems;
        return this;
    }

    public TypeDeclBuilder types(Map<String, TypeDecl> types) {
        this.types = types;
        return this;
    }

    public TypeDecl build() {
        TypeDecl typeDecl = new TypeDecl(type,
                defaultValue,
                enumValues,
                format,
                nullable,
                multipleOf,
                minimum,
                maximum,
                exclusiveMinimum,
                exclusiveMaximum,
                patternRaw,
                minLength,
                maxLength,
                fields,
                items,
                minItems,
                maxItems,
                types);
        return typeDecl;
    }
}
