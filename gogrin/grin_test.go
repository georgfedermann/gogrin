package grin_test

import (
	"testing"

	grin "github.com/georgfedermann/gogrin/gogrin"
)

func TestGrin(t *testing.T) {
	expect := ":-> yay! \\./"
	actual := grin.Grin()
	if actual != expect {
		t.Fatal("expect:", expect, "but was actual:", actual)
	}
}
