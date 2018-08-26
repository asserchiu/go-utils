package strconcator

import (
	"strings"
)

// StringConcator the string concatenation helper
type StringConcator struct {
	strings.Builder
}

// New create a new StringConcator and append string(s)
func New(ss ...string) (sc *StringConcator) {
	sc = &StringConcator{}
	sc.WriteStrings(ss...)
	return sc
}

// WriteString append a string to the StringConcator
// Impl. stringWriter interface
func (sc *StringConcator) WriteString(s string) (n int, err error) {
	return sc.Builder.WriteString(s)
}

// WriteStrings append a list(slice) of string to the StringConcator
func (sc *StringConcator) WriteStrings(ss ...string) (n int, err error) {
	var l = 0
	for _, s := range ss {
		sc.Builder.WriteString(s)
		if err != nil {
			return l, err
		}
		l += len(s)
	}
	return l, nil
}

// String return the string content of the StringConcator
// Impl. Stringer interface
func (sc *StringConcator) String() string {
	return sc.Builder.String()
}
