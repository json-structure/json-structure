package org.jsonstructure.jackson.validator.error;

import java.util.ArrayList;
import java.util.List;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

public class CompositeError extends JSONStructureError {

    @Nonnull
    private final List<JSONStructureError> children;

    public CompositeError() {
        children = new ArrayList<>();
    }

    public void add(@Nullable JSONStructureError ex) {
        if (ex == null) {
            return;
        }
        if (ex instanceof CompositeError) {
            CompositeError cex = (CompositeError) ex;
            for (JSONStructureError child : cex.children) {
                add(child);
            }
        } else {
            children.add(ex);
        }
    }

    public int size() {
        return children.size();
    }

    public JSONStructureError simplify() {
        if (children.size() == 0) {
            return null;
        } else if (children.size() == 1) {
            return children.get(0);
        } else {
            return this;
        }
    }

}
