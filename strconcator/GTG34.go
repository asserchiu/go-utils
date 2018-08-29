package strconcator

import "strings"

//// before
//
//func GTG34(a, b string) (result string) {
//	result = a
//	for i := 0; i < 3; i++ {
//		result += b
//	}
//	return result
//}

// after

func GTG34(a, b string) (result string) {
	var s strings.Builder
	s.WriteString(a)
	for i := 0; i < 3; i++ {
		s.WriteString(b)
	}
	return s.String()
}
