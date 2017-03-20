package org.jsonstructure.jackson.validator;

import java.util.Iterator;
import java.util.Map;

import javax.annotation.Nonnull;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.node.ArrayNode;
import com.fasterxml.jackson.databind.node.DecimalNode;
import com.fasterxml.jackson.databind.node.ObjectNode;

public class UnifyNumbers {

    @Nonnull
    public static JsonNode unify(@Nonnull JsonNode node) {
        if (node.isBigDecimal()) {
            return node;
        } else if (node.isNumber()) {
            return new DecimalNode(node.decimalValue());
        } else if (node.isObject()) {
            ObjectNode obj = (ObjectNode) node;
            Iterator<Map.Entry<String, JsonNode>> iterator = obj.fields();
            while (iterator.hasNext()) {
                Map.Entry<String, JsonNode> next = iterator.next();
                next.setValue(unify(next.getValue()));
            }
        } else if (node.isArray()) {
            ArrayNode arr = (ArrayNode) node;
            int size = arr.size();
            for (int i = 0; i < size; i++) {
                arr.set(i, unify(arr.get(i)));
            }
        }
        return node;
    }
}
