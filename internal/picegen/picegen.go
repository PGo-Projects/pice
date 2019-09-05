package picegen

import (
	"errors"
)

type Options struct {
	UseCaps          bool
	UseNumbers       bool
	UsePunctunations bool
}

func Generate(wordCount uint64, options Options) error {
	return errors.New("Not implemented!")
}
