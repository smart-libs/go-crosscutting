package transaction

import "github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"

type (
	defaultCommittedState struct {
		baseState
	}
)

func (d defaultCommittedState) Status() transaction.Status  { return transaction.CommittedStatus }
func (d defaultCommittedState) Accept(visitor stateVisitor) { visitor.VisitCommitted(d) }
func newCommittedState(b baseState) committedState          { return defaultCommittedState{b} }
