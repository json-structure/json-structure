package org.jsonstructure.jackson.validator;

import java.math.BigInteger;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.BigIntegerNode;
import com.fasterxml.jackson.databind.node.ObjectNode;
import com.fasterxml.jackson.databind.node.TextNode;

import org.jsonstructure.jackson.validator.error.ValidationError;
import org.junit.Test;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNull;

public class ComposeTest {

    @Test
    public void compose() throws Exception {
        String input = "{" +
                "\"fragments\": {" +
                "\"req\": {" +
                "\"multipleOf\": 2" +
                "}" +
                "}," +
                "\"types\": {" +
                "\"bar\": {" +
                "\"type\": \"integer\"," +
                "\"multipleOf\": 4" +
                "}," +
                "\"foo\": {" +
                "\"\u0ADD\": [ \"req\", \"bar\" ]," +
                "\"type\": \"number\"" +
                "}" +
                "}," +
                "\"main\": {" +
                "\"type\": \"foo\"" +
                "}" +
                "}";
        ObjectMapper objectMapper = Jackson.OBJECT_MAPPER;
        JsonNode tree = objectMapper.readTree(input);

        assertEquals(new TextNode("number"), tree.get("types").get("foo").get("type"));
        assertNull(tree.get("types").get("foo").get("multipleOf"));

        ValidationError errors = Compose.compose((ObjectNode) tree);
        assertNull(errors);

        assertEquals(new TextNode("number"), tree.get("types").get("foo").get("type"));
        assertEquals(new BigIntegerNode(BigInteger.valueOf(4)), tree.get("types").get("foo").get("multipleOf"));
    }

}
