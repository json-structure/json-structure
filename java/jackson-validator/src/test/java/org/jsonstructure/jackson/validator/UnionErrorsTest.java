package org.jsonstructure.jackson.validator;

import java.io.ByteArrayInputStream;
import java.io.InputStream;
import java.net.URL;
import java.nio.charset.StandardCharsets;

import org.jsonstructure.jackson.validator.error.ValidationError;
import org.jsonstructure.jackson.validator.loanword.Result;
import org.junit.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import static org.junit.Assert.assertNotNull;
import static org.junit.Assert.fail;

public class UnionErrorsTest {

    private static final Logger log = LoggerFactory.getLogger(UnionErrorsTest.class);


    @Test
    public void unionErrors() throws Exception {
        ClassLoader classLoader = getClass().getClassLoader();
        InputStream inputStream = classLoader.getResourceAsStream("union.json");
        Result<Structure, ValidationError> result = Structure.create(inputStream, Options.defaultOpt());
        if (result.isError()) {
            assertNotNull(result.getErr());
            fail(result.getErr().toString());
        }
        String input = "{\"type\": \"foo\"}";

        Structure structure = result.getOk();
        assertNotNull(structure);

        InputStream stream = new ByteArrayInputStream(input.getBytes(StandardCharsets.UTF_8));
        structure.options.unionErrors = Options.UnionErrorReport.PRIORITY;
        ValidationError error = structure.validateValue(stream);
        log.info("{}", error);

        stream = new ByteArrayInputStream(input.getBytes(StandardCharsets.UTF_8));
        structure.options.unionErrors = Options.UnionErrorReport.ALL;
        error = structure.validateValue(stream);
        log.info("{}", error);

    }

}
