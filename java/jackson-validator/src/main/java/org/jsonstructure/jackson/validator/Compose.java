package org.jsonstructure.jackson.validator;

import java.util.HashSet;
import java.util.Iterator;
import java.util.Map;
import java.util.Set;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.node.ArrayNode;
import com.fasterxml.jackson.databind.node.ObjectNode;

import org.jsonstructure.jackson.validator.error.CompositeError;
import org.jsonstructure.jackson.validator.error.ValidationError;
import org.jsonstructure.jackson.validator.loanword.Slice;

import static org.jsonstructure.jackson.validator.error.ValidationError.errorAt;

public class Compose {

    public static final String COMPOSE_SYMBOL = "\u0ADD";

    private static ValidationError mapMerge(ObjectNode dst, ObjectNode src, Slice<String> scope) {
        CompositeError errors = new CompositeError();
        Iterator<Map.Entry<String, JsonNode>> iterator = src.fields();
        while (iterator.hasNext()) {
            Map.Entry<String, JsonNode> entry = iterator.next();
            Slice<String> newScope = scope.append(entry.getKey());
            JsonNode dstValue = dst.get(entry.getKey());
            if (dstValue != null) {
                boolean srcMap = entry.getValue().isObject();
                boolean dstMap = dstValue.isObject();
                if ((srcMap && !dstMap) || (!srcMap && dstMap)) {
                    ValidationError err = errorAt("Attempted merge between map and non-map types", newScope);
                    errors.add(err);
                } else if (srcMap && dstMap) {
                    ValidationError err = mapMerge((ObjectNode) dstValue, (ObjectNode) entry.getValue(), newScope);
                    errors.add(err);
                } else {
                    dst.set(entry.getKey(), entry.getValue().deepCopy());
                }
            } else {
                dst.set(entry.getKey(), entry.getValue().deepCopy());
            }
        }
        return errors;
    }

    private static ValidationError doCompose(JsonNode target,
                                             ObjectNode types,
                                             ObjectNode fragments,
                                             Set<String> prev,
                                             Slice<String> scope) {
        CompositeError errors = new CompositeError();
        if (!target.isObject()) {
            return errorAt("definition is not a JSON object", scope);
        }
        ObjectNode object = (ObjectNode) target;
        Iterator<Map.Entry<String, JsonNode>> iterator = object.fields();
        while (iterator.hasNext()) {
            Map.Entry<String, JsonNode> entry = iterator.next();
            String key = entry.getKey();
            JsonNode value = entry.getValue();
            if (value.isObject()) {
                Slice<String> newScope = scope.append(key);
                errors.add(doCompose(value, types, fragments, prev, newScope));
            }
        }
        if (errors.size() > 0) {
            return errors;
        }
        JsonNode composeNode = object.get(COMPOSE_SYMBOL);
        if (composeNode == null) {
            return null;
        }
        if (!composeNode.isArray()) {
            return errorAt("'\\u0ADD' property must be a string array", scope);
        }
        ArrayNode defs = (ArrayNode) composeNode;
        ObjectNode dest = new ObjectNode(Jackson.NODE_FACTORY);
        Iterator<JsonNode> defIter = defs.elements();
        while (defIter.hasNext()) {
            JsonNode defNode = defIter.next();
            if (!defNode.isTextual()) {
                errors.add(errorAt("'\\u0ADD' property must be a string array", scope));
                continue;
            }
            String defName = defNode.textValue();
            if (prev.contains(defName)) {
                errors.add(errorAt("cycle detected with definitions " + prev.toString(), scope));
                continue;
            }
            ObjectNode def;
            JsonNode tDef = types.get(defName);
            JsonNode fDef = fragments.get(defName);
            if ((tDef == null) && (fDef == null)) {
                errors.add(errorAt("Unknown definition " + defName, scope));
                continue;
            }
            if ((tDef != null) && (fDef != null)) {
                errors.add(errorAt("Internal error duplicate definition " + defName, scope));
                continue;
            }
            if (tDef != null) {
                Slice<String> newScope = Slice.<String>create("types", defName);
                Set<String> next = new HashSet<>(prev);
                next.add(defName);
                ValidationError err = doCompose(tDef, types, fragments, next, newScope);
                errors.add(err);
                if (err != null) {
                    continue;
                }
                def = (ObjectNode) tDef;
            } else {
                Slice<String> newScope = Slice.<String>create("fragments", defName);
                Set<String> next = new HashSet<>(prev);
                next.add(defName);
                ValidationError err = doCompose(fDef, types, fragments, next, newScope);
                errors.add(err);
                if (err != null) {
                    continue;
                }
                def = (ObjectNode) fDef;
            }
            ValidationError err = mapMerge(dest, def, scope);
            errors.add(err);
        }
        ValidationError err = mapMerge(dest, object, scope);
        errors.add(err);
        if (errors.size() > 0) {
            return errors;
        }
        object.removeAll();
        object.setAll(dest);
        object.remove(COMPOSE_SYMBOL);
        return null;
    }

    public static ValidationError compose(JsonNode shellNode) {
        CompositeError errors = new CompositeError();
        if (!shellNode.isObject()) {
            return errorAt("root object must be a JSON object", Slice.empty());
        }
        ObjectNode shell = (ObjectNode) shellNode;
        JsonNode typesNode = shell.get("types");
        JsonNode fragmentsNode = shell.get("fragments");
        JsonNode mainNode = shell.get("main");
        if ((typesNode == null) || typesNode.isNull()) {
            typesNode = new ObjectNode(Jackson.NODE_FACTORY);
        }
        if ((fragmentsNode == null) || fragmentsNode.isNull()) {
            fragmentsNode = new ObjectNode(Jackson.NODE_FACTORY);
        }
        if (!typesNode.isObject()) {
            errors.add(errorAt("'types' property must be a JSON object", Slice.empty()));
        }
        if (!fragmentsNode.isObject()) {
            errors.add(errorAt("'fragments' property must be a JSON object", Slice.empty()));
        }
        if (errors.size() > 0) {
            return errors.simplify();
        }
        ObjectNode types = (ObjectNode) typesNode;
        ObjectNode fragments = (ObjectNode) fragmentsNode;
        removeNullValues(types);
        removeNullValues(fragments);
        Set<String> dupls = intersection(types, fragments);
        if (dupls.size() > 0) {
            return errorAt("Duplicate keys across 'types' and 'fragments': " + dupls, Slice.empty());
        }
        Iterator<Map.Entry<String, JsonNode>> iterator = fragments.fields();
        while (iterator.hasNext()) {
            Map.Entry<String, JsonNode> entry = iterator.next();
            Slice<String> scope = Slice.<String>create("fragments", entry.getKey());
            Set<String> prev = new HashSet<>();
            prev.add(entry.getKey());
            if (PrimitiveTypes.PRIMITIVE_TYPES.contains(entry.getKey())) {
                errors.add(errorAt("Cannot declare fragment with primitive type name", scope));
            } else {
                errors.add(doCompose(entry.getValue(), types, fragments, prev, scope));
            }
        }
        iterator = types.fields();
        while (iterator.hasNext()) {
            Map.Entry<String, JsonNode> entry = iterator.next();
            Slice<String> scope = Slice.<String>create("types", entry.getKey());
            Set<String> prev = new HashSet<>();
            prev.add(entry.getKey());
            if (PrimitiveTypes.PRIMITIVE_TYPES.contains(entry.getKey())) {
                errors.add(errorAt("Cannot declare type with primitive type name", scope));
            } else {
                errors.add(doCompose(entry.getValue(), types, fragments, prev, scope));
            }
        }
        if (mainNode != null) {
            Slice<String> scope = Slice.<String>create("main");
            Set<String> prev = new HashSet<>();
            errors.add(doCompose(mainNode, types, fragments, prev, scope));
        }
        return errors.simplify();
    }

    private static void removeNullValues(ObjectNode types) {
        Iterator<Map.Entry<String, JsonNode>> iterator = types.fields();
        while (iterator.hasNext()) {
            JsonNode next = iterator.next().getValue();
            if (next.isNull()) {
                iterator.remove();
            }
        }
    }

    private static Set<String> intersection(ObjectNode a, ObjectNode b) {
        Set<String> aNames = new HashSet<>();
        Set<String> bNames = new HashSet<>();
        a.fieldNames().forEachRemaining(aNames::add);
        b.fieldNames().forEachRemaining(bNames::add);
        aNames.retainAll(bNames);
        return aNames;
    }
}
