package org.jsonstructure.jackson.validator.error;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

import org.jsonstructure.jackson.validator.loanword.Slice;

public class EnumError extends ValidationError {

    EnumError(@Nullable String message, @Nullable String scope) {
        super(message, scope);
    }

    public static EnumError enumError(@Nonnull String message, @Nullable Slice<String> scope) {
        if (scope == null) {
            return new EnumError(message, null);
        } else {
            String scopeStr = "/" + join(scope);
            return new EnumError(message, scopeStr);
        }
    }
}
