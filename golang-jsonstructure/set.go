package jsonstructure

func cloneSet(src map[string]bool) map[string]bool {
	res := make(map[string]bool)
	for k := range src {
		res[k] = true
	}
	return res
}

func keysSet(src map[string]bool) []string {
	keys := make([]string, 0, len(src))
	for k := range src {
		keys = append(keys, k)
	}
	return keys
}
