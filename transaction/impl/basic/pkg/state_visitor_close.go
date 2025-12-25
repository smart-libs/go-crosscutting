package transaction

import "github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"

type (
	closeStateVisitor struct {
		newState state
		error
	}
)

func (v *closeStateVisitor) VisitActive(s activeState) {
	v.newState = s.Commit()
	getErrorCMD := getErrorCommandStateVisitor{}
	if v.newState.Status() != transaction.CommittedStatus {
		v.newState.Accept(&getErrorCMD)
	}
	v.error = getErrorCMD.error
}

func (v *closeStateVisitor) VisitRolledBack(s rolledBackState) {
	v.newState = s
}

func (v *closeStateVisitor) VisitCommitted(s committedState) {
	v.newState = s
}

func (v *closeStateVisitor) VisitNotCreated(s notCreatedState) {
	v.newState = s
}

func (v *closeStateVisitor) VisitMarkedToRollback(s markedToRollbackState) {
	v.newState = s.Rollback()
	getErrorCMD := getErrorCommandStateVisitor{}
	if v.newState.Status() != transaction.RolledBackStatus {
		v.newState.Accept(&getErrorCMD)
	}
	v.error = getErrorCMD.error
}

func (v *closeStateVisitor) VisitFailed(s failedState) {
	v.newState = s
}
