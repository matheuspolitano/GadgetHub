package chat

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseInt(t *testing.T) {
	value := "  1231   "
	valueParse, err := parseInt(value)

	require.NoError(t, err)
	require.NotEmpty(t, valueParse)

	require.Equal(t, "1231", valueParse)
}
