package util_test

import (
	"testing"

	"github.com/crispybaccoon/hayashi/util"
)

func TestRemoveExtension(t *testing.T) {
	t.Run("*.*", func(t *testing.T) {
		str := "hi.world"
		exp := "hi"
		result := util.RemoveExtension(str)
		if result != exp {
			t.Fatal(result, "!=", exp)
		}
		})
	t.Run("*.", func(t *testing.T) {
		str := "hi."
		exp := "hi"
		result := util.RemoveExtension(str)
		if result != exp {
			t.Fatal(result, "!=", exp)
		}
		})
	t.Run(".*", func(t *testing.T) {
		str := ".hi"
		exp := ""
		result := util.RemoveExtension(str)
		if result != exp {
			t.Fatal(result, "!=", exp)
		}
		})
}
