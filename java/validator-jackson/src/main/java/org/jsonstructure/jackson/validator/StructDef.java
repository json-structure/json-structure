package org.jsonstructure.jackson.validator;

import java.io.IOException;
import java.io.InputStream;
import java.util.Collections;
import java.util.HashMap;
import java.util.Map;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import org.jsonstructure.jackson.validator.error.ValidationError;
import org.jsonstructure.jackson.validator.loanword.Result;

public class StructDef {

    @Nullable
    final String title;

    @Nullable
    final String description;

    @Nonnull
    final Map<String, String> imports;

    @Nonnull
    final Map<String, JsonNode> fragments;

    @Nonnull
    final Map<String, TypeDecl> types;

    @Nullable
    final TypeDecl main;

    @JsonCreator
    public StructDef(@Nullable @JsonProperty("title") String title,
                     @Nullable @JsonProperty("description") String description,
                     @Nullable @JsonProperty("imports") Map<String, String> imports,
                     @Nullable @JsonProperty("fragments") Map<String, JsonNode> fragments,
                     @Nullable @JsonProperty("types") Map<String, TypeDecl> types,
                     @Nullable @JsonProperty("main") TypeDecl main) {
        this.title = title;
        this.description = description;
        this.imports = (imports == null) ? Collections.emptyMap() : imports;
        this.fragments = (fragments == null) ? Collections.emptyMap() : fragments;
        this.types = (types == null) ? Collections.emptyMap() : types;
        this.main = main;
    }

    @Nonnull
    public static Result<StructDef, ValidationError> create(@Nonnull InputStream inputStream)
            throws IOException {

        ObjectMapper objectMapper = Jackson.OBJECT_MAPPER;
        JsonNode tree = objectMapper.readTree(inputStream);
        return createNode(tree);
    }

    @Nonnull
    public static Result<StructDef, ValidationError> createNode(@Nonnull JsonNode tree)
            throws IOException {
        ValidationError error = Compose.compose(tree);
        if (error != null) {
            return Result.err(error);
        }
        StructDef definition = Jackson.OBJECT_MAPPER.treeToValue(tree, StructDef.class);
        return Result.ok(definition);
    }


}