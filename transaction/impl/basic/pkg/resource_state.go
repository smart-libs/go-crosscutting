package transaction

type (
	ResourceStateVisitor interface {
		VisitOK()
		VisitCommitFailed(err error)
		VisitRollbackFailed(err error)
	}

	ResourceState interface {
		Accept(visitor ResourceStateVisitor)
	}

	commitSucceeded   struct{}
	commitFailed      struct{ err error }
	rollbackSucceeded struct{ lastState ResourceState }
	rollbackFailed    struct {
		lastState ResourceState
		err       error
	}
)

func (c commitSucceeded) Accept(v ResourceStateVisitor)   { v.VisitOK() }
func (c commitFailed) Accept(v ResourceStateVisitor)      { v.VisitCommitFailed(c.err) }
func (r rollbackSucceeded) Accept(v ResourceStateVisitor) { v.VisitOK() }
func (r rollbackFailed) Accept(v ResourceStateVisitor)    { v.VisitRollbackFailed(r.err) }
