package org.jsonstructure.jackson.validator.format;

import java.net.URI;
import java.net.URISyntaxException;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

import com.fasterxml.jackson.databind.JsonNode;

import org.jsonstructure.jackson.validator.Format;

public class UriFormat implements Format {

    @Override
    public boolean accept(@Nonnull String type) {
        return "string".equals(type);
    }

    @Override
    @Nullable
    public String apply(@Nonnull JsonNode value, @Nonnull String type) {
        String str = value.textValue();
        try {
            new URI(str);
            return null;
        } catch (URISyntaxException ex) {
            return "value is not a valid URI";
        }
    }
}
