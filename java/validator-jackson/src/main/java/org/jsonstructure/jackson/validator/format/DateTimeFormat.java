package org.jsonstructure.jackson.validator.format;

import java.time.ZonedDateTime;
import java.time.format.DateTimeFormatter;
import java.time.format.DateTimeParseException;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

import com.fasterxml.jackson.databind.JsonNode;

import org.jsonstructure.jackson.validator.Format;

public class DateTimeFormat implements Format {

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
        try {
            String str = value.textValue();
            str = str.replace("-00:00", "+00:00");
            ZonedDateTime.parse(str, DateTimeFormatter.ISO_OFFSET_DATE_TIME);
        } catch (DateTimeParseException ex) {
            return "value is not a valid rfc 3339 date-time format: " + ex.getMessage();
        }
        return null;
    }
}
