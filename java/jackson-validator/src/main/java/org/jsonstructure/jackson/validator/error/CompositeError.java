package org.jsonstructure.jackson.validator.error;

import java.util.ArrayList;
import java.util.List;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

public class CompositeError extends ValidationError {

    @Nonnull
    public final List<ValidationError> children;

    public CompositeError() {
        children = new ArrayList<>();
    }

    public void add(@Nullable ValidationError ex) {
        if (ex == null) {
            return;
        }
        if (ex instanceof CompositeError) {
            CompositeError cex = (CompositeError) ex;
            for (ValidationError child : cex.children) {
                add(child);
            }
        } else {
            children.add(ex);
        }
    }

    public int size() {
        return children.size();
    }

    public ValidationError simplify() {
        if (children.size() == 0) {
            return null;
        } else if (children.size() == 1) {
            return children.get(0);
        } else {
            return this;
        }
    }

    @Override
    public String toString() {
        StringBuilder builder = new StringBuilder();
        builder.append(children.size());
        builder.append(" errors occurred:\n\n");
        for (int i = 0; i < children.size(); i++) {
            builder.append(children.get(i).toString());
            if (i < (children.size() - 1)) {
                builder.append("\n");
            }
        }
        return builder.toString();
    }

}
