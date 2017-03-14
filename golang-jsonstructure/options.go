package jsonstructure

type RegexFlavor int

type UnionErrorReport int

const (
	StrictRegex RegexFlavor = iota
	NativeRegex
)

const (
	PriorityUnionReport UnionErrorReport = iota
	AllUnionReport
)

type JSONStructureOptions struct {
	Formats    map[string]*Format
	Regex      RegexFlavor
	UnionError UnionErrorReport
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
