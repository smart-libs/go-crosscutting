package tests

import (
	"github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransactionFactory(t *testing.T, factory transaction.Factory) {
	t.Run("A factory must be provided", func(t *testing.T) {
		require.NotNil(t, factory)
	})

	t.Run("The factory must create a transaction if no argument is given", func(t *testing.T) {
		tx, err := factory.Create()
		if assert.NoError(t, err) {
			require.NotNil(t, tx)
		}
	})

	t.Run("The factory must create a transaction if nil is given as argument", func(t *testing.T) {
		tx, err := factory.Create(nil)
		if assert.NoError(t, err) {
			require.NotNil(t, tx)
		}
	})
}
