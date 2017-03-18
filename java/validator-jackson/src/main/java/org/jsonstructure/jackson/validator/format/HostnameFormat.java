package org.jsonstructure.jackson.validator.format;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

import com.fasterxml.jackson.databind.JsonNode;

import org.jsonstructure.jackson.validator.Format;

public class HostnameFormat implements Format {

    private static final Pattern ILLEGAL_LABEL_CHARS = Pattern.compile("[^0-9a-zA-Z.\\-]");

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
        if (str.length() > 253) {
            return "hostname is greater than 253 characters";
        }
        String[] components = str.split("\\.");
        for (int i = 0; i < components.length; i++) {
            String label = components[i];
            if (label.length() < 1) {
                return String.format("label %d of hostname is 0 characters", i+1);
            }
            if (label.length() > 63) {
                return String.format("label %d of hostname > 63 characters", i+1);
            }
        }
        Matcher matcher = ILLEGAL_LABEL_CHARS.matcher(str);
        if (matcher.find()) {
            return "hostname has illegal character " + matcher.group();
        }
        return null;
    }
}
