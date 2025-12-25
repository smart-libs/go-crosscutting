package transaction

import "github.com/smart-libs/go-crosscutting/transaction/spec/lib/pkg"

type (
	baseState struct {
		notifier transaction.EventNotifier
		holder   *ResourcesHolder
	}
)

func (d baseState) GetResource(key string) transaction.Resource {
	return d.holder.GetResource(key)
}

func (d baseState) notify(e transaction.Event) {
	if EventNotifier != nil {
		EventNotifier.Notify(e)
	}
	d.notifier.Notify(e)
}
