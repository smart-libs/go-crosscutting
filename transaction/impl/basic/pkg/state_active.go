package transaction

import "github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"

type (
	defaultActiveState struct {
		baseState
	}
)

func (d defaultActiveState) Status() transaction.Status { return transaction.ActiveStatus }
func (d defaultActiveState) Accept(v stateVisitor)      { v.VisitActive(d) }

func (d defaultActiveState) MarkToRollback() (newState state) {
	return newMarkedToRollbackState(d.baseState)
}

func (d defaultActiveState) Commit() (newState state) {
	err := prepareCommand{d.baseState}.Execute()
	lastResourceIndex := 0
	if err == nil {
		lastResourceIndex, err = commitCommand{d.baseState}.Execute()
		if err == nil {
			return newCommittedState(d.baseState)
		}
	}
	firstErr := err
	err = rollbackCommand{baseState: d.baseState, resourceIndexToStart: lastResourceIndex}.Execute()
	if err == nil && lastResourceIndex == 0 {
		return newRolledBackStateWithError(d.baseState, firstErr)
	}
	return newFailedState(d.baseState, firstErr)
}

func (d defaultActiveState) Rollback() (newState state) {
	err := rollbackCommand{baseState: d.baseState}.Execute()
	if err == nil {
		return newRolledBackState(d.baseState)
	}
	return newFailedState(d.baseState, err)
}

func (d defaultActiveState) EnlistResource(key string, r transaction.Resource) {
	d.holder.EnlistResource(key, r)
}

func newActiveState(b baseState) activeState { return &defaultActiveState{b} }
