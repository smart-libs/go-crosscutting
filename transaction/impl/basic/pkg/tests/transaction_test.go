package tests

import (
	"github.com/smart-libs/go-crosscutting/event/spec/lib/pkg"
	impl "github.com/smart-libs/go-crosscutting/transaction/impl/basic/pkg"
	"github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"
	"github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg/tests"
	"testing"
)

func Test_Transaction(t *testing.T) {
	tests.TestTransactionHappyPath(t, impl.NewTransactionFactory(event.DumbNotifierFactory[transaction.Event]{}))
	tests.TestTransactionFailScenarios(t, impl.NewTransactionFactory(event.DumbNotifierFactory[transaction.Event]{}))
	tests.TestTxResources(t, impl.NewTransactionFactory(event.DumbNotifierFactory[transaction.Event]{}))
	tests.TestGlobalEvents(
		t,
		func(notifier transaction.EventNotifier) {
			impl.EventNotifier = notifier
		},
		impl.NewTransactionFactory(event.DumbNotifierFactory[transaction.Event]{}),
	)
	tests.TestTransactionEvents(
		t,
		func(factory transaction.EventNotifierFactory) transaction.Factory {
			return impl.NewTransactionFactory(factory)
		},
	)
}
