package transaction

import "github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"

type (
	enlistResourceCommandStateVisitor struct {
		panicCommandStateVisitor
		key      string
		resource transaction.Resource
	}
)

func (v *enlistResourceCommandStateVisitor) VisitActive(s activeState) {
	s.EnlistResource(v.key, v.resource)
}
