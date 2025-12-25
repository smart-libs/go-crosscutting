package transaction

import (
	"github.com/smart-libs/go-crosscutting/event/spec/lib/pkg"
	"github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"
)

type (
	// simpleTransaction is a simple implementation of Transaction without handling the Two Phase Commit protocol
	simpleTransaction struct {
		notifier transaction.EventNotifier
		state
	}

	factory struct {
		notifierFactory transaction.EventNotifierFactory
	}
)

func NewTransactionFactory(notifierFactory transaction.EventNotifierFactory) transaction.Factory {
	return &factory{notifierFactory: notifierFactory}
}

func (f factory) Create(_ ...transaction.FactoryOptions) (transaction.Transaction, error) {
	notifier := f.notifierFactory.Create()
	tx := &simpleTransaction{
		state: newActiveState(baseState{
			notifier: notifier,
			holder:   &ResourcesHolder{},
		}),
		notifier: notifier,
	}
	EventNotifier.Notify(NewCreatedEvent(transaction.ActiveStatus))
	return tx, nil
}

func (t *simpleTransaction) Status() transaction.Status { return t.state.Status() }

// Close a simpleTransaction means to commit if it is active or rollback it if marked to rollback
func (t *simpleTransaction) Close() error {
	closeCMD := closeStateVisitor{}
	t.state.Accept(&closeCMD)
	_ = t.notifier.Close()
	return closeCMD.error
}

// EnlistResource add a resource to the simpleTransaction activating it if it was not done before
func (t *simpleTransaction) EnlistResource(key string, r transaction.Resource) (err error) {
	enlistResource := enlistResourceCommandStateVisitor{key: key, resource: r}
	t.state.Accept(&enlistResource)
	return nil
}

// GetResource returns a transactional resource enlisted in the simpleTransaction
func (t *simpleTransaction) GetResource(key string) transaction.Resource {
	getResource := getResourceCommandStateVisitor{key: key}
	t.state.Accept(&getResource)
	return getResource.resource
}

func (t *simpleTransaction) Commit() (err error) {
	defer func() {
		err = MakeError(recover(), err)
	}()
	commit := commitCommandStateVisitor{}
	t.state.Accept(&commit)
	t.state = commit.newState
	getError := &getErrorCommandStateVisitor{}
	t.state.Accept(getError)
	err = getError.error
	return
}

func (t *simpleTransaction) IsMarkedRollback() bool {
	return t.state != nil && t.state.Status() == transaction.MarkedRollbackStatus
}

func (t *simpleTransaction) Rollback() (err error) {
	defer func() {
		err = MakeError(recover(), err)
	}()
	rollback := rollbackCommandStateVisitor{}
	t.state.Accept(&rollback)
	t.state = rollback.newState
	getError := &getErrorCommandStateVisitor{}
	t.state.Accept(getError)
	err = getError.error
	return
}

func (t *simpleTransaction) MarkToRollback() {
	markToRollback := markToRollbackCommandStateVisitor{}
	t.state.Accept(&markToRollback)
	t.state = markToRollback.newState
}

func (t *simpleTransaction) Register(l event.Listener[transaction.Event]) {
	t.notifier.Register(l)
}

func (t *simpleTransaction) Unregister(l event.Listener[transaction.Event]) {
	t.notifier.Unregister(l)
}

var (
	_ transaction.EventSource = &simpleTransaction{}
)
