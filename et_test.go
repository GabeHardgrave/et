package et

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

// Some example ErrorTags. For library consumers, these can be as many as you want
const (
	NotFound ErrorTag = iota
	Unauthorized
	ConnectionDropped
)

func TestTaggedError(t *testing.T) {
	baseErr := errors.New("some error")
	taggedErr := Tag(baseErr, NotFound)
	require.Equal(t, NotFound, taggedErr.Tag())
	require.Equal(t, "some error", taggedErr.Error())
	require.ErrorIs(t, taggedErr, baseErr)

	wrapped := errors.Wrap(taggedErr, "here's some more context")
	require.True(t, Tagged(wrapped, NotFound))
	require.False(t, Tagged(wrapped, Unauthorized))

	anotherTaggedErr := Tag(wrapped, Unauthorized)
	require.True(t, Tagged(anotherTaggedErr, NotFound))
	require.True(t, Tagged(anotherTaggedErr, Unauthorized))

	wrappedAgain := errors.Wrap(anotherTaggedErr, "even more nesting")

	var unauthorizedErr TaggedError
	ok := errors.As(wrappedAgain, &unauthorizedErr)
	require.True(t, ok)
	require.Equal(t, Unauthorized, unauthorizedErr.Tag())
}

func TestUntaggedError(t *testing.T) {
	err := errors.New("some error")
	require.False(t, Tagged(err, NotFound))
}

func TestNilTagging(t *testing.T) {
	taggedErr := Tag(nil, Unauthorized)
	require.Nil(t, taggedErr)
	require.False(t, Tagged(taggedErr, Unauthorized))
}
