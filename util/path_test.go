package util_test

import (
	"testing"

	"github.com/crispybaccoon/hayashi/util"
)

func TestPathDetermine(t *testing.T) {
	t.Run("relative path", func(t *testing.T) {
		str := "./../core.yaml"
		exp := false
		result, is_dir, err := util.PathDetermine(str)
		if err != nil {
			t.Fatal(err)
		}
		if is_dir != exp {
			t.Fatal(result, "!=", exp)
		}
		t.Logf("%s -> %s", str, result)
		})
	t.Run("absolute path", func(t *testing.T) {
		cwd, err := util.GetCwd()
		if err != nil {
		  t.Fatal(err)
		}
		str := cwd + "/../core.yaml"
		exp := false
		result, is_dir, err := util.PathDetermine(str)
		if err != nil {
			t.Fatal(err)
		}
		if is_dir != exp {
			t.Fatal(result, "!=", exp)
		}
		t.Logf("%s -> %s", str, result)
		})
	t.Run("relative directory", func(t *testing.T) {
		str := "../../hayashi"
		exp := true
		result, is_dir, err := util.PathDetermine(str)
		if err != nil {
			t.Fatal(err)
		}
		if is_dir != exp {
			t.Fatal(result, "!=", exp)
		}
		t.Logf("%s -> %s", str, result)
		})
	t.Run(".", func(t *testing.T) {
		str := "."
		exp := true
		result, is_dir, err := util.PathDetermine(str)
		if err != nil {
			t.Fatal(err)
		}
		if is_dir != exp {
			t.Fatal(result, "!=", exp)
		}
		t.Logf("%s -> %s", str, result)
		})
}

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
