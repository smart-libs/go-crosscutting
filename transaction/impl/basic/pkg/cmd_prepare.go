package transaction

import (
	"github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"
)

type (
	prepareCommand struct {
		baseState
	}
)

func (p prepareCommand) notifyPrepareEvent(prepare func() error) error {
	err := prepare()
	if err != nil {
		p.notify(NewPrepareFailedEvent(transaction.RollingBackStatus, err))
	}
	return err
}

func (p prepareCommand) doPrepare() error {
	_, err := p.holder.foreach(0, func(handle *ResourceHandle) error {
		return handle.Prepare()
	})
	return err
}

func (p prepareCommand) Execute() error {
	return p.notifyPrepareEvent(p.doPrepare)
}
