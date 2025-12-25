package transaction

import (
	"github.com/smart-libs/go-crosscutting/event/spec/lib/pkg"
	"github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"
)

type (
	detailEvent struct {
		status transaction.Status
	}

	failedEvent struct {
		err error
	}

	createdEvent struct {
		detailEvent
	}

	prepareFailedEvent struct {
		detailEvent
		failedEvent
	}

	committedEvent struct {
		detailEvent
	}

	commitFailedEvent struct {
		detailEvent
		failedEvent
	}

	rolledBackEvent struct {
		detailEvent
	}

	rollbackFailedEvent struct {
		detailEvent
		failedEvent
	}
)

var (
	// EventNotifier must be set by the component that uses the transaction implementation.
	EventNotifier transaction.EventNotifier = &event.OrdinaryNotifier[transaction.Event]{}
)

func (d detailEvent) GetTransactionStatus() transaction.Status  { return d.status }
func (f failedEvent) GetError() error                           { return f.err }
func (p prepareFailedEvent) Accept(v transaction.EventVisitor)  { v.VisitPrepareFailedEvent(p) }
func (p commitFailedEvent) Accept(v transaction.EventVisitor)   { v.VisitCommitFailedEvent(p) }
func (p rollbackFailedEvent) Accept(v transaction.EventVisitor) { v.VisitRollbackFailedEvent(p) }
func (p createdEvent) Accept(v transaction.EventVisitor)        { v.VisitCreatedEvent(p) }
func (p committedEvent) Accept(v transaction.EventVisitor)      { v.VisitCommittedEvent(p) }
func (p rolledBackEvent) Accept(v transaction.EventVisitor)     { v.VisitRolledBackEvent(p) }

func NewCreatedEvent(s transaction.Status) transaction.CreatedEvent {
	return createdEvent{
		detailEvent: detailEvent{status: s},
	}
}

func NewPrepareFailedEvent(s transaction.Status, e error) transaction.PrepareFailedEvent {
	return prepareFailedEvent{
		detailEvent: detailEvent{status: s},
		failedEvent: failedEvent{err: e},
	}
}

func NewCommitFailedEvent(s transaction.Status, e error) transaction.CommitFailedEvent {
	return commitFailedEvent{
		detailEvent: detailEvent{status: s},
		failedEvent: failedEvent{err: e},
	}
}

func NewCommittedEvent(s transaction.Status) transaction.CommittedEvent {
	return committedEvent{
		detailEvent: detailEvent{status: s},
	}
}

func NewRollbackFailedEvent(s transaction.Status, e error) transaction.RollbackFailedEvent {
	return rollbackFailedEvent{
		detailEvent: detailEvent{status: s},
		failedEvent: failedEvent{err: e},
	}
}

func NewRolledBackEvent(s transaction.Status) transaction.RolledBackEvent {
	return rolledBackEvent{
		detailEvent: detailEvent{status: s},
	}
}
