package chat

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadAction(t *testing.T) {
	actions, err := loadFlowConfig("./templates")
	require.NoError(t, err)
	require.NotEmpty(t, actions)

}
