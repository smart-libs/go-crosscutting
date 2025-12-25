package transaction

import "github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"

type (
	// TODO: Implementar os states

	activeState interface {
		state
		MarkToRollback() (newState state)
		Commit() (newState state)
		Rollback() (newState state)
		EnlistResource(key string, r transaction.Resource)
		GetResource(key string) transaction.Resource
	}

	markedToRollbackState interface {
		state
		Rollback() (newState state)
		GetResource(key string) transaction.Resource
	}

	rolledBackState interface {
		state
		GetResource(key string) transaction.Resource
		// GetRollbackReason returns the error that made a commit becomes a rollback
		GetRollbackReason() error
	}

	committedState interface {
		state
		GetResource(key string) transaction.Resource
	}

	failedState interface {
		state
		GetResource(key string) transaction.Resource
		// GetReason returns the error that causes the state to be inconsistent
		GetReason() error
	}

	notCreatedState interface{ state }

	stateVisitor interface {
		VisitActive(activeState)
		VisitRolledBack(rolledBackState)
		VisitCommitted(committedState)
		VisitNotCreated(notCreatedState)
		VisitMarkedToRollback(markedToRollbackState)
		VisitFailed(failedState)
	}

	state interface {
		Status() transaction.Status
		Accept(stateVisitor)
	}
)
