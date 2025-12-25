package transaction

import (
	event "github.com/smart-libs/go-crosscutting/event/spec/lib/pkg"
)

// The events specified in this package are published if there is a EventNotifierFactory available
type (
	DetailEvent interface {
		GetTransactionStatus() Status
	}

	FailedEvent interface {
		GetError() error
	}

	CreatedEvent interface {
		Event
		DetailEvent
	}
	PrepareFailedEvent interface {
		Event
		DetailEvent
		FailedEvent
	}
	CommittedEvent interface {
		Event
		DetailEvent
	}
	CommitFailedEvent interface {
		Event
		DetailEvent
		FailedEvent
	}
	RolledBackEvent interface {
		Event
		DetailEvent
	}
	RollbackFailedEvent interface {
		Event
		DetailEvent
		FailedEvent
	}

	EventVisitor interface {
		VisitCreatedEvent(CreatedEvent)
		VisitCommittedEvent(CommittedEvent)
		VisitPrepareFailedEvent(PrepareFailedEvent)
		VisitCommitFailedEvent(CommitFailedEvent)
		VisitRolledBackEvent(RolledBackEvent)
		VisitRollbackFailedEvent(RollbackFailedEvent)
	}

	Event interface {
		Accept(EventVisitor)
	}

	EventListener        event.Listener[Event]
	EventNotifier        event.Notifier[Event]
	EventNotifierFactory event.NotifierFactory[Event]

	// EventSource can be optionally implemented by Transaction. The user can access it using type cast like
	// t.(EventSource). It is also recommended that the implementation provides a global EventSource so that
	// the user can receive events from any transaction created.
	EventSource event.Source[Event]
)
