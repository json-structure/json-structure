package org.jsonstructure.jackson.validator.error;

import java.util.List;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

import org.jsonstructure.jackson.validator.loanword.Slice;

public class JSONStructureError {

    @Nullable
    private final String message;

    @Nullable
    private final String scope;

    JSONStructureError() {
        this.message = null;
        this.scope = null;
    }

    private static String join(Slice<String> scope) {
        List<String> values = scope.toList();
        int size = values.size();
        StringBuilder builder = new StringBuilder();
        for (int i = 0; i < size; i++) {
            builder.append(values.get(i));
            if (i < (size - 1)) {
                builder.append("/");
            }
        }
        return builder.toString();
    }

    public static JSONStructureError errorAt(@Nonnull String message, @Nonnull Slice<String> scope) {
        String scopeStr = "/" + join(scope);
        return new JSONStructureError(message, scopeStr);
    }

    JSONStructureError(@Nullable String message, @Nullable String scope) {
        this.message = message;
        this.scope = scope;
    }
}
