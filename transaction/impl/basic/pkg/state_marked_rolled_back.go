package transaction

import "github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"

type (
	defaultMarkedToRollbackState struct {
		baseState
	}
)

func (d defaultMarkedToRollbackState) Status() transaction.Status {
	return transaction.MarkedRollbackStatus
}

func (d defaultMarkedToRollbackState) Accept(visitor stateVisitor) { visitor.VisitMarkedToRollback(d) }

func (d defaultMarkedToRollbackState) Rollback() (newState state) {
	err := rollbackCommand{baseState: d.baseState}.Execute()
	if err == nil {
		return newRolledBackState(d.baseState)
	}
	return newFailedState(d.baseState, err)
}

func newMarkedToRollbackState(b baseState) markedToRollbackState {
	return defaultMarkedToRollbackState{b}
}
