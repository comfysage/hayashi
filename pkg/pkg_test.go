package pkg

import (
	"testing"
)

func TestInferInfo(t *testing.T) {
	t.Run("Normal url", (func(t *testing.T) {
		pkg := Pkg{
			Url: "https://github.com/user/repo",
		}
		pkg.InferInfo()
		if pkg.Url == "https://github.com/user/repo" {
			return
		}
		t.Fatal(pkg.Url)
	}))
	t.Run("Go style url", (func(t *testing.T) {
		pkg := Pkg{
			Url: "github.com/user/repo",
		}
		pkg.InferInfo()
		if pkg.Url == "https://github.com/user/repo" {
			return
		}
		t.Fatal(pkg.Url)
	}))
	t.Run("short git url", (func(t *testing.T) {
		pkg := Pkg{
			Url: "user/repo",
		}
		pkg.InferInfo()
		if pkg.Url == "https://github.com/user/repo" {
			return
		}
		t.Fatal(pkg.Url)
	}))
}
