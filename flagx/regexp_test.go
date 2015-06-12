package flagx_test

import (
	"regexp/syntax"
	"testing"

	"github.com/tv42/cliutil/flagx"
)

func TestRegexpSetOk(t *testing.T) {
	var r flagx.Regexp
	if err := r.Set("foo"); err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}
}

func TestRegexpSetBad(t *testing.T) {
	var r flagx.Regexp
	err := r.Set("[")
	if err == nil {
		t.Fatal("expected an error")
	}
	if err, ok := err.(*syntax.Error); !ok || err.Code != syntax.ErrMissingBracket {
		t.Fatalf("got unexpected error: %T: %v", err, err)
	}
}
