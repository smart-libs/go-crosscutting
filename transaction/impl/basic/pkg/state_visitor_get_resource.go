package transaction

import "github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"

type (
	getResourceCommandStateVisitor struct {
		doNothingCommandStateVisitor
		key      string
		resource transaction.Resource
	}
)

func (v *getResourceCommandStateVisitor) VisitActive(s activeState) {
	v.resource = s.GetResource(v.key)
}

func (v *getResourceCommandStateVisitor) VisitRolledBack(s rolledBackState) {
	v.resource = s.GetResource(v.key)
}

func (v *getResourceCommandStateVisitor) VisitCommitted(s committedState) {
	v.resource = s.GetResource(v.key)
}

func (v *getResourceCommandStateVisitor) VisitMarkedToRollback(s markedToRollbackState) {
	v.resource = s.GetResource(v.key)
}

func (v *getResourceCommandStateVisitor) VisitFailed(s failedState) {
	v.resource = s.GetResource(v.key)
}
