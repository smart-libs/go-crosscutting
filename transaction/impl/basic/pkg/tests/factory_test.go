package tests

import (
	event "github.com/smart-libs/go-crosscutting/event/spec/lib/pkg"
	impl "github.com/smart-libs/go-crosscutting/transaction/impl/basic/pkg"
	transaction "github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"
	"github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg/tests"
	"testing"
)

func Test_Factory(t *testing.T) {
	tests.TestTransactionFactory(t, impl.NewTransactionFactory(event.DumbNotifierFactory[transaction.Event]{}))
}
