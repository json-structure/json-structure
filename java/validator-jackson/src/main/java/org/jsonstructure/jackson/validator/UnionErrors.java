package org.jsonstructure.jackson.validator;

import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import javax.annotation.Nonnull;

import org.jsonstructure.jackson.validator.error.CompositeError;
import org.jsonstructure.jackson.validator.error.EnumError;
import org.jsonstructure.jackson.validator.error.ValidationError;
import org.jsonstructure.jackson.validator.loanword.Slice;

import static org.jsonstructure.jackson.validator.error.ValidationError.errorAt;

public class UnionErrors {

    public static ValidationError filterErrors(@Nonnull Map<String, ValidationError> errors,
                                               @Nonnull Structure structure,
                                               @Nonnull Slice<String> scope) {
        if (structure.options.unionErrors == Options.UnionErrorReport.PRIORITY) {
            return unionPriorityErrors(errors, scope);
        }
        return unionAllErrors(errors);
    }

    private static ValidationError filterPriority(List<ValidationError> errs, Slice<String> scope) {
        CompositeError result = new CompositeError();
        for (ValidationError err : errs) {
            if (err instanceof EnumError) {
                String errScope = err.scope;
                if (errScope != null) {
                    int count = (errScope.length() == 1) ? 0 : (int) errScope.chars().filter(ch -> ch =='/').count();
                    if (count == (scope.length() + 1)) {
                        result.add(err);
                    }
                }
            }
        }
        return result.simplify();
    }

    private static ValidationError unionPriorityErrors(@Nonnull Map<String, ValidationError> errors,
                                                       @Nonnull Slice<String> scope) {
        CompositeError errs = new CompositeError();
        errs.add(errorAt("union validation failure. Reporting a subset of errors for each type", null));
        Map<String, ValidationError> lowPriority = new HashMap<>();
        boolean hasPriority = false;
        for (Map.Entry<String, ValidationError> entry : errors.entrySet()) {
            ValidationError current = entry.getValue();
            ValidationError filter;
            if (current instanceof CompositeError) {
                filter = filterPriority(((CompositeError) current).children, scope);
            } else {
                filter = filterPriority(Collections.singletonList(current), scope);
            }
            lowPriority.put(entry.getKey(), filter);
            if (filter == null) {
                hasPriority = true;
            }
        }
        for (Map.Entry<String, ValidationError> entry : errors.entrySet()) {
            ValidationError p = lowPriority.get(entry.getKey());
            if (hasPriority && (p != null)) {
                continue;
            }
            String msg = String.format("Attempted to validate against union type '%s' - %s",
                    entry.getKey(), entry.getValue());
            errs.add(errorAt(msg, null));
        }
        return errs;
    }

    private static ValidationError unionAllErrors(@Nonnull Map<String, ValidationError> errors) {
        CompositeError errs = new CompositeError();
        errs.add(errorAt("union validation failure. Reporting all errors for each type", null));
        for (Map.Entry<String, ValidationError> entry : errors.entrySet()) {
            String msg = String.format("Attempted to validate against union type '%s' - %s",
                    entry.getKey(), entry.getValue());
            errs.add(errorAt(msg, null));
        }
        return errs;
    }

}
