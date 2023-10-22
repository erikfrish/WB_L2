package main

import (
	"strings"
	"testing"
)

func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}

type validateStringTest struct {
	str         []rune
	expectedErr string
}

var validateStringTests = []validateStringTest{
	{[]rune("5abcd"), "error: string starts with a number"},
	{[]rune(`abcd\`), "error: last symbol is escaped symbol"},
	{[]rune("abcde"), ""},
}

func TestValidateString(t *testing.T) {
	for _, test := range validateStringTests {
		if err := validateString(test.str); !ErrorContains(err, test.expectedErr) {
			t.Errorf("Output %v not equal to expected %v", err, test.expectedErr)
		}
	}
}

type unpackTest struct {
	str         string
	expected    string
	expectedErr string
}

var unpackTests = []unpackTest{
	{"", "", ""},
	{"abcd", "abcd", ""},
	{"a4bc2d5e", "aaaabccddddde", ""},
	{"45", "", "error: string starts with a number"},
	{`qwe\4\5`, "qwe45", ""},
	{`qwe\45`, "qwe44444", ""},
	{`qwe\\5`, `qwe\\\\\`, ""},
	{`\`, "", "error: last symbol is escaped symbol"},
}

func TestUnpack(t *testing.T) {
	for _, test := range unpackTests {
		if result, err := Unpack(test.str); result != test.expected || !ErrorContains(err, test.expectedErr) {
			t.Errorf("Output %v, %v not equal to expected %v, %v", result, err, test.expected, test.expectedErr)
		}
	}
}
