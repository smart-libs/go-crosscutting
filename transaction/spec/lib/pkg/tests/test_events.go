package tests

//go:generate mockgen -destination=./event_notifier_mock.go -package=tests github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg EventNotifier
//go:generate mockgen -destination=./event_notifier_factory_mock.go -package=tests github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg EventNotifierFactory

import (
	"errors"
	"github.com/smart-libs/go-crosscutting/event/spec/lib/pkg"
	"github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestGlobalEvents(
	t *testing.T,
	// the globalEventNotifierSetter() function will be used by the test to create the EventNotifier instances
	// to be used by the transaction implementation
	globalEventNotifierSetter func(transaction.EventNotifier),
	factory transaction.Factory,
) {
	givenError := errors.New("any")
	givenOtherError := errors.New("any2")

	var events []transaction.Event
	eventNotifierMock := event.NotifierMock[transaction.Event]{
		OnNotify: func(t transaction.Event) {
			events = append(events, t)
		},
	}

	globalEventNotifierSetter(eventNotifierMock)

	t.Run("The global notifier must notify when a transaction is created", func(t *testing.T) {
		events = nil
		_, err := factory.Create()
		if assert.NoError(t, err) {
			if assert.Equal(t, 1, len(events)) {
				assert.Implements(t, new(transaction.CreatedEvent), events[0])
			}
		}
	})
	t.Run("The global notifier must notify when a transaction is committed", func(t *testing.T) {
		tx, err := factory.Create()
		if assert.NoError(t, err) {
			events = nil
			if assert.NoError(t, tx.Commit()) {
				if assert.Equal(t, 1, len(events)) {
					assert.Implements(t, new(transaction.CommittedEvent), events[0])
				}
			}
		}
	})
	t.Run("The global notifier must notify when a transaction is closed", func(t *testing.T) {
		tx, err := factory.Create()
		if assert.NoError(t, err) {
			events = nil
			if assert.NoError(t, tx.Close()) {
				if assert.Equal(t, 1, len(events)) {
					assert.Implements(t, new(transaction.CommittedEvent), events[0])
				}
			}
		}
	})
	t.Run("The global notifier must notify prepareFailed and rollback when a transaction prepare fails", func(t *testing.T) {
		tx, err := factory.Create()
		if assert.NoError(t, err) {
			events = nil
			// Build resource mock to fail in the prepare step
			ctrl := gomock.NewController(t)
			resourceMock := NewMockResource(ctrl)
			resourceMock.EXPECT().Prepare().Return(givenError)
			// When commit fails, the rollback is invoked
			resourceMock.EXPECT().Rollback().Return(nil)
			require.NoError(t, tx.EnlistResource("any", resourceMock))

			// Commit must fail with the prepare error
			if assert.Equal(t, givenError, tx.Commit()) {
				if assert.Equal(t, 2, len(events)) {
					assert.Implements(t, new(transaction.PrepareFailedEvent), events[0])
					assert.Implements(t, new(transaction.RolledBackEvent), events[1])
				}
			}
			ctrl.Finish()
		}
	})
	t.Run("The global notifier must notify prepareFailed and rollbackFailed when a transaction prepare and rollback fails", func(t *testing.T) {
		tx, err := factory.Create()
		if assert.NoError(t, err) {
			events = nil
			// Build resource mock to fail in the prepare step
			ctrl := gomock.NewController(t)
			resourceMock := NewMockResource(ctrl)
			resourceMock.EXPECT().Prepare().Return(givenError)
			// When commit fails, the rollback is invoked
			resourceMock.EXPECT().Rollback().Return(givenOtherError)
			require.NoError(t, tx.EnlistResource("any", resourceMock))

			// Commit must fail with the prepare error
			if assert.Equal(t, givenError, tx.Commit()) {
				if assert.Equal(t, 2, len(events)) {
					assert.Implements(t, new(transaction.PrepareFailedEvent), events[0])
					assert.Implements(t, new(transaction.RollbackFailedEvent), events[1])
				}
			}
			ctrl.Finish()
		}
	})
	t.Run("The global notifier must notify commitFailed only when a transaction commit fails", func(t *testing.T) {
		tx, err := factory.Create()
		if assert.NoError(t, err) {
			events = nil
			// Build resource mock to fail in the commit step
			ctrl := gomock.NewController(t)
			resourceMock := NewMockResource(ctrl)
			resourceMock.EXPECT().Prepare().Return(nil)
			resourceMock.EXPECT().Commit().Return(givenError)
			// Since we have just one resource and the commit fails, then no rollback is invoked for this resource
			require.NoError(t, tx.EnlistResource("any", resourceMock))

			// Commit must fail with the prepare error
			if assert.Equal(t, givenError, tx.Commit()) {
				if assert.Equal(t, 2, len(events)) {
					assert.Implements(t, new(transaction.CommitFailedEvent), events[0])
					assert.Implements(t, new(transaction.RolledBackEvent), events[1])
				}
			}
			ctrl.Finish()
		}
	})
	t.Run("The global notifier must notify when a transaction is rolled back", func(t *testing.T) {
		tx, err := factory.Create()
		if assert.NoError(t, err) {
			events = nil
			if assert.NoError(t, tx.Rollback()) {
				if assert.Equal(t, 1, len(events)) {
					assert.Implements(t, new(transaction.RolledBackEvent), events[0])
				}
			}
		}
	})
}

func TestTransactionEvents(
	t *testing.T,
	// the eventNotifierFactorySetter() function will be used by the test to set the EventNotifierFactory the
	// TransactionFactory implementation must use in the tests
	eventNotifierFactorySetter func(transaction.EventNotifierFactory) transaction.Factory,
) {
	givenError := errors.New("any")

	ctrl := gomock.NewController(t)
	factoryMock := NewMockEventNotifierFactory(ctrl)
	factoryMock.EXPECT().Create(gomock.Any()).Return(&event.OrdinaryNotifier[transaction.Event]{}).AnyTimes()

	txFactory := eventNotifierFactorySetter(factoryMock)
	t.Run("the transaction notifier must notify when the commit is done with success", func(t *testing.T) {
		tx, err := txFactory.Create()
		if assert.NoError(t, err) {
			eventSource, casted := tx.(transaction.EventSource)
			if assert.Truef(t, casted, "Transaction must implement EventSource") {
				if assert.NotNil(t, eventSource) {
					var listenerEvents []transaction.Event

					listener := event.ListenerMock[transaction.Event]{
						OnEventFunc: func(t transaction.Event) {
							listenerEvents = append(listenerEvents, t)
						},
					}

					eventSource.Register(listener)
					if assert.NoError(t, tx.Commit()) {
						assert.Equal(t, 1, len(listenerEvents))
						assert.Implements(t, new(transaction.CommittedEvent), listenerEvents[0])
					}
				}
			}
		}
	})
	t.Run("the transaction notifier must notify when the prepare fails", func(t *testing.T) {
		tx, err := txFactory.Create()
		if assert.NoError(t, err) {
			eventSource, casted := tx.(transaction.EventSource)
			if assert.Truef(t, casted, "Transaction must implement EventSource") {
				if assert.NotNil(t, eventSource) {
					var listenerEvents []transaction.Event

					listener := event.ListenerMock[transaction.Event]{
						OnEventFunc: func(t transaction.Event) {
							listenerEvents = append(listenerEvents, t)
						},
					}

					eventSource.Register(listener)

					resourceMock := NewMockResource(ctrl)
					resourceMock.EXPECT().Prepare().Return(givenError)
					// When commit fails, the rollback is invoked
					resourceMock.EXPECT().Rollback().Return(nil)
					require.NoError(t, tx.EnlistResource("any", resourceMock))

					if assert.Equal(t, givenError, tx.Commit()) {
						assert.Equal(t, 2, len(listenerEvents))
						assert.Implements(t, new(transaction.PrepareFailedEvent), listenerEvents[0])
						assert.Implements(t, new(transaction.RolledBackEvent), listenerEvents[0])
					}
				}
			}
		}
	})
	t.Run("the transaction notifier must notify when the rollback is done with success", func(t *testing.T) {
		tx, err := txFactory.Create()
		if assert.NoError(t, err) {
			eventSource, casted := tx.(transaction.EventSource)
			if assert.Truef(t, casted, "Transaction must implement EventSource") {
				if assert.NotNil(t, eventSource) {
					var listenerEvents []transaction.Event

					listener := event.ListenerMock[transaction.Event]{
						OnEventFunc: func(t transaction.Event) {
							listenerEvents = append(listenerEvents, t)
						},
					}

					eventSource.Register(listener)
					if assert.NoError(t, tx.Rollback()) {
						assert.Equal(t, 1, len(listenerEvents))
						assert.Implements(t, new(transaction.RolledBackEvent), listenerEvents[0])
					}
				}
			}
		}
	})
	ctrl.Finish()
}
