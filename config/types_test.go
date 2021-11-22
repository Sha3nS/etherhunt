package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInitConfig(t *testing.T) {
	err := InitConfig()
	require.NoError(t, err)

	require.NotEqual(t, Config, HostConfig{})
}
