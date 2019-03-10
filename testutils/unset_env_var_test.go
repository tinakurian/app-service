package testutils

import (
	"os"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestUnsetEnvVar(t *testing.T) {
	t.Run("check unsetting and resetting of previously unset variable", func(t *testing.T) {
		// given
		varName := uuid.NewV4().String()
		os.Unsetenv(varName)
		_, present := os.LookupEnv(varName)
		require.False(t, present)

		// when
		resetFn := UnsetEnvVar(varName)

		// then
		_, present = os.LookupEnv(varName)
		require.False(t, present)

		// finally
		resetFn()
		_, present = os.LookupEnv(varName)
		require.False(t, present)
	})

	t.Run("check unsetting and resetting of previously set variable", func(t *testing.T) {
		// given
		varName := uuid.NewV4().String()
		val := "somevalue"
		os.Setenv(varName, val)
		_, present := os.LookupEnv(varName)
		require.True(t, present)

		// when
		resetFn := UnsetEnvVar(varName)

		// then
		_, present = os.LookupEnv(varName)
		require.False(t, present)

		// finally
		resetFn()
		valAfterRestoring, present := os.LookupEnv(varName)
		require.True(t, present)
		require.Equal(t, val, valAfterRestoring)
	})
}
