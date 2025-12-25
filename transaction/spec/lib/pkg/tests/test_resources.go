package tests

//go:generate mockgen -destination=./resource_mock.go -package=tests github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg Resource

import (
	"github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestTxResources(t *testing.T, factory transaction.Factory) {
	t.Run("The Resource enlisted in transaction must be retrieved using the same key used to enlist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		resourceMock := NewMockResource(ctrl)

		key := "test"
		tx, err := factory.Create()
		require.NoError(t, err)
		assert.NoError(t, tx.EnlistResource(key, resourceMock))
		assert.Equal(t, resourceMock, tx.GetResource(key))
		ctrl.Finish()
	})
	t.Run("The Resource.Prepare() and Commit() methods must be invoked when the transaction is committed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		resourceMock := NewMockResource(ctrl)
		resourceMock.EXPECT().Prepare().Return(nil)
		resourceMock.EXPECT().Commit().Return(nil)

		tx, err := factory.Create()
		require.NoError(t, err)
		assert.NoError(t, tx.EnlistResource("test", resourceMock))
		assert.NoError(t, tx.Commit())
		ctrl.Finish()
	})
	t.Run("The Resource.Rollback() method must be invoked when the transaction is rolled back", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		resourceMock := NewMockResource(ctrl)
		resourceMock.EXPECT().Rollback().Return(nil)

		tx, err := factory.Create()
		require.NoError(t, err)
		assert.NoError(t, tx.EnlistResource("test", resourceMock))
		assert.NoError(t, tx.Rollback())
		ctrl.Finish()
	})
}
