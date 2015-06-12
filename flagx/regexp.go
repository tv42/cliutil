package flagx

import (
	"flag"
	"regexp"
)

// Regexp is a flag.Value that accepts regular expressions.
type Regexp struct {
	*regexp.Regexp
}

var _ flag.Value = (*Regexp)(nil)

func (r *Regexp) Set(value string) error {
	rr, err := regexp.Compile(value)
	if err != nil {
		return err
	}
	r.Regexp = rr
	return nil
}
