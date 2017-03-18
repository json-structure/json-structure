package org.jsonstructure.jackson.validator;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

import com.fasterxml.jackson.databind.JsonNode;

public interface Format {

    boolean accept(@Nonnull String type);

    @Nullable
    String apply(@Nonnull JsonNode value, @Nonnull String type);
}
