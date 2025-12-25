package transaction

type (
	markToRollbackCommandStateVisitor struct {
		panicCommandStateVisitor
		newState state
	}
)

func (v *markToRollbackCommandStateVisitor) VisitActive(s activeState) {
	v.newState = s.MarkToRollback()
}
func (v *markToRollbackCommandStateVisitor) VisitMarkedToRollback(s markedToRollbackState) {
	// Do nothing because it is already marked to rollback
	v.newState = s
}
