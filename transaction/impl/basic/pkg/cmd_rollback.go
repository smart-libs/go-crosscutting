package transaction

import (
	"github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"
)

type (
	rollbackCommand struct {
		baseState
		resourceIndexToStart int
	}
)

func (d rollbackCommand) notifyRollbackEvent(rollback func() error) error {
	err := rollback()
	if err != nil {
		d.notify(NewRollbackFailedEvent(transaction.UnknownStatus, err))
	} else {
		d.notify(NewRolledBackEvent(transaction.RolledBackStatus))
	}
	return err
}

func (d rollbackCommand) doRollback() (err error) {
	_, _ = d.holder.foreach(d.resourceIndexToStart, func(handle *ResourceHandle) error {
		innerError := handle.RollbackResource()
		if err == nil {
			err = innerError
		}
		return nil
	})
	return err
}

func (d rollbackCommand) Execute() error {
	return d.notifyRollbackEvent(d.doRollback)
}
