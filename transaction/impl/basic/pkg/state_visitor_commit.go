package transaction

type (
	commitCommandStateVisitor struct {
		panicCommandStateVisitor
		newState state
	}
)

func (v *commitCommandStateVisitor) VisitActive(s activeState) { v.newState = s.Commit() }
