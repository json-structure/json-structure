package org.jsonstructure.jackson.validator.simpleregexp;

import java.util.BitSet;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;

import org.antlr.v4.runtime.ANTLRErrorListener;
import org.antlr.v4.runtime.ANTLRInputStream;
import org.antlr.v4.runtime.CommonTokenStream;
import org.antlr.v4.runtime.Parser;
import org.antlr.v4.runtime.RecognitionException;
import org.antlr.v4.runtime.Recognizer;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.atn.ATNConfigSet;
import org.antlr.v4.runtime.dfa.DFA;

public class SimpleRegexp {

    private static class CapturingListener implements ANTLRErrorListener {

        String error;


        @Override
        public void syntaxError(Recognizer<?, ?> recognizer, Object o, int i, int i1, String s, RecognitionException e) {
            error = s;
        }

        @Override
        public void reportAmbiguity(Parser parser, DFA dfa, int i, int i1, boolean b, BitSet bitSet, ATNConfigSet atnConfigSet) {

        }

        @Override
        public void reportAttemptingFullContext(Parser parser, DFA dfa, int i, int i1, BitSet bitSet, ATNConfigSet atnConfigSet) {

        }

        @Override
        public void reportContextSensitivity(Parser parser, DFA dfa, int i, int i1, int i2, ATNConfigSet atnConfigSet) {

        }
    }

    @Nullable
    public static String accept(@Nonnull String input) {
        CapturingListener listener = new CapturingListener();
        ANTLRInputStream inputStream = new ANTLRInputStream(input);
        SimpleRegexpLexer lexer = new SimpleRegexpLexer(inputStream);
        TokenStream tokenStream = new CommonTokenStream(lexer);
        SimpleRegexpParser parser = new SimpleRegexpParser(tokenStream);
        parser.removeErrorListeners();
        parser.addErrorListener(listener);
        parser.start();
        return listener.error;
    }
}
