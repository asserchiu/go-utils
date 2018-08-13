package strconcator

// StringConcator the string concatenation helper
type StringConcator struct {
	raw []byte
}

// New create a new StringConcator and append string(s)
func New(ss ...string) (sc *StringConcator) {
	sc = &StringConcator{}
	sc.WriteStrings(ss...)
	return sc
}

// WriteString append a string to the StringConcator
// Impl. stringWriter interface
func (r *StringConcator) WriteString(s string) (n int, err error) {
	r.raw = append(r.raw, s...)
	return len(s), nil
}

// WriteStrings append a list(slice) of string to the StringConcator
func (r *StringConcator) WriteStrings(ss ...string) {
	for _, s := range ss {
		r.raw = append(r.raw, s...)
	}
}

// String return the string content of the StringConcator
// Impl. Stringer interface
func (r *StringConcator) String() string {
	return string(r.raw)
}
