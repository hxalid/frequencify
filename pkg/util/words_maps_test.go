package util

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	extract  = "? ? ? = = = . . . Stack . . . Overflow is a question and answer site for professional and enthusiast."
	extractU = "月 ? Ɵɵ, Şş, Əçü ъзчиьбюжэё Stack Overflow is a question and answer site for professional and enthusiast."
	stripStr = "Stack test HTML.Stack test again stack."
)

func TestReverseMap(t *testing.T) {
	wc := map[string]int{
		"answer":   2,
		"overflow": 3,
		"site":     1,
		"stack":    1,
	}
	rwcExpected := map[int]string{
		1: "site,stack",
		2: "answer",
		3: "overflow",
	}
	rwc := ReverseMap(wc)
	cur := strings.Split(rwcExpected[1], ",")
	expected := strings.Split(rwc[1], ",")
	sort.Strings(cur)
	sort.Strings(expected)
	assert.True(t, reflect.DeepEqual(cur, expected))
}

func TestWordCount(t *testing.T) {
	wc := WordCount(extract)
	wcExpected := map[string]int{
		"answer":       1,
		"enthusiast":   1,
		"overflow":     1,
		"professional": 1,
		"question":     1,
		"site":         1,
		"stack":        1,
	}
	wcNotExpected := map[string]int{
		"a":            10,
		"and":          2,
		"answer":       1,
		"enthusiast":   1,
		"for":          1,
		"is":           1,
		"overflow":     1,
		"professional": 1,
		"question":     1,
		"site":         1,
		"stack":        1,
	}
	assert.True(t, reflect.DeepEqual(wcExpected, wc))
	assert.False(t, reflect.DeepEqual(wcNotExpected, wc))
}

func TestStripDelimiters(t *testing.T) {
	expected := "Stack test HTML Stack test again stack"
	s := strings.TrimSpace(stripDelimiters(stripStr))

	assert.Equal(t, expected, s)
}

func TestValid(t *testing.T) {
	testStrs := map[string]bool{
		"1abcd":    false,
		".=":       false,
		"abc":      false,
		"abcd":     true,
		"abcdefgh": true,
	}
	for k, v := range testStrs {
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, valid(k), v)
		})
	}
}
