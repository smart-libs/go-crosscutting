package transaction

import (
	"github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"
)

type ResourceHandle struct {
	ResourceState
	transaction.Resource
}

func (h *ResourceHandle) CommitResource() error {
	if h == nil {
		return nil
	}
	err := h.Commit()
	if err != nil {
		h.ResourceState = commitFailed{err: err}
	} else {
		h.ResourceState = commitSucceeded{}
	}
	return err
}

func (h *ResourceHandle) RollbackResource() error {
	if h == nil {
		return nil
	}
	err := h.Rollback()
	if err != nil {
		h.ResourceState = rollbackFailed{lastState: h.ResourceState, err: err}
	} else {
		h.ResourceState = rollbackSucceeded{lastState: h.ResourceState}
	}
	return err
}

func (h *ResourceHandle) PrepareResource() error {
	if h == nil {
		return nil
	}
	return h.Prepare()
}
