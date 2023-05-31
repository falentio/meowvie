package internal_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "meowvie/internal"
)

func TestSigner(t *testing.T) {
	s := NewSigner("testing")
	var signature string
	t.Run("sign", func(t *testing.T) {
		var err error
		signature, err = s.Sign("foo")
		require.Nil(t, err)
		require.NotNil(t, signature)
		require.Equal(t, len(signature), 64)
	})

	t.Run("compare equal", func(t *testing.T) {
		err := s.Compare("foo", signature)
		require.Nil(t, err)
	})

	t.Run("compare non equal", func(t *testing.T) {
		err := s.Compare("bar", signature)
		require.NotNil(t, err)
		require.EqualError(t, err, ErrSignatureNotMatch.Error())
	})
}
