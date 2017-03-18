package org.jsonstructure.jackson.validator;

import java.util.Collections;
import java.util.HashMap;
import java.util.Map;

import javax.annotation.Nonnull;

import org.jsonstructure.jackson.validator.format.DateTimeFormat;
import org.jsonstructure.jackson.validator.format.EmailFormat;
import org.jsonstructure.jackson.validator.format.HostnameFormat;
import org.jsonstructure.jackson.validator.format.Ipv4Format;
import org.jsonstructure.jackson.validator.format.Ipv6Format;
import org.jsonstructure.jackson.validator.format.UriFormat;

public class Options {

    @Nonnull
    final Map<String, Format> formats;

    public static final Map<String, Format> DEFAULT_FORMATS;

    static {
        Map<String, Format> formats = new HashMap<>();
        formats.put("date-time", new DateTimeFormat());
        formats.put("email", new EmailFormat());
        formats.put("hostname", new HostnameFormat());
        formats.put("ipv4", new Ipv4Format());
        formats.put("ipv6", new Ipv6Format());
        formats.put("uri", new UriFormat());
        DEFAULT_FORMATS = Collections.unmodifiableMap(formats);
    }

    Options(@Nonnull Map<String, Format> formats) {
        this.formats = formats;
    }

    public static Options defaultOpt() {
        return new Options(DEFAULT_FORMATS);
    }
}
