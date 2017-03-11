package jsonstructure

type RegexFlavor int

const (
	StrictRegex RegexFlavor = iota
	NativeRegex
)

type JSONStructureOptions struct {
	Formats map[string]*Format
	Regex   RegexFlavor
}

func DefaultOptions() *JSONStructureOptions {
	formats := make(map[string]*Format)
	for k, v := range defaultFormats {
		formats[k] = v
	}
	flavor := StrictRegex
	return &JSONStructureOptions{
		Formats: formats,
		Regex:   flavor,
	}
}
