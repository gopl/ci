package foo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestF(t *testing.T) {
	require.Equal(t, "foo.F()", F())
}
