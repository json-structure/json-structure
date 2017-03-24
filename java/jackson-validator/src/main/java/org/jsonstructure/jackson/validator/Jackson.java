package org.jsonstructure.jackson.validator;

import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.JsonNodeFactory;

public class Jackson {

    public static final JsonNodeFactory NODE_FACTORY = JsonNodeFactory.instance;

    public static final ObjectMapper OBJECT_MAPPER;

    static {
        OBJECT_MAPPER = new ObjectMapper();
        OBJECT_MAPPER.enable(DeserializationFeature.USE_BIG_DECIMAL_FOR_FLOATS);
        OBJECT_MAPPER.enable(DeserializationFeature.USE_BIG_INTEGER_FOR_INTS);
    }
}
