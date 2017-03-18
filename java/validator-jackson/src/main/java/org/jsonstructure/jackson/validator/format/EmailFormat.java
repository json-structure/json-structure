package org.jsonstructure.jackson.validator.format;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

import com.fasterxml.jackson.databind.JsonNode;

import org.apache.commons.validator.routines.EmailValidator;
import org.jsonstructure.jackson.validator.Format;

public class EmailFormat implements Format {

    @Override
    public boolean accept(@Nonnull String type) {
        return "string".equals(type);
    }

    @Override
    @Nullable
    public String apply(@Nonnull JsonNode value, @Nonnull String type) {
        if (!value.isTextual()) {
            return null;
        }
        String str = value.textValue();
        if (!EmailValidator.getInstance().isValid(str)) {
            return "value is not a valid email address";
        }
        return null;
    }
}
