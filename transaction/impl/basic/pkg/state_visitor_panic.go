package transaction

import (
	"fmt"
	"github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"
	"strings"
)

type (
	// panicCommandStateVisitor panics for all visit. "subclasses" should "override" the methods they want to handle.
	panicCommandStateVisitor struct {
		expectedStatus []transaction.Status
	}
)

func ToString[T any](v ...T) string {
	buf := strings.Builder{}
	buf.WriteString("[")
	for i, value := range v {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf("%v", value))
	}
	buf.WriteString("]")
	return buf.String()
}

func (v panicCommandStateVisitor) createError(s transaction.Status) error {
	return WrapAsIllegalStateError(fmt.Errorf("illegal status=[%s], expected %s", s, ToString(v.expectedStatus)))
}

func (v panicCommandStateVisitor) VisitActive(_ activeState) {
	panic(v.createError(transaction.ActiveStatus))
}

func (v panicCommandStateVisitor) VisitRolledBack(_ rolledBackState) {
	panic(v.createError(transaction.RolledBackStatus))
}

func (v panicCommandStateVisitor) VisitCommitted(_ committedState) {
	panic(v.createError(transaction.CommittedStatus))
}

func (v panicCommandStateVisitor) VisitNotCreated(_ notCreatedState) {
	panic(v.createError(transaction.NoTransactionStatus))
}

func (v panicCommandStateVisitor) VisitMarkedToRollback(_ markedToRollbackState) {
	panic(v.createError(transaction.MarkedRollbackStatus))
}

func (v panicCommandStateVisitor) VisitFailed(_ failedState) {
	panic(v.createError(transaction.UnknownStatus))
}

var (
	_ stateVisitor = panicCommandStateVisitor{}
)
