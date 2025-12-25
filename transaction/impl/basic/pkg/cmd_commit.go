package transaction

import (
	"github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"
)

type (
	commitCommand struct {
		baseState
	}
)

func (d commitCommand) notifyCommitEvent(commit func() (int, error)) (int, error) {
	lastResourceIndex, err := commit()
	if err != nil {
		d.notify(NewCommitFailedEvent(transaction.RollingBackStatus, err))
	} else {
		d.notify(NewCommittedEvent(transaction.CommittedStatus))
	}
	return lastResourceIndex, err
}

func (d commitCommand) doCommit() (lastResourceIndex int, err error) {
	return d.holder.foreach(0, func(handle *ResourceHandle) error {
		return handle.CommitResource()
	})
}

func (d commitCommand) Execute() (lastResourceIndex int, err error) {
	return d.notifyCommitEvent(d.doCommit)
}
