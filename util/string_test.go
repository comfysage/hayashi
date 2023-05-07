package util_test

import (
	"testing"

	"github.com/crispybaccoon/hayashi/util"
)

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestStringSplit(t *testing.T) {
	t.Run("no quotes", (func(t *testing.T) {
		str := "echo hello world"
		exp := []string{"echo", "hello", "world"}
		s := util.StringSplit(str)
		if Equal(s, exp) {
			t.Log(s, "==", exp)
		} else {
			t.Fatal(s, "!=", exp)
		}
	}))
	t.Run("with quotes", (func(t *testing.T) {
		str := "echo \"hello world\""
		exp := []string{"echo", "hello world"}
		s := util.StringSplit(str)
		if Equal(s, exp) {
			t.Log(s, "==", exp)
		} else {
			t.Fatal(s, "!=", exp)
		}
		str = "echo \"hello\"world"
		exp = []string{"echo", "\"hello\"world"}
		s = util.StringSplit(str)
		if Equal(s, exp) {
			t.Log(s, "==", exp)
		} else {
			t.Fatal(s, "!=", exp)
		}
	}))
	t.Run("mixed quotes", (func(t *testing.T) {
		str := "echo \"hello world\" 'hi mom'"
		exp := []string{"echo", "hello world", "hi mom"}
		s := util.StringSplit(str)
		if Equal(s, exp) {
			t.Log(s, "==", exp)
		} else {
			t.Fatal(s, "!=", exp)
		}
		str = "echo 'hello world' \"hi mom\""
		s = util.StringSplit(str)
		if Equal(s, exp) {
			t.Log(s, "==", exp)
		} else {
			t.Fatal(s, "!=", exp)
		}
	}))
	t.Run("nested quotes", (func(t *testing.T) {
		str := "echo \"hello 'world'\" 'hi \" mom'"
		exp := []string{"echo", "hello 'world'", "hi \" mom"}
		s := util.StringSplit(str)
		if Equal(s, exp) {
			t.Log(s, "==", exp)
		} else {
			t.Fatal(s, "!=", exp)
		}
		str = "echo 'hello \"world\"' \"hi ' mom\""
		exp = []string{"echo", "hello \"world\"", "hi ' mom"}
		s = util.StringSplit(str)
		if Equal(s, exp) {
			t.Log(s, "==", exp)
		} else {
			t.Fatal(s, "!=", exp)
		}
	}))
}
