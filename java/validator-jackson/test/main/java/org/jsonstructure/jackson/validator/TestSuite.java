package org.jsonstructure.jackson.validator;

import java.io.File;
import java.io.IOException;
import java.io.InputStream;
import java.net.URISyntaxException;
import java.nio.file.FileVisitResult;
import java.nio.file.FileVisitor;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.file.attribute.BasicFileAttributes;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import org.jsonstructure.jackson.validator.error.ValidationError;
import org.jsonstructure.jackson.validator.loanword.Result;
import org.junit.ClassRule;
import org.junit.Rule;
import org.junit.Test;
import org.junit.rules.ErrorCollector;

import jdk.internal.org.objectweb.asm.TypeReference;

import static org.junit.Assert.*;

public class TestSuite {

    @ClassRule
    public static ErrorCollector collector = new ErrorCollector();

    public static class TestDeclaration {
        @JsonProperty
        public String description;
        @JsonProperty
        public JsonNode structure;
        @JsonProperty
        public boolean valid;
        @JsonProperty
        public TestCase[] tests;
    }

    public static class TestCase {
        @JsonProperty
        public String description;
        @JsonProperty
        public JsonNode data;
        @JsonProperty
        public boolean valid;
    }

    private static class TestCaseVisitor implements FileVisitor<Path> {

        @Override
        public FileVisitResult preVisitDirectory(Path dir, BasicFileAttributes attrs) throws IOException {
            return FileVisitResult.CONTINUE;
        }

        @Override
        public FileVisitResult visitFile(Path file, BasicFileAttributes attrs) throws IOException {
            try {
                InputStream inputStream = Files.newInputStream(file);
                ObjectMapper mapper = Jackson.OBJECT_MAPPER;
                TestDeclaration[] decls = mapper.readValue(inputStream, mapper.getTypeFactory().constructArrayType(TestDeclaration.class));
                for (TestDeclaration decl : decls) {
                    Result<Structure, ValidationError> result = Structure.createFromNode(decl.structure, Options.defaultOpt());
                    if (result.isError() && decl.valid) {
                        collector.addError(new Throwable(
                                String.format("%s, %s.\nUnexpected JSON structure validation error: %s",
                                        file, decl.description, result.getErr())));
                    } else if (result.isOK() && !decl.valid) {
                        collector.addError(new Throwable(
                                String.format("%s, %s.\nJSON structure validation did not fail",
                                        file, decl.description)));
                    } else if (result.isOK()) {
                        Structure structure = result.getOk();
                        assertNotNull(structure);
                        for (TestCase testcase : decl.tests) {
                            ValidationError error = structure.validateFromNode(testcase.data);
                            if ((error != null) && testcase.valid) {
                                collector.addError(new Throwable(
                                        String.format("%s, %s, %s.\nUnexpected object validation error: %s",
                                                file, decl.description, testcase.description, error)));
                            } else if ((error == null) && !testcase.valid) {
                                collector.addError(new Throwable(
                                        String.format("%s, %s, %s.\nJSON object validation did not fail.",
                                                file, decl.description, testcase.description)));
                            }
                        }
                    }
                }
            } catch (Exception ex) {
                collector.addError(ex);
            }
            return FileVisitResult.CONTINUE;
        }

        @Override
        public FileVisitResult visitFileFailed(Path file, IOException exc) throws IOException {
            collector.addError(exc);
            return FileVisitResult.CONTINUE;
        }

        @Override
        public FileVisitResult postVisitDirectory(Path dir, IOException exc) throws IOException {
            return FileVisitResult.CONTINUE;
        }
    }

    @Test
    public void testSuite() throws IOException, URISyntaxException {
        ClassLoader classLoader = getClass().getClassLoader();
        Path root = Paths.get(classLoader.getResource("testsuite").toURI());
        System.out.println(root);
        Files.walkFileTree(root, new TestCaseVisitor());
    }
}
