package transaction

type (
	// getErrorCommandStateVisitor returns the error, if any, that caused the state chang
	getErrorCommandStateVisitor struct {
		doNothingCommandStateVisitor
		error
	}
)

func (v *getErrorCommandStateVisitor) VisitRolledBack(s rolledBackState) {
	v.error = s.GetRollbackReason()
}

func (v *getErrorCommandStateVisitor) VisitFailed(s failedState) {
	v.error = s.GetReason()
}
