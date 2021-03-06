package org.jsonstructure.jackson.validator.error;

import java.util.List;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

import org.jsonstructure.jackson.validator.loanword.Slice;

public class ValidationError {

    @Nullable
    public final String message;

    @Nullable
    public final String scope;

    ValidationError() {
        this.message = null;
        this.scope = null;
    }

    static String join(Slice<String> scope) {
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

    public static ValidationError errorAt(@Nonnull String message, @Nullable Slice<String> scope) {
        if (scope == null) {
            return new ValidationError(message, null);
        } else {
            String scopeStr = "/" + join(scope);
            return new ValidationError(message, scopeStr);
        }
    }

    ValidationError(@Nullable String message, @Nullable String scope) {
        this.message = message;
        this.scope = scope;
    }

    @Override
    public String toString() {
        if (scope != null) {
            return String.format("At %s: %s", scope, message);
        } else {
            return message;
        }
    }
}
