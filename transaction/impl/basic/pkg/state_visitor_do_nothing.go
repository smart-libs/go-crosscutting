package transaction

type (
	// doNothingCommandStateVisitor does nothing when the visit method is invoked. It was designed to be
	// used by commands that do not want to return an error for not desired method.
	doNothingCommandStateVisitor struct {
	}
)

func (v doNothingCommandStateVisitor) VisitActive(_ activeState)                     {}
func (v doNothingCommandStateVisitor) VisitCommitted(_ committedState)               {}
func (v doNothingCommandStateVisitor) VisitNotCreated(_ notCreatedState)             {}
func (v doNothingCommandStateVisitor) VisitMarkedToRollback(_ markedToRollbackState) {}
func (v doNothingCommandStateVisitor) VisitRolledBack(_ rolledBackState)             {}
func (v doNothingCommandStateVisitor) VisitFailed(_ failedState)                     {}

var _ stateVisitor = doNothingCommandStateVisitor{}
