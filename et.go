package et

import "errors"

// ErrorTag is any arbitrary tag attached to an error
type ErrorTag int

// Tag associates the tag to the given error
func Tag(e error, tag ErrorTag) TaggedError {
	if e == nil {
		return nil
	}
	return taggedError{e, tag}
}

// TaggedError is an error that has been tagged
type TaggedError interface {
	error
	Tag() ErrorTag
	Unwrap() error
}

// Tagged returns true if the error has the given tag
func Tagged(e error, tag ErrorTag) bool {
	for {
		if taggedErr, ok := e.(TaggedError); ok {
			if taggedErr.Tag() == tag {
				return true
			}
			e = taggedErr.Unwrap()
		} else if unwrapped := errors.Unwrap(e); unwrapped != nil {
			e = unwrapped
		} else {
			return false
		}
	}
}

type taggedError struct {
	error
	tag ErrorTag
}

func (te taggedError) Tag() ErrorTag {
	return te.tag
}

func (te taggedError) Unwrap() error {
	return te.error
}
