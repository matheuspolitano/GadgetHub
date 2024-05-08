package chat

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadAction(t *testing.T) {
	template, err := loadChatTemplate("../.")
	require.NoError(t, err)
	require.NotEmpty(t, template)
}

func TestGetFlow(t *testing.T) {
	template, err := loadChatTemplate("../.")
	require.NoError(t, err)
	require.NotEmpty(t, template)

	flow, err := template.GetFlow(ProductReview)
	require.NoError(t, err)
	require.NotEmpty(t, flow)
}
