package org.jsonstructure.jackson.validator;

import java.util.Arrays;
import java.util.Collections;
import java.util.HashSet;
import java.util.Set;

public class TypeDecl {

    public static final Set<String> PRIMITIVE_TYPES;
    static {
        Set<String> types = new HashSet<>(Arrays.asList(
                "boolean",
                "integer",
                "number",
                "string",
                "json",
                "struct",
                "array",
                "set",
                "map",
                "union"));
        PRIMITIVE_TYPES = Collections.unmodifiableSet(types);
    }
}
