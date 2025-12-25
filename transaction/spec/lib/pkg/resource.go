package transaction

type (
	//Resource is the kind of resource that can be enlisted in a Transaction
	Resource interface {
		Prepare() error
		Commit() error
		Rollback() error
	}
)
