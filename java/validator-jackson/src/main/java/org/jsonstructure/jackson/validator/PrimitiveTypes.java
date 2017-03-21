package org.jsonstructure.jackson.validator;

import java.util.Arrays;
import java.util.Collections;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;

import javax.annotation.Nonnull;

public class PrimitiveTypes {

    public static final Set<String> PRIMITIVE_TYPES;

    public static final Map<String, PermissibleFields> PERMISSIBLE_FIELDS;

    static {
        Map<String, PermissibleFields> permissible = new HashMap<>();
        permissible.put("boolean", PermissibleFields.BOOLEAN);
        permissible.put("integer", PermissibleFields.INTEGER);
        permissible.put("number", PermissibleFields.NUMBER);
        permissible.put("string", PermissibleFields.STRING);
        permissible.put("json", PermissibleFields.JSON);
        permissible.put("struct", PermissibleFields.STRUCT);
        permissible.put("array", PermissibleFields.ARRAY);
        permissible.put("set", PermissibleFields.SET);
        permissible.put("map", PermissibleFields.MAP);
        permissible.put("union", PermissibleFields.UNION);
        PRIMITIVE_TYPES = Collections.unmodifiableSet(permissible.keySet());
        PERMISSIBLE_FIELDS = Collections.unmodifiableMap(permissible);
    }


    public enum PermissibleFields {
        BOOLEAN(Collections.emptyList()),
        INTEGER(Arrays.asList("multipleOf", "minimum", "maximum", "exclusiveMinimum", "exclusiveMaximum")),
        NUMBER(Arrays.asList("multipleOf", "minimum", "maximum", "exclusiveMinimum", "exclusiveMaximum")),
        STRING(Arrays.asList("pattern", "minLength", "maxLength")),
        JSON(Collections.emptyList()),
        STRUCT(Collections.singletonList("fields")),
        ARRAY(Arrays.asList("items", "minItems", "maxItems")),
        SET(Arrays.asList("items", "minItems", "maxItems")),
        MAP(Arrays.asList("items", "minItems", "maxItems")),
        UNION(Collections.singletonList("types")),
        USER_DEFINED();

        @Nonnull
        private final Set<String> permissible;

        PermissibleFields() {
            this.permissible = Collections.emptySet();
        }

        PermissibleFields(@Nonnull List<String> permissible) {
            this.permissible = new HashSet<>(permissible);
            this.permissible.add("format");
            this.permissible.add("nullable");
            this.permissible.add("optional");
            this.permissible.add("default");
            this.permissible.add("enum");
        }

        public boolean contains(String value) {
            return permissible.contains(value);
        }
    }

}
