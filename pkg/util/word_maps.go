package util

import (
	"strings"
	"unicode"
)

// ReverseMap - reverses a map from (k, v) to (v, k)
func ReverseMap(m map[string]int) map[int]string {
	n := make(map[int]string)
	for k, v := range m {
		if _, ok := n[v]; ok {
			n[v] = strings.Join([]string{n[v], k}, ",")
		} else {
			n[v] = k
		}
	}
	return n
}

// WordCount -returns a map of words and their counts
func WordCount(s string) map[string]int {
	strSlice := strings.Fields(stripDelimiters(s))
	result := make(map[string]int)

	for _, str := range strSlice {
		if valid(str) {
			result[strings.ToLower(str)]++
		}
	}

	return result
}

// valid - validates if the string
// is at least 4 alphabetic chars
func valid(s string) bool {
	if len(s) < 4 {
		return false
	}

	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// stripDelimiters - returns a string that is
// a copy of the orginal without the delimiters
// if we don't strip delimiters we can miss words
// that end with those delimiters
func stripDelimiters(text string) string {
	r := []rune(text)
	size := len(r)
	var prev string
	tmp := ""

	for i := 0; i < size; i++ {
		s := string(r[i])
		// Can we extend the list of delimiters?
		delimiter := strings.Contains("?!.;,=:*()[]{}/\\", s)
		if !delimiter || (s == " " && prev != " ") {
			prev = s
			tmp += s
		} else if delimiter && prev != " " {
			tmp += " "
		}
	}
	return tmp
}
