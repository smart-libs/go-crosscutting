package transaction

import "fmt"

type (
	// Status represents the transaction state
	Status int
)

const (
	// ActiveStatus indicating an active
	ActiveStatus Status = 1 + iota
	// MarkedRollbackStatus code indicating a transaction that has been marked for  rollback only.
	MarkedRollbackStatus
	// PreparedStatus see below
	/**
	 *  PreparedStatus code indicating a transaction that has completed the first
	 *  phase of the two-phase commit protocol, but not yet begun the
	 *  second phase.
	 *  Probably the transaction is waiting for instruction from a superior
	 *  coordinator on how to proceed.
	 */
	PreparedStatus
	// CommittedStatus code indicating a transaction that has been committed.
	CommittedStatus
	// RolledBackStatus code indicating a transaction that has been rolled back.
	RolledBackStatus
	// UnknownStatus code indicating that the transaction state could not be
	UnknownStatus
	// NoTransactionStatus code indicating that no transaction exists.
	NoTransactionStatus
	// PreparingStatus see below
	/**
	 *  PreparingStatus code indicating a transaction that has begun the first
	 *  phase of the two-phase commit protocol, not not yet completed
	 *  this phase.
	 */
	PreparingStatus
	// CommittingStatus see below
	/**
	 *  CommittingStatus code indicating a transaction that has begun the second
	 *  phase of the two-phase commit protocol, but not yet completed
	 *  this phase.
	 */
	CommittingStatus
	// RollingBackStatus code indicating a transaction that is in the process of rolling back.
	RollingBackStatus
)

func (s Status) String() string {
	switch s {
	case ActiveStatus:
		return "ActiveStatus"
	case CommittedStatus:
		return "CommittedStatus"
	case CommittingStatus:
		return "CommittingStatus"
	case MarkedRollbackStatus:
		return "MarkedRollbackStatus"
	case NoTransactionStatus:
		return "NoTransactionStatus"
	case PreparedStatus:
		return "PreparedStatus"
	case PreparingStatus:
		return "PreparingStatus"
	case RolledBackStatus:
		return "RolledBackStatus"
	case UnknownStatus:
		return "UnknownStatus"
	default:
		return fmt.Sprintf("unknown transaction state=(%d)", s)
	}
}
