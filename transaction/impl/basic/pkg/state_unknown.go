package transaction

import "github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"

type (
	defaultFailedState struct {
		baseState
		err error
	}
)

func (d defaultFailedState) GetReason() error            { return d.err }
func (d defaultFailedState) Status() transaction.Status  { return transaction.UnknownStatus }
func (d defaultFailedState) Accept(visitor stateVisitor) { visitor.VisitFailed(d) }

func newFailedState(b baseState, err error) failedState {
	return defaultFailedState{baseState: b, err: err}
}
