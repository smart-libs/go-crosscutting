package tests

import (
	"github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransactionHappyPath(t *testing.T, factory transaction.Factory) {
	t.Run("I can commit a new transaction", func(t *testing.T) {
		tx, err := factory.Create()
		assert.Equal(t, transaction.ActiveStatus, tx.Status())
		require.NoError(t, err)
		assert.NoError(t, tx.Commit())
		assert.Equal(t, transaction.CommittedStatus, tx.Status())
	})
	t.Run("I can rollback a new transaction", func(t *testing.T) {
		tx, err := factory.Create()
		assert.Equal(t, transaction.ActiveStatus, tx.Status())
		require.NoError(t, err)
		assert.NoError(t, tx.Rollback())
		assert.Equal(t, transaction.RolledBackStatus, tx.Status())
	})
	t.Run("I can mark to rollback a new transaction", func(t *testing.T) {
		tx, err := factory.Create()
		assert.Equal(t, transaction.ActiveStatus, tx.Status())
		require.NoError(t, err)
		tx.MarkToRollback()
		assert.Equal(t, transaction.MarkedRollbackStatus, tx.Status())
		assert.NoError(t, tx.Rollback())
		assert.Equal(t, transaction.RolledBackStatus, tx.Status())
	})
}

func TestTransactionFailScenarios(t *testing.T, factory transaction.Factory) {
	t.Run("I receive an error if invoke commit twice", func(t *testing.T) {
		tx, err := factory.Create()
		require.NoError(t, err)
		assert.NoError(t, tx.Commit())
		assert.Error(t, tx.Commit())
	})
	t.Run("I receive an error if invoke rollback twice", func(t *testing.T) {
		tx, err := factory.Create()
		require.NoError(t, err)
		assert.NoError(t, tx.Rollback())
		assert.Error(t, tx.Rollback())
	})
	t.Run("I receive NO error if invoke mark to rollback twice", func(t *testing.T) {
		tx, err := factory.Create()
		require.NoError(t, err)
		tx.MarkToRollback()
		tx.MarkToRollback()
		assert.NoError(t, tx.Rollback())
	})
	t.Run("I receive an error if invoke commit and next rollback", func(t *testing.T) {
		tx, err := factory.Create()
		require.NoError(t, err)
		assert.NoError(t, tx.Commit())
		assert.Error(t, tx.Rollback())
	})
	t.Run("I receive an error if invoke rollback and next commit", func(t *testing.T) {
		tx, err := factory.Create()
		require.NoError(t, err)
		assert.NoError(t, tx.Rollback())
		assert.Error(t, tx.Commit())
	})
}
