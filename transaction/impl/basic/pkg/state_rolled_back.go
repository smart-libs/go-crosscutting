package transaction

import "github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"

type (
	defaultRolledBackState struct {
		baseState
		error
	}
)

func (d defaultRolledBackState) GetRollbackReason() error    { return d.error }
func (d defaultRolledBackState) Status() transaction.Status  { return transaction.RolledBackStatus }
func (d defaultRolledBackState) Accept(visitor stateVisitor) { visitor.VisitRolledBack(d) }
func newRolledBackState(b baseState) rolledBackState         { return defaultRolledBackState{baseState: b} }

func newRolledBackStateWithError(b baseState, err error) rolledBackState {
	return defaultRolledBackState{baseState: b, error: err}
}
