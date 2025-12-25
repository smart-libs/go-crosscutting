package transaction

type (
	rollbackCommandStateVisitor struct {
		panicCommandStateVisitor
		newState state
	}
)

func (v *rollbackCommandStateVisitor) VisitActive(s activeState) { v.newState = s.Rollback() }
func (v *rollbackCommandStateVisitor) VisitMarkedToRollback(s markedToRollbackState) {
	v.newState = s.Rollback()
}
